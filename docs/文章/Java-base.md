# Java 面试题-基础

## 面向对象？多态？重载？

面向对象的三大特点：

- 继承
- 封装
- 多态

多态是指同一个方法名在不同的对象上可以有不同的行为。可通过继承和接口实现：

- 子类继承父类，可重写父类的方法
- 不同的类可以实现相同的接口

## 重载和重写的区别是什么？

- 重载方法名相同，参数列表不同（参数的类型，数量，顺序），返回值可能也不同
- 重写是子类继承父类，子类重新实现父类的方法，方法名和参数列表均相同

实现一个常用的多态的例子：

```
public class Animal {
    public void say() {
        System.out.println("dddd");
    }

    public static void main(String[] args) {
        Dog dog = new Dog();
        dog.say();

        Pig pig = new Pig();
        pig.say();

        Bird bird = new Bird();
        bird.say();
    }
}

class Dog extends Animal {
    public void say() {
        System.out.println("汪");
    }
}

class Pig extends Animal {
    public void say() {
        System.out.println("哼");
    }
}

class Bird extends Animal {
    public void say() {
        System.out.println("叽喳");
    }
}
```

如上例子实现了一个 `Animal` 类，`Dog`, `Pig`, `Bird` 类分别继承了 `Animal` 类，并重写了 `say()` 方法，这样实例化不同的类，调用同一个方法 `say()` 时，会打印不同的叫声。

## HashMap

### HashMap 的数据结构是什么？

- JDK1.7 及之前版本采用数组+链表的方式， JDK1.8 开始采用数组+链表/红黑树的方式
- 链表长度大于 8 会改变成红黑树，小于 6 时会从红黑树退化为链表

### HashMap 的扩容机制？

- 判断老表容量是否超过上限，是修改为 Integer.MAX_VALUE
- 否将容量和阈值都修改为原来 2 倍，遍历老数组，如果索引位置有一个节点，直接迁移到新位置，如果大于 1 个，则判断是红黑树节点还是链表节点，如果是链表，先保存头结点，然后依次计算后续的节点
- Jdk1.8 之前头插法，1.8 开始尾插法

### HashMap 是线程安全的吗？

线程不安全。

- JDK1.7 采用数据+链表，多线程下，扩容时存在 Entry 链死循环和数据丢失问题。
- JDK1.8`resize()` 时，旧数据还没有被转移到重新哈希后的位置，但这时请求的 `key` 已经会被定位到重新哈希后的位置，导致获取到空值，这条暂时不确定
- 多线程 `put` 时可能会数据覆盖。如果两个不同的 key 发生哈希冲突，可能会只新增一个列表节点而不是两个

### HashTable 和 ConcurrentHashMap 的区别是什么？

- HashTable 在整个方法加锁，ConcurrentHashMap 在每个链表头节点加锁，不会发生锁冲突
- HashTable 使用 `synchronized` 加锁，ConcurrentHashMap 使用 CAS，后者效率更高
- HashTable `resize` 时旧元素搬到新空间，然后释放旧空间，大量拷贝，效率低；ConcurrentHashMap 每次拷贝一部分，新旧空间同时存在
- HashTable `get` 加锁，ConcurrentHashMap `get` 不加锁，原因是 Node 中的 val(和 next) 使用 volatile 修饰

### 为什么引入红黑树，不引入其他树？

红黑树相比于其他树，性能和稳定性更好，具体来说：

1. 为什么不用二叉排序树？二叉排序树在极端条件下可能出现线性结构，比如每次添加的元素均小于或大于当前所有元素，则该树变为线性结构。这时查询的效率和链表一样，所以不用二叉排序树
2. 为什么不用平衡二叉树（AVL 树）？AVL 树是严格平衡的二叉树，而红黑树是基于 2-3 树演变而来，没有严格平衡，因此 AVL 树为了保持平衡进行的旋转次数要多于红黑树，性能也就不如红黑树

### HashMap 出现红黑树会一直增高变成无限高的情况吗？

不会。集合中节点的数量超过阈值，HashMap 会进行扩容，原始红黑树的节点会被打散，可能会退款成链表结构。

### HashMap 读和写的时间复杂度是多少？

读、写（插入、更新、删除）的时间复杂度均为 O(1)。使用键值对存储数据，可直接计算出哈希值来定位到对应的存储位置。

### 怎么解决 HashMap 线程不安全的问题？

- 使用 `ConcurrentHashMap`：ConcurrentHashMap 是线程安全的哈希表实现，通过分段锁和 CAS 保证线程安全
- 使用 `Collections.synchronizedMap`：该方法会返回同步 Map 对象，但性能不如 ConcurrentHashMap

### 拓展：解决线程安全问题还有哪些办法？

- 使用同步关键字 `synchronized`：同一时刻只有一个线程可以访问共享资源
- 使用 `volatile` 关键字：一个线程修改了共享变量的值，其他线程可以立即看到该值
- 使用线程安全的工具类：如 `AtomicInteger`,`AutomicLong`,`CountDownLatch` 等线程安全的工具类
- 使用多线程并发容器：如 `ConcurrentLinkedQueue`, `CopyOnWriteArrayList` 等

## Java 并发

### volatile 关键字底层是怎么实现的？如何保证内存可见性？

volatile 通过以下两种机制保证内存可见性：

- 禁止指令重排：在程序执行时，为了提高性能，编译器和处理器可能会指令重排序，导致变更的更新操作被延迟执行或乱序执行，其他线程无法看到最新的值。使用 volatile 关键字修饰的变量会禁止指令重排序，保证变量更新操作按照代码顺序执行。
- 内存屏障：在多核处理器架构下，每个线程都有自己的缓存，volatile 关键字会在写操作后插入写屏障（Write Barrier），在读操作前插入读屏障（Read Barrier），确保变量的更新能够立即被其他线程看到，保证内存的可见性。

通过指令重排序和插入内存屏障，volatile 关键字能够保证被修饰变量的更新操作对其他线程是可见的，从而有效解决了多线程环境下的内存可见性问题。

