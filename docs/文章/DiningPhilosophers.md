# 哲学家就餐：死锁及解决方案 Java

哲学家就餐问题是计算机科学中的一个经典问题，1971 年由荷兰计算机科学家艾兹格·迪科斯彻提出，五台计算机都试图访问五份共享的磁带时会产生问题，后来东尼·霍尔将其重新表述为哲学家就餐问题。[1]问题的详细描述可以参考 [链接](https://zh.wikipedia.org/wiki/%E5%93%B2%E5%AD%A6%E5%AE%B6%E5%B0%B1%E9%A4%90%E9%97%AE%E9%A2%98)。

## 死锁的产生

其实哲学家就餐问题描述的就是计算机中的死锁问题。那么死锁是怎么产生的呢？

五个哲学家同时拿起左边的餐叉，然后准备同时拿右边餐叉时，发生死锁。因为这时每个哲学家都想在等在右边的餐叉而无法吃饭。

## 解决方案

个人总结下来，解决这类问题主要有两个思路：互斥和顺序。

互斥就是对于一个共享资源，当一个线程占有该资源的锁时，其他线程不能占有该资源的锁。顺序就是获取多个锁时，尽量按照先后顺序获取，避免交叉。

基于以上两个思路，有一些通用的解决方案供我们学习。

### Dijkstra 算法

Dijkstra 的解决方案是每个哲学家维护三个状态：THINKING，HUNGRY，EATING，以及各自的信号量。每个哲学家对应的信号量代表左右两个餐叉是否可以被拿起来。Java 代码实现如下：

```
import java.util.Random;
import java.util.concurrent.Semaphore;
import java.util.concurrent.TimeUnit;

public class DiningPhilosophersDijkstra {

    private static final int N = 5;

    private enum State {THINKING, HUNGRY, EATING}

    private static State[] state = new State[N];
    private static final Object[] forks = new Object[N];
    private static Semaphore[] bothForksAvailable = new Semaphore[N];

    static {
        for (int i = 0; i < N; i++) {
            state[i] = State.THINKING;
            forks[i] = new Object();
            bothForksAvailable[i] = new Semaphore(0);
        }
    }

    private static int left(int i) {
        return (i - 1 + N) % N;
    }

    private static int right(int i) {
        return (i + 1) % N;
    }

    private static int myRand(int min, int max) {
        Random rnd = new Random();
        return rnd.nextInt(max - min + 1) + min;
    }

    private static void test(int i) {
        if (state[i] == State.HUNGRY &&
                state[left(i)] != State.EATING &&
                state[right(i)] != State.EATING) {
            state[i] = State.EATING;
            bothForksAvailable[i].release();
        }
    }

    private static void think(int i) throws InterruptedException {
        int duration = myRand(400, 800);
        System.out.println(i + " is thinking " + duration + "ms");
        TimeUnit.MILLISECONDS.sleep(duration);
    }

    private static void takeForks(int i) throws InterruptedException {
        synchronized (forks[i]) {
            state[i] = State.HUNGRY;
            System.out.println("\t\t" + i + " is HUNGRY");
            test(i);
        }
        bothForksAvailable[i].acquire();
    }

    private static void eat(int i) throws InterruptedException {
        int duration = myRand(400, 800);
        System.out.println("\t\t\t\t" + i + " is eating " + duration + "ms");
        TimeUnit.MILLISECONDS.sleep(duration);
    }

    private static void putForks(int i) {
        synchronized (forks[i]) {
            state[i] = State.THINKING;
            test(left(i));
            test(right(i));
        }
    }

    private static void philosopher(int i) {
        while (true) {
            try {
                think(i);
                takeForks(i);
                eat(i);
                putForks(i);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        System.out.println("dp_14");

        Thread t0 = new Thread(() -> philosopher(0));
        Thread t1 = new Thread(() -> philosopher(1));
        Thread t2 = new Thread(() -> philosopher(2));
        Thread t3 = new Thread(() -> philosopher(3));
        Thread t4 = new Thread(() -> philosopher(4));

        t0.start();
        t1.start();
        t2.start();
        t3.start();
        t4.start();
    }
}
```

### 共享资源优先级算法

顾名思义就是对餐叉划分一个优先级，必须 0～4，每个哲学家只能先拿优先级小的餐叉，后拿优先级大的餐叉。如果 4 个哲学家同时拿起左边编号较小的餐叉，第 5 个哲学家左边是优先级 4，右边是优先级 0，因为 0 已经被占用了，所以他无法拿起餐叉，因此不会发生死锁。这里也运用了顺序的思想。

但是这个方案有两个问题：

1. 如果一个工作单元持有 3 和 5，需要资源 2，需要以下操作：
    - 释放 5
    - 释放 3
    - 获取 2
    - 获取 3
    - 获取 5
2. 每个哲学家不能公平地获取到餐叉，如果哲学家 1 拿餐叉的速度很慢，哲学家 2 思考速度很快，每次都会很快把餐叉拿起来，那么哲学家 1 永远不能拿起右手边的餐叉

Java 代码实现：

```
import java.util.Random;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class DiningPhilosophersResourceHierarchy {

    static Random rnd = new Random();

    static int myRand(int min, int max) {
        return rnd.nextInt(max - min + 1) + min;
    }

    static void philosopher(int ph, Lock ma, Lock mb, Lock mo) {
        while (true) {
            int duration = myRand(200, 800);
            synchronized (mo) {
                System.out.println(ph + " thinks " + duration + "ms");
            }
            try {
                Thread.sleep(duration);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            synchronized (mo) {
                System.out.println("\t\t" + ph + " is hungry");
            }
            synchronized (ma) {
                synchronized (mb) {
                    duration = myRand(200, 800);
                    synchronized (mo) {
                        System.out.println("\t\t\t\t" + ph + " eats " + duration + "ms");
                    }
                    try {
                        Thread.sleep(duration);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        }
    }

    public static void main(String[] args) {
        System.out.println("dining Philosophers C++11 with Resource hierarchy");
        Lock m1 = new ReentrantLock();
        Lock m2 = new ReentrantLock();
        Lock m3 = new ReentrantLock();
        Lock m4 = new ReentrantLock();
        Lock m5 = new ReentrantLock();
        Lock mo = new ReentrantLock();

        Thread t1 = new Thread(() -> philosopher(1, m1, m2, mo));
        Thread t2 = new Thread(() -> philosopher(2, m2, m3, mo));
        Thread t3 = new Thread(() -> philosopher(3, m3, m4, mo));
        Thread t4 = new Thread(() -> philosopher(4, m4, m5, mo));
        Thread t5 = new Thread(() -> philosopher(5, m1, m5, mo));

        t1.start();
        t2.start();
        t3.start();
        t4.start();
        t5.start();

        try {
            t1.join();
            t2.join();
            t3.join();
            t4.join();
            t5.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
```

### Chandy-Misra 算法

这个算法可以很好地避免死锁的发生，具体流程是：

1. 每个餐叉有个状态，脏或者干净，初始所有的餐叉都是脏的
2. 每个哲学家只会获取左右哲学手中的餐叉，每次尝试获取左右两个餐叉，如果某个餐叉被占用，则向对应的哲学家发送一个消息
3. 哲学家收到消息，如果餐叉是干净的，则不理会；如果是脏的，则擦干净并交出餐叉
4. 哲学家吃完后餐叉就变脏了，如果这时有哲学家请求，就擦干净并交出餐叉

Java 实现如下：

```
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * forks:
 * 0 x 1 x 2 x 3 x 4 x
 * philosophers:
 * x 0 x 1 x 2 x 3 x 4
 */
public class DiningPhilosophersChandyMisra {
    private final int NP;
    private final Lock[] forks;
    private final boolean[] dirty;
    private final boolean[] hungry;

    public DiningPhilosophersChandyMisra(int NP) {
        this.NP = NP;
        this.forks = new Lock[NP];
        this.dirty = new boolean[NP];
        this.hungry = new boolean[NP];
        for (int i = 0; i < NP; i++) {
            this.hungry[i] = true;
            this.forks[i] = new ReentrantLock();
            this.dirty[i] = true;
        }
    }

    private int leftFork(int p) {
        return p;
    }

    private int rightFork(int p) {
        return (p + 1) % NP;
    }

    private int leftPhilosopher(int p) {
        return (p - 1) % NP;
    }

    private int rightPhilosopher(int p) {
        return (p + 1) % NP;
    }

    private void think(int p) {
        System.out.println(Thread.currentThread().getName() + ": philosopher " + p + " is thinking.");
        hungry[p] = true;
    }

    private void eat(int p) {
        System.out.println(Thread.currentThread().getName() + ": philosopher " + p + " is eating.");
        forks[leftFork(p)].unlock();
        forks[rightFork(p)].unlock();
        hungry[p] = false;
    }

    private boolean obtainBothForks(int p) {
        boolean leftForkAcquired = forks[leftFork(p)].tryLock();
        boolean rightForkAcquired = forks[rightFork(p)].tryLock();
        if (leftForkAcquired && rightForkAcquired) {
            return true;
        } else {
            if (leftForkAcquired) {
                forks[leftFork(p)].unlock();
            } else {
                receiveRequest(leftPhilosopher(p), true);
            }

            if (rightForkAcquired) {
                forks[rightFork(p)].unlock();
            } else {
                receiveRequest(rightPhilosopher(p), false);
            }
            return false;
        }
    }

    private void receiveRequest(int p, boolean rightFork) {
        int forkId = rightFork ? rightFork(p) : leftFork(p);
        if (!dirty[forkId]) {
            return;
        }

        dirty[forkId] = false;
        if (forks[forkId].tryLock()) {
            forks[forkId].unlock();
        }
    }

    public void startDining() {
        Thread[] philosophers = new Thread[NP];
        for (int i = 0; i < NP; i++) {
            final int philosopherId = i;
            philosophers[i] = new Thread(() -> {
                while (true) {
                    if (obtainBothForks(philosopherId) && hungry[philosopherId]) {
                        eat(philosopherId);
                    } else {
                        think(philosopherId);
                    }
                }
            });
            philosophers[i].start();
        }
    }

    public static void main(String[] args) {
        int NP = 5;
        DiningPhilosophersChandyMisra diningPhilosophersChandyMisra = new DiningPhilosophersChandyMisra(NP);
        diningPhilosophersChandyMisra.startDining();
    }

}
```

### 服务生解法

服务生解法就是引入服务生，每次必须经过服务生的允许后才能获取餐叉。因为服务生知道当前哪个餐叉正在被使用，所以可以避免死锁发生。

假如当前有 0～4 五个哲学家，0,2 正在吃饭，这时 1 由于两个餐叉都被占用而无法吃饭，对于 3，如果这时拿起剩余的一只餐叉，就有可能发生死锁。如果引入服务生，则服务生会让 3 等待，直到两只餐叉都可用时才去拿。

但是这里我一直不太理解的是，为什么 3 会拿起右边的餐叉？3 等待 2 吃完后先拿左边的餐叉，再拿右边的餐叉不就好了吗？希望有懂得大佬帮忙解释下。

### 限制就餐人数

这个思路和服务生解法类似，只是方式不同。这种方式要求每次只能有 `n-1` 个哲学家就餐，这样每次都会剩余一个哲学家处于等待状态，当有一个哲学家吃完时他才能做下就餐。

## 总结

学习技术时要从根源学起，文档也最好追溯到原版英文的文档。哲学家就餐问题属于经典的死锁问题，所以有必要好好学习一番。wiki 中的英文文档介绍得非常详细，我这里只是抛砖引玉了，除此之外，wiki 中引用的参考文献都是深入了解死锁和哲学家就餐问题的极好材料。


## 参考资料

1. [哲学家就餐问题 wiki](https://zh.wikipedia.org/wiki/%E5%93%B2%E5%AD%A6%E5%AE%B6%E5%B0%B1%E9%A4%90%E9%97%AE%E9%A2%98)
2. [经典并发问题: 哲学家就餐问题](https://colobu.com/2022/02/13/dining-philosophers-problem/)
