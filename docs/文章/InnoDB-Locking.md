# InnoDB 锁介绍

## 前言

个人学习 MySQL 锁时，网上的文章基本都是按照不同类型做了不同分类，很难记住，原因主要有下：

- 还是没有理解为什么会出现这些锁，这些锁解决了什么问题
- 网络上的文章大多是你参考我，我参考你，每个人加入自己的理解，导致主观的成分太多

本着查找根源的思路，我在 MySQL 的[官方文档](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)上找到了最原始对于 InnoDB 的讲解，以下是对该文档的笔记和理解，仅供参考学习之用。

## InnoDB 锁

InnoDB 共有以下几种锁：

- 共享锁和排他锁(Shared(S) and Exclusive(X) Locks)
- 意向锁(Intention Locks)
- 记录锁(Record Locks)
- 间隙锁(Gap Locks)
- 临键锁(Next-Key Locks)
- 插入意向锁(Insert Intention Locks)
- 自增锁(AUTO-INC Locks)
- 空间索引锁(Predicate Locks for Spatial Indexes)

### 共享锁和排他锁(Shared(S) and Exclusive(X) Locks)

InnoDB 实现了标准的行级锁定，其中有两种类型的锁：共享锁(S)和排他锁(X)。

- 共享锁允许持有该锁的事务读取行
- 排他锁允许持有该锁的事务更新或删除行

如果事务 T1 持有 r 行的 S 锁，那么另一个事务 T2 对 r 行的请求按如下方式处理：

1. T2 对 S 锁的请求可以立刻批准，结果是 T1 和 T2 都持有 r 行的 S 锁
2. T2 对 X 锁的请求无法立刻批准

如果事务 T1 持有 r 行的 X 锁，那么另一个事务 T2 对 r 行任一类型的锁请求都不会立即批准，只能等待 T1 释放 X 锁。

### 意向锁(Intention Locks)

InnoDB 支持多粒度锁定，允许行锁和表锁共存。意向锁是表级别的锁，表明稍后需要对表中的行设置那种类型的锁。意向锁有两种类型：

- 意向共享锁(IS): 表示事务打算在个别行上设置共享锁
- 意向排他锁(IX): 表示事务打算在个别行上设置排他锁

使用 `select ... for share` 设置 IS 锁，使用 `select ... for update` 设置 IX 锁。

规则如下：

1. 在事务可以获取 S 锁之前，必须先获取 IS 锁或更强的锁
2. 在事务可以获取 X 锁之前，必须先获取 IX 锁

兼容性如下：

||X|IX|S|IS|
|:-:|:-:|:-:|:-:|:-:|
|X|Conflict|Conflict|Conflict|Conflict|
|IX|Conflict|Compatible|Conflict|Compatible|
|S|Conflict|Conflict|Compatible|Compatible|
|IS|Conflict|Compatible|Compatible|Compatible|

如果一个事务请求加的锁与现有的锁兼容，那么该事务将获得该锁，否则事务将等待，直到现有锁被释放。如果请求的锁冲突，并可能发生死锁，这时会报错。

意向锁只会发送一个全表的请求(比如 LOCK TABLES ... WRITE)，不会阻塞任何东西，它只是表明有人正在锁定一行，或者即将锁定一行。

使用命令 `SHOW ENGINE INNODB STATUS` 将显示如下信息：

```
TABLE LOCK table `test`.`t` trx id 10080 lock mode IX
```



## 总结

## 参考

