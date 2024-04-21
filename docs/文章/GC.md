Java 内存的自动管理，关键要解决内存的自动分配和自动回收。本文基于周志明的经典著作《深入理解 JAVA 虚拟机》介绍了内存分配会回收的一些基本策略。我们一方面要理解这些基本策略，另一方面要会通过代码验证、测试这些回收策略，且掌握这些分析方法比策略本身更有效。

## 1. 新对象优先分配在 Eden 区

当 Eden 区没有足够的空间存放对象时，将触发一次 Minor GC。

```java
public class Allocation {
    private static final int _1MB = 1024 * 1024;
    /**
     * -Xms 初始堆大小
     * -Xmx 最大堆大小
     * -Xmn 新生代大小
     * VM参数:-verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8
     */
    public static void testAllocation() {
        byte[] allocation1, allocation2, allocation3, allocation4;
        allocation1 = new byte[2 * _1MB];
        allocation2 = new byte[2 * _1MB];
        allocation3 = new byte[2 * _1MB];
        allocation4 = new byte[4 * _1MB];
    }

    public static void main(String[] args) {
        testAllocation();
    }
}
```

## 2. 大对象直接分配在老年代

避免创建“朝生夕灭”的“短命大对象”。创建大对象会导致明明还有很多内存却提前触发了垃圾回收，以获取足够的连续空间才能安置好它们，而复制对象时，意味着高额的复制开销。在 HotSpot 中可以使用 `-XX:PretenureSizeThreshold` 指定大于该值的对象直接在老年代分配，避免在 Eden 区和 两个 Survivor 区之间来回复制，产生大量的内存复制操作。
 
```java
public class Allocation2 {
    private static final int _1MB = 1024 * 1024;
    /**
     * VM参数:-verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8 * -XX:PretenureSizeThreshold=3145728
     */
    public static void testPretenureSizeThreshold() { byte[] allocation;
        allocation = new byte[4 * _1MB]; //直接分配在老年代中
    }

    public static void main(String[] args) {
        testPretenureSizeThreshold();
    }
}
```

## 3. 长期存活的对象将进入老年代

如果对象在 Survivor 区中每熬过一次 Minor GC，年龄就会增加一次。年龄是虚拟机为每个对象定义的 Age 计数器，保存在对象头中。

在 HotSpot 虚拟机中，对象在堆内存中存储布局分别为对象头（Header）、实例数据（Instance Data）和对齐填充（Padding）。其中对象头除了保存 GC 分代年龄外，还保存了哈希码（HashCode）、锁状态标志、线程持有的锁、偏向线程 ID、偏向时间戳等，这些被称为 Mark Word。

当年龄到达 15 岁时，就会晋升到老年代。15 是默认值，可以通过 `-XX: MaxTenuringThreshold` 设置其他阈值。

```
public class Allocation3 {
    private static final int _1MB = 1024 * 1024;
    /**
     * VM参数:-verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:Survivor-
     Ratio=8 -XX:MaxTenuringThreshold=1 * -XX:+PrintTenuringDistribution
     */
    @SuppressWarnings("unused")
    public static void testTenuringThreshold() {
        byte[] allocation1, allocation2, allocation3;
        // 什么时候进入老年代决定于XX:MaxTenuringThreshold 设置
        allocation1 = new byte[_1MB / 4];
        allocation2 = new byte[4 * _1MB];
        allocation3 = new byte[4 * _1MB];
        allocation3 = null;
        allocation3 = new byte[4 * _1MB];
    }

    public static void main(String[] args) {
        testTenuringThreshold();
    }
}
```

## 4. 动态对象年龄判定

HotSpot 虚拟机并不要求年龄必须达到 15 才进入老年代，还会根据一个动态机制来判断，即：如果在 Survivor 空间中相同年龄的所有对象总和大于 Survivor 空间的一半，年龄大于或等于该年龄的对象就可以直接进入老年代，无需等到 -XX:MaxTenuringThreshold=1 中配置的年龄。

```java
public class Allocation4 {
    private static final int _1MB = 1024 * 1024;
    /**
     * VM参数:-verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8
     * -XX:MaxTenuringThreshold=15
     * -XX:+PrintTenuringDistribution
     */
    @SuppressWarnings("unused")
    public static void testTenuringThreshold() {
        byte[] allocation1, allocation2, allocation3, allocation4;
        allocation1 = new byte[_1MB / 4]; // allocation1+allocation2大于survivo空间一半
        allocation2 = new byte[_1MB / 4];
        allocation3 = new byte[4 * _1MB];
        allocation4 = new byte[4 * _1MB];
        allocation4 = null;
        allocation4 = new byte[4 * _1MB];
    }

    public static void main(String[] args) {
        testTenuringThreshold();
    }
}
```

## 5. 空间分配担保

这里的规则不难，但是需要理解。在每次 Minor GC 之前，虚拟机都会检查老年代最大可用连续空间是否大于新生代的所有对象总空间，如果成立，我们认为这次 Minor GC 是安全的。因为 Minor GC 后的垃圾会进入老年代，老年代的空间够用，所以是安全的。

如果老年的空间不够呢？虚拟机会进行如下步骤：

-  检查 -XX:HandlePromotionFailur 的值
    - 如果为 True，检查老年代最大可用的连续空间是否大于历次晋升到老年代对象的平均大小
        - 是：进行 Minor GC，因为是根据“经验”判断的，所以此次 GC 可能有风险 
        - 否：进行 Full GC，不去“冒险”
    - 如果为 False，进行 Full GC

JDK6 update24 之后的规则变为：只要老年代剩余连续空间大于新生代对象总大小，就会进行 Minor GC，否则进行 Full GC。

```java
public class Allocation5 {
    private static final int _1MB = 1024 * 1024;
    /**
     * VM参数:-Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8 -XX:-Handle-
     PromotionFailure */
    @SuppressWarnings("unused")
    public static void testHandlePromotion() {
        byte[] allocation1, allocation2, allocation3, allocation4, allocation5, allocation6, allocation7;
        allocation1 = new byte[2 * _1MB];
        allocation2 = new byte[2 * _1MB];
        allocation3 = new byte[2 * _1MB];
        allocation1 = null;
        allocation4 = new byte[2 * _1MB];
        allocation5 = new byte[2 * _1MB];
        allocation6 = new byte[2 * _1MB];
        allocation4 = null;
        allocation5 = null;
        allocation6 = null;
        allocation7 = new byte[2 * _1MB];
    }


    public static void main(String[] args) {
        testHandlePromotion();
    }
}
```

## 参考

1. 深入理解 Java 虚拟机：JVM 高级特性与最佳实践. 周志明
