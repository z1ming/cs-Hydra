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

- `resize()` 时，旧数据还没有被转移到重新哈希后的位置，但这时请求的 `key` 已经会被定位到重新哈希后的位置，导致获取到空值
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
