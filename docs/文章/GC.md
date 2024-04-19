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
