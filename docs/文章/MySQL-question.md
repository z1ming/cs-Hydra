# MySQL 高频面试题

## MySQL 索引

### 聊聊 MySQL 的索引结构，为什么使用 B+ 树而不是 B 树？

- B+ 树的非叶子节点不存放实际的数据，只存放索引，因此数据量相同的条件下，相比既存数据又存索引的 B 树，B+ 树可以存放更多的索引，B+ 树可以比 B 树更矮胖，查询底层节点的磁盘 IO 次数会变少
- B+ 树有大量冗余节点，所有非叶子节点都是冗余索引，这些冗余索引让 B+ 树在插入、删除时的效率更高，相比 B 树有更少的结构变换
- B+ 树叶子节点之间用链表连接，有利于范围查询，而 B 树要实现范围查询，只能遍历树，效率更低。

### 你是怎么建立索引的？一般给哪些字段建立索引？

- 字段有唯一性限制的，比如商品编码
- 经常使用 `WHERE` 查询条件的字段，如果查询字段不是一个，可以建立联合索引
- 建立联合索引时，按照最左匹配原则，且稀疏度越大的越放在左边
- 经常用 `GROUP BY` 的字段，这样查询的时候数据就是排序好的，因为建立索引后数据在 B+ 树中的记录是排序好的

### 如何确定语句是否走了索引？

在 SQL 前添加 `explain` 查看执行计划：

- possible_keys： 可能用到的索引
- key：实际用到的索引，如果为 null，表示没有走索引
- key_len：索引的长度
- rows：扫描的数据行数
- type：扫描的数据类型，如果为 all，表示没有走索引，进行了全表扫描

### 什么是联合索引？

联合索引由多个字段组合而成。在 B+ 树中的非叶子节点中保存了联合索引的多个字段。联合查询时，先按照第一个索引字段查询，相同的话按第二个字段查询，以此类推。存在**最左匹配原则**。

### 如果要建立联合索引，字段的顺序有什么需要注意的吗？

- 最左匹配原则
- 稀疏度高的放在左边：稀疏度高就是区分度高的字段，比如唯一 id 要放在类型左边

## MySQL 事务

### MySQL 有哪些隔离级别？可能产生什么问题？

- 读未提交：当一个事务还未提交时，它的变更就能被其他事务看到
- 读已提交：当一个事务被提交后，它的变更就能被其他事务看到
- 可重复读：InnoDB 默认隔离级别，事务执行时看到的数据和启动时一样
- 串行化：会对记录加上读写锁，多个事务对一条记录进行读写操作时，后一个事务需要等待前一个事务执行完成释放锁后才能继续执行

不同的隔离级别可能产生不同的问题：

- 读未提交：脏读、不可重复读、幻读
- 读已提交：不可重复读、幻读
- 可重复读：幻读
- 串行化：无

### InnoDB 如何避免不可重复读？

InnoDB 默认隔离级别是可重复读，可重复读隔离级别在开启事务后，执行第一个 select 语句时，会生成一个 Read View，后面整个事务 select 期间都在用这个 Read View，所以这期间读取的数据都是一致的，不会出现前后读取的数据不一致的问题，避免了不可重复读。

## MySQL 锁

### MySQL 锁是怎么产生的？如何解决？

假设有两个事务 A 和 B[1]：

1. A: `select * from table where id = 1;`，事务 A 对 id = 1 的记录上了 X 行锁
2. B: `select * from table where id = 3;`，事务 B 对 id = 3 的记录上了 X 行锁
3. A: `select * from table where id = 3;`，事务 A 等待 B 释放 id = 3 的锁后再加锁
4. B: `select * from table where id = 1;`，事务 B 等待 A 释放 id = 1 的锁后再加锁

这时事务 A 和事务 B 都在等待对方释放锁后才能加锁，陷入无限等待，这时进入死锁。

解决办法是加锁时添加顺序，比如必须按照这个顺序：先获取 id = 1 的锁，再获取 id = 3 的锁。那么以上步骤就会变为 1，2，4，3，这时就不会发生死锁了。

### InnoDB 有哪些锁？

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

### 可重复读隔离级别下，以下 SQL 会发生什么？

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

