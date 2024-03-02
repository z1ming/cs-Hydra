# MySQL 锁高频面试题

## MySQL 锁是怎么产生的？如何解决？

假设有两个事务 A 和 B[1]：

1. A: `select * from table where id = 1;`，事务 A 对 id = 1 的记录上了 X 行锁
2. B: `select * from table where id = 3;`，事务 B 对 id = 3 的记录上了 X 行锁
3. A: `select * from table where id = 3;`，事务 A 等待 B 释放 id = 3 的锁后再加锁
4. B: `select * from table where id = 1;`，事务 B 等待 A 释放 id = 1 的锁后再加锁

这时事务 A 和事务 B 都在等待对方释放锁后才能加锁，陷入无限等待，这时进入死锁。

解决办法是加锁时添加顺序，比如必须按照这个顺序：先获取 id = 1 的锁，再获取 id = 3 的锁。那么以上步骤就会变为 1，2，4，3，这时就不会发生死锁了。

## InnoDB 有哪些锁？

InnoDB 有如下锁[2]：

- Shared(S) and Exclusive(X) Locks 
- Intention Locks
- Record Locks
- Gap Locks
- Next-Key Locks
- Insert Intention Locks
- AUTO-INC Locks
- Predicate Locks for Spatial Indexes

具体参考[InnoDB 锁介绍](./InnoDB-Locking.md)

小林版锁：

- 全局锁
	- FTWRL
- 表锁
	- 表锁
	- 元数据锁
	- 意向锁
	- AUTO-INC 锁
- 行锁
	- Record Lock
	- Gap Lock
	- Next-Key Lock

## 可重复读隔离级别下，以下 SQL 会发生什么？

```
(id, no, name, age, score)

(15, S0001, Bob, 25, 34)
(18, S0002, Alice, 24, 77)
(20, S0003, Jim, 24, 5)
(30, S0004, Eric, 23, 91)
(37, S0005, Tom, 22, 22)
(49, S0006, Tom, 25, 83)
(50, S0007, Rose, 23, 89)

事务A: 
time1: update students set score=100 where id = 25
time3: insert into students(id,no,name,age,score) value (25,'S0025','sony',28,90)

事务B: 
time2: update students set score=100 where id = 26
time4: insert into students(id,no,name,age,score) value (26,'S0026','ace',28,90)
```

解答:

- time1: 事务 A 获取 Gap Lock（间隙锁），范围 20~30
- time2: 事务 B 获取 Gap Lock（间隙锁），范围 20~30，间隙锁可以共存[3]
- time3: 事务 A 生成 Insert Intention Lock（插入意向锁），等待
- time4: 事务 B 生成 Insert Intention Lock（插入意向锁），等待
- 由于双方都在等待对方释放间隙锁，进入死锁

## 参考

1. 《MySQL 是怎样运行的：从根儿上理解 MySQL》第 22 章 第 6 节
2. [MySQL 8.0 Reference Manual-17.7.1](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)
3. [小林 coding](https://www.xiaolincoding.com/mysql/lock/show_lock.html#time-2-%E9%98%B6%E6%AE%B5%E5%8A%A0%E9%94%81%E5%88%86%E6%9E%90)

