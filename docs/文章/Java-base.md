# Java 基础面试题

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
- 链表长度大于 8 会改变成红黑树，8 是发生冲突的概率和转化红黑树开销之间的一种取舍

### HashMap 的扩容机制？

- 达到负载因子（默认 0.75）时，首先将数组长度扩容原来的二倍
- 扩容后，对原有节点重新哈希；如果节点是链表，则从头节点开始重新计算哈希
- 扩容期间新增的元素按照新的哈希值保存

