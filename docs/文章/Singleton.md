1、可重入的方式如下，其中 volatile 保证多线程场景下实例的可见性

```java
public class Singleton {
    private volatile static Singleton instence = null;
    private Singleton() {
        
    }
    
    public static Singleton getInstance() {
        if (instence == null) {
            synchronized (Singleton.class) {
                if (instence == null) {
                    instence = new Singleton();
                }
            }
        }
        return instence;
    }
}
```

2. 枚举式实现

《Effective Java》一书说枚举式是最佳实践，如下：

```java
public enum Singleton1 {
    INSTANCE;

    public void doSomething() {
        System.out.println("doSomething");
    }

    public static void main(String[] args) {
        Singleton1.INSTANCE.doSomething();
    }
}
```
