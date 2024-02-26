# 分布式全局唯一 ID 实现

## 一、Twitter Snowflake 算法思想

采用 64 位长整形实现，划分为四个区域：[1]

- 符号，占用 1 位
- 毫秒级时间戳，占用 41 位
- 5 位数据中心 + 5 位机器 ID，占用 10 位
- 序列号，占用 12 位

```
 +------+--------------------+-----------------------+----------+
 | sign | delta milliSeconds | dataCenter | workerId | sequence |
 +------+--------------------+-----------------------+----------+
   1bit        41bits                  10bits           12bits
```

[Java/Go 实现](https://github.com/z1ming/Twitter-Snowflake-implementation)

## 二、美团 Leaf 方案

Leaf这个名字是来自德国哲学家、数学家莱布尼茨的一句话： >There are no two identical leaves in the world > “世界上没有两片相同的树叶”[2]。目前Leaf的性能在4C8G的机器上QPS能压测到近5万/s，TP999 1ms，已经能够满足大部分的业务的需求。

Leaf 方案分为两种：

- Leaf-segment：利用数据库自增 ID 特性，每个机器获取不同分段，然后在各自分段范围内自增。如机器 A 1～1000，机器 B 1001～2000，机器 C 2001～3000
- ：基于雪花算法思想，更适用于订单 ID，避免了 Leaf-segment 可以推算出订单量的问题。

### Leaf-segment 

以下是美团官方架构图：

![Leaf-segment](../images/文章/id/Leaf-segment.png)

主要思想在服务层和 DB Server 之间添加 Proxy Server，Proxy Server 的作用是批量从 DB Server 获取 segment（大小为 step）号段的值，用完之后再去数据库获取新的号段。这样能极大减小请求 DB 的次数，读写数据库的频率用 `1` 变成了 `1 / step`。

**DB Server 表设计**

|    Field    | Type         | Null | Key | Default           | Extra                       |
|:-----------:|:------------:|:----:|:---:|:-------:|:-----:|
|   biz_tag   | varchar(128) | NO   | PRI |                   |                             |
|   max_id    | bigint(20)   | NO   |     | 1                 |                             |
|    step     | int(11)      | NO   |     | NULL              |                             |
|    desc     | varchar(256) | YES  |     | NULL              |                             |
| update_time | timestamp    | NO   |     | CURRENT_TIMESTAMP | on update CURRENT_TIMESTAMP |

biz_tag 用来区分业务，max_id 代表当前业务号段的最大 ID，step 代表步长，比如一次取 1000 个 ID，step = 1000。

**更新号段的 SQL**

如上图，假如 `biz_tag` 为 `test` 的业务在 Leaf1 的号段是 1～1000，当这个号段用完了，但 Leaf2、Leaf3 都没有更新，这时会从 DB Server 请求更新号段为 3001～4000，DB Server 中的 `max_id` 会更新为 4000，更新的 SQL 如下： 

```sql
Begin
UPDATE table SET max_id = max_id + step WHERE biz_tag = xxx
SELECT tag, max_id, step FROM table WHERE biz_tag = xxx
Commit
```

**优点**

- Leaf 服务可以很方便地线程拓展
- ID 是 8 bit 64位数字，符合主键要求
- 容灾性高：Leaf 服务内部有缓存，即使 DB Server 挂掉了，Leaf 服务在号码段用尽前仍然能继续提供服务
- 可以自定义 max_id 大小，方便服务迁移

**缺点**
- ID 不够随机，有泄露 ID 数量的风险
- 号段用完时更新 DB Server 仍然会阻塞线程，tg999 会出现突刺
- DB Server 挂掉时整个系统会不可用

为了解决取号段时阻塞的问题，美团使用了双 Buffer 的方式，即缓存两个号段。在 Leaf 内部提前获取号段，比如当号段使用率达到 10% 时就获取下一号段并缓存，循环往复。具体可查看文末的原文。

### Leaf-snowflake

snowflake 的思路也是基于推特的雪花算法，在此基础改进得来，我们称之为类雪花算法。

![Leaf-snowflake](../images/文章/id/Leaf-snowflake.png)

改算法主要解决了两个问题：

- 弱依赖 Zookeeper：workerId 由 Zookeeper 生成，为了避免强依赖第三方组件，在机器内部缓存了 workerId 文件，这样即使 ZK 挂了也能正常获取到 workerId
- 时钟回拨问题：雪花算法涉及到时间戳的生成，如果机器时钟回拨，则可能生成重复 ID。所以在 Leaf 服务启动前，增加时钟回拨校验，如果确实发生了回拨，则 Leaf 服务启动失败

## 三、百度 UidGenerator



## 参考

1. [Twitter Snowflake](https://github.com/twitter-archive/snowflake)
2. [Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)
2. [百度 UidGenerator](https://github.com/baidu/uid-generator/blob/master/README.zh_cn.md)
3. [滴滴 Tinyid](https://github.com/didi/tinyid)

