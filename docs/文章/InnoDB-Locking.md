# InnoDB 锁介绍

## 前言

个人学习 MySQL 锁时，网上的文章基本都是按照不同类型做了不同分类，很难记住，原因主要有下：

- 还是没有理解为什么会出现这些锁，这些锁解决了什么问题
- 网络上的文章大多是你参考我，我参考你，每个人加入自己的理解，导致主观的成分太多

本着查找根源的思路，我在 MySQL 的[官方文档](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)上找到了最原始对于 InnoDB 的讲解，以下是对该文档的笔记和理解，仅供参考学习之用。

## InnoDB 锁

InnoDB 共有以下几种锁：

- Shared(S) and Exclusive(X) Locks 
- Intention Locks
- Record Locks
- Gap Locks
- Next-Key Locks
- Insert Intention Locks
- AUTO-INC Locks
- Predicate Locks for Spatial Indexes

### Shared(S) and Exclusive(X) Locks

InnoDB 实现了标准的行级锁定，其中有两种类型的锁：共享锁(S)和排他锁(X)。

- 共享锁允许持有该锁的事务读取行
- 排他锁允许持有该锁的事务更新或删除行

如果事务 T1 持有 r 行的 S 锁，那么另一个事务 T2 对 r 行的请求按如下方式处理：

1. T2 对 S 锁的请求可以立刻批准，结果是 T1 和 T2 都持有 r 行的 S 锁
2. T2 对 X 锁的请求无法立刻批准

如果事务 T1 持有 r 行的 X 锁，那么另一个事务 T2 对 r 行任一类型的锁请求都不会立即批准，只能等待 T1 释放 X 锁。

## 总结

## 参考

