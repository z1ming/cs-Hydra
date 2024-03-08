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

### 为什么需要保证内存可见性

为了解决数据一致性。如果不保证内存可见性，一个线程改变了共享变量的值，另一个线程无法看到最新的值，没有使用最新的数据，导致产生错误的结果。

### volatile 为什么要禁止指令重拍，能举一个具体的指令重拍出现的例子吗？

禁止指令重拍是为了保证程序的执行顺序和编写的顺序一致，特别是在多线程环境下，避免出现意外的结果。

如:

```
int a = 0;
boolean flag = false;

a = 1; 
flag = false;

if (flag) {
    System.out.println(a);
}
```

如果发生指令重排，可能发生在判断 flag 先于 `a` 的赋值操作，从而打印出 `0`。如果禁止指令重排，则在多线程环境下保证代码按顺序执行，保证 `System.out.println(a);` 可以输出 1。

### Synchronized 的底层原理是什么，锁升级的过程了解吗？

- Synchronized 底层使用 monitor 对象锁实现，每一个对象关联一个 monitor 对象，而 monitor 对象锁是互斥的，同一个时刻只能有一个线程持有对象锁，其他线程想再获取对象锁时会被阻塞住，这样就能保证拥有对象锁的线程可以安全执行临界区的代码
- 锁升级是指 jvm 根据锁的竞争对象和对象的状态，将对象的锁从偏向锁、轻量锁升级为重量级锁的过程。
    - 偏向锁是指针对无竞争的情况下，锁会偏向于第一个获取锁的线程
    - 轻量锁是指段时间内只有一个线程竞争锁的情况下，使用 CAS 操作来避免阻塞
    - 重量级锁是指多个线程竞争同一个锁时，通过操作系统的互斥量来实现线程阻塞和唤醒。
    - 锁升级为了提供多线程并发访问的效率和性能

### 线程是怎么拿到锁的？

检查锁的状态，并尝试获取锁。在 JVM 中，锁信息具体是存放在 Java 对象头中的。

当一个线程尝试进入 synchronized 代码块或方法时，JVM 会检查对应对象的锁状态。如果一个对象的锁未被其他线程持有，即锁状态为可获取，那么该线程将成功获取锁并进入临界区执行代码。

### 锁信息具体是放在哪里的？

锁状态信息是 Java 对象头中的 Mark Word 字段，保存了锁的信息、垃圾回收信息等。Java 对象在内存中有如下字段：

- 对象头
    - Mark Word
    - Class Pointer
    - Length
- 对象实际数据
    - Instance Data/Array Data
- 对齐填充
    - Padding

JVM 通过操作对象的头部信息来实现锁的获取、释放以及等待队列的管理。当线程成功获取锁后，对象的头部信息会被更新为当前线程的标识，表示该线程拥有了这个锁。

其他线程在尝试获取同一个锁时，会检查对象的头部信息，如果锁已经被其他线程持有，他们将会被阻塞直到锁被释放。

### Synchronized 锁和 ReentrantLock 加锁有什么区别？

- 用法不同：synchronized 可用来修饰普通方法、静态方法和代码块，而 ReentrantLock 只能用在代码块上
- 获取锁和释放锁的方式不同：synchronized 会自动加锁和释放锁，当进入 synchronized 修饰的代码块之后会自动加锁，当离开 synchronized 的代码段之后会自动释放锁。ReentrantLock 需要手动加锁和释放锁
- 锁类型不同：synchronized 属于非公平锁，而 ReentrantLock 既可以是公平锁也可以是非公平锁
- 响应中断不同：ReentrantLock 可以响应中断，synchronized 不可以
- 底层实现不同：synchronized 是 JVM 层面通过监视器实现的，ReentrantLock 是基于 AQS 实现的

## Java 线程池

### 线程池了解过吗？有哪些核心参数？

线程吃是为了减少频繁的创建线程和销毁线程带来的性能损耗。

线程池分为核心线程池，线程池的最大容量，还有等待任务的队列，提交一个任务，如果核心线程没有满，就创建一个线程，如果满了就加入等待队列，如果等待队列满了，就会增加线程，如果达到最大线程数量，就按照丢弃策略处理。

一共有 7 个参数：

```
public ThreadPoolExecutor(int corePoolSize,
                          int maximumPoolSize,
                          long keepAliveTime,
                          TimeUnit unit,
                          BlockingQueue<Runnable> workQueue,
                          ThreadFactory threadFactory,
                          RejectedExecutionHandler handler) 
```

- corePoolSize: 核心线程数。默认情况下，线程池中线程的数量 <= corePoolSize，即使线程处于空闲状态也不会被销毁
- maximumPoolSize: 最大线程数，即线程池中最多可以容纳的线程数量。
- keepAliveTime: 当线程池中线程的数量大于 corePoolSize，并且某个线程的空闲时间超过了 keepAliveTime，那么这个线程就会被销毁
- unit: keepAliveTime 的时间单位
- workQueue: 工作队列。当没有空闲的线程执行新任务时，该任务就会被放入工作队列
- threadFactory: 线程工厂。可以给线程取名字等
- handler: 拒绝策略。当一个新任务交给线程池，如果此时线程池中有空闲的线程，就会直接执行，如果没有空闲的的线程，就放入阻塞队列中，如果阻塞队列满了，就会创建一个新线程，从阻塞队列头部取出一个任务来执行，并将新任务加入到阻塞队列的末尾。如果当前线程池中线程的数量等于 maximumPoolSize，就不会创建新线程，就会去执行拒绝策略

### 为什么核心线程满了之后是先加入阻塞队列而不是直接加到总线程？

- 线程池创建线程需要获取 mainLock 这个全局锁，会影响并发效率，所以使用阻塞队列吧第一步创建核心线程与第三步创建最大线程隔离开来，起一个缓冲的作用
- 引入阻塞队列，是为了在执行 `execute()` 方法时，尽可能地避免获取全局锁

### 核心线程数一搬设置为多少？

假设机器有 N 个 CPU:

- 如果是 CPU 密集型应用，则线程池大小设置为 N+1，线程的应用场景：主要是复杂算法
- 如果是 IO 密集型应用，则线程池大小设置为 2N+1，线程的应用场景：数据库数据的交互，文件的上传下载，网络数据传输等

如果同时有计算工作和 IO 工作的任务，应该考虑使用两个线程池，一个处理计算任务，一个处理 IO 任务，分别对两个线程池按照计算密集型和 IO 密集型来设置线程数。

### IO 密集型线程数为什么设置为 2N+1？

在 IO 密集型任务中，线程通畅会因为 IO 操作而阻塞，此时可以让其他线程继续执行，充分利用 CPU 字段。设置为 2N+1 可以保证在有多个线程阻塞时，仍有足够的线程可以继续执行。

### String，StringBuilder，StringBuffer 区别？单线程大量操作字符串用哪个？

- String 是不可变字符序列，每次对 String 进行修改时都会创建一个新的 String 对象，因此在大量操作字符串时，使用 String 会频繁创建对象，导致性能较低
- StringBuilder 线程不安全，
- StringBuffer 线程安全，性能不如 StringBuilder，因为 StringBuffer 所有共有的方法都是同步的

因此单线程场景下，使用 StringBuilder 性能更好，多线程场景下使用 StringBuffer 能保证线程安全。

### synchronized 偏向锁直接升级为重量级锁吗？重量级锁是怎么实现的？

偏向锁不会直接升级为重量级锁，而是先升级为轻量级锁，如果轻量级锁竞争失败，则再升级为重量级锁。

重量级锁一般是通过操作系统的互斥量（mutex）来实现的，当一个线程获取重量级锁时，会将该线程挂起，直到锁被释放。这种锁的性能比较低，因为每次加锁和释放锁都需要涉及到操作系统的系统调用，开销比较大。因此在实际应用中，应尽量避免使用重量级锁。

### Java 中的异常分类？

- Checked Exception（受检异常）：这种异常在编译时就可以被检测出来，必须在代码中声明或者抛出，否则编译不通过。一般由外部环境引起，如 `IOException`，`SQLException` 等
- Unchecked Exception（非受检异常）：程序内部错误导致，这类代码不用显式声明抛出，抛出后不处理程序会崩溃，如 `NullPointerException`, `ArrayIndexOutOfBoundsException`, `IllegalArgumentException` 等
- Error（错误）：这类错误无法捕获，通常由 JVM 或内存不足引起。如 `OutOfMemory Error`, `StackOverFlowError` 等

Java 中，可以使用 `try ... catch` 捕获异常，使用 `throw` 手动抛出异常。

