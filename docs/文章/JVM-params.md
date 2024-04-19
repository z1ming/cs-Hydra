# -XX:SurvivorRatio

SurvivorRatio 称为幸存者比例。幸存者就是每次垃圾回收后没有被回收的对象，他们被放到了 Survivor1 和 Survivor2，所以 Survivor1 和 Survivor2 就是幸存者。

`-XX:SurvivorRatio=6` 就代表每个 Survivor 和 Eden 的比率是 1:6，由于有两个 Survivor，所以 Survivor1、Survivor2 各占 1/8，Eden 占 2/8。 如果 Survivor 空间太小，回收的对象会直接溢出到老年代；如果 Survivor 空间太大，会有大量空闲空间。在每次 GC 之前，JVM 都会确定对象在保留之前可以复制的次数，成为保留阈值。选择这个阈值是为了使 survivor 空间保持半满。

使用选项 `-XX:+PrintTenuringDistribution` 可以显示新生代中对象的阈值和年龄，对于观察应用程序的生命周期分布很有用。
