# InnoDB 锁介绍

## 前言

个人学习 MySQL 锁时，网上的文章基本都是按照不同类型做了不同分类，很难记住，原因主要有下：

- 还是没有理解为什么会出现这些锁，这些锁解决了什么问题
- 网络上的文章大多是你参考我，我参考你，每个人加入自己的理解，导致主观的成分太多

本着查找根源的思路，我在 MySQL 的官方文档[1]上找到了最原始对于 InnoDB 的讲解，以下是对该文档的笔记和理解，仅供参考学习之用。

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

### 记录锁(Record Locks)

记录锁是索引记录上的锁。例如 `select c1 from t where c1 = 10 for update;`，这个 SQL 可以防止任何其他事务插入，更新或删除 `t.c1=10` 的行。

记录锁锁定的是索引记录，如果表定义为没有索引，InnoDB 创建一个隐藏的聚集索引[2]来锁定记录。

使用 `SHOW ENGINE INNODB STATUS` 输出类似数据：

```
RECORD LOCKS space id 58 page no 3 n bits 72 index `PRIMARY` of table `test`.`t`
trx id 10078 lock_mode X locks rec but not gap
Record lock, heap no 2 PHYSICAL RECORD: n_fields 3; compact format; info bits 0
 0: len 4; hex 8000000a; asc     ;;
 1: len 6; hex 00000000274f; asc     'O;;
 2: len 7; hex b60000019d0110; asc        ;;
```

### 间隙锁(Gap Locks)

间隙锁是对索引之间，或第一个索引之前，或最后一个索引之后间隙的锁定。如 `select c1 from t where c1 between 10 and 20 for update;`，这时要插入 `t.c1=15` 的值，由于处于间隙区间，因此会阻塞。

间隙可能跨越单个、多个、空的索引值。

如果通过唯一索引搜索某一行的数据，这时不需要用到间隙锁。如果搜索条件包含多个唯一索引列的某几个列，这时还是会发生间隙锁定的。例如对于以下语句：

```
select * from child where id = 100;
```

如果 id 是唯一索引，则不会发生间隙锁定；如果 id 不是索引、或 id 不是唯一的索引，则会间隙锁定 id 前面的间隙。

不同事务可以在间隙上持有冲突锁。比如事务 A 在某个间隙上持有共享间隙锁（S 锁），而事务 B 在同一间隙上持有排他间隙锁（X 锁）。允许冲突间隙锁的原因是，如果从索引中清楚记录，则必须合并不同事务在该记录上持有的间隙锁。

InnoDB 的间隙锁是“纯粹抑制性的”，它们的唯一目的是防止其他事务插入到间隙中。间隙锁可以共存，即一个事务获取的间隙锁不会阻止另一个事务在同一间隙上获取间隙锁。共享间隙锁和排他间隙锁一样，功能也相同，彼此之间不冲突。

### 临键锁(Next-Key Locks)

Next-Key 锁是索引记录上的纪录锁和索引记录之前的间隙上的间隙锁的结合。

假设索引值为 10、11、13 和 20，则该索引的 Next-Key 锁涵盖以下区间：

```
(-♾️, 10]
(10, 11]
(11, 13]
(13, 20]
(20, +♾️)
```

最后一个区间 `(20, +♾️)`，正无穷不是实际值，所以上界不是方括号。

使用 `SHOW ENGINE INNODB STATUS` 输出类似数据：

```
RECORD LOCKS space id 58 page no 3 n bits 72 index `PRIMARY` of table `test`.`t`
trx id 10080 lock_mode X
Record lock, heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0
 0: len 8; hex 73757072656d756d; asc supremum;;

Record lock, heap no 2 PHYSICAL RECORD: n_fields 3; compact format; info bits 0
 0: len 4; hex 8000000a; asc     ;;
 1: len 6; hex 00000000274f; asc     'O;;
 2: len 7; hex b60000019d0110; asc        ;;
```

### 插入意向锁(Insert Intention Locks)

插入意向锁是一种间隙锁，如果插入同一间隙的多个事务没有插入间隙内的同一位置，则无需相互等待。比如有索引 4 和 7，有两个事务分别想要插入值 5 和 6，那么：

- 这两个事务分别使用插入意向锁锁定 4 和 7 之间的间隙，不会发生阻塞，因为不冲突
- 这两个事务单独各自获取插入行上的排他锁

假如客户端 A 创建一个包含两条索引记录的（90 和 102）的表，然后启动一个事务，对 ID 大于 100 的索引记录放置排他锁，排他锁包含 102 之前的间隙锁：

```
mysql> CREATE TABLE child (id int(11) NOT NULL, PRIMARY KEY(id)) ENGINE=InnoDB;
mysql> INSERT INTO child (id) values (90),(102);

mysql> START TRANSACTION;
mysql> SELECT * FROM child WHERE id > 100 FOR UPDATE;
+-----+
| id  |
+-----+
| 102 |
+-----+
```

客户端 B 开始事务以将记录插入到间隙中。事务在等待获取排它锁时获取插入意向锁。

```
mysql> START TRANSACTION;
mysql> INSERT INTO child (id) VALUES (101);
```

使用 `SHOW ENGINE INNODB STATUS` 输出如下内容：

```
RECORD LOCKS space id 31 page no 3 n bits 72 index `PRIMARY` of table `test`.`child`
trx id 8731 lock_mode X locks gap before rec insert intention waiting
Record lock, heap no 3 PHYSICAL RECORD: n_fields 3; compact format; info bits 0
 0: len 4; hex 80000066; asc    f;;
 1: len 6; hex 000000002215; asc     " ;;
 2: len 7; hex 9000000172011c; asc     r  ;;...
```

### 自增锁(AUTO-INC Locks)

如果一个表拥有属性是 `AUTO_INCREMENT` 的列（比如主键 ID），则一个事务插入数据时会添加自增锁，其他事务必须等待上一个事务插入表数据完成是才能继续。

### 空间索引锁(Predicate Locks for Spatial Indexes)

空间索引锁是为了支持 SPATIAL 索引产生的，SPATIAL 是一个空间索引，而 Next-Key 锁无法判断在空间上的间隙，所以空间索引锁可以解决这个问题。 

## 总结

数据库强大而复杂，很难想象 InnoDB 中的锁机制是怎么想到并设计出来的，这可能就是普通人和大佬的差距吧，共勉～

## 参考

1. MySQL 官方文档：https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html
2. 聚集索引和二级索引：https://dev.mysql.com/doc/refman/8.0/en/innodb-index-types.html
