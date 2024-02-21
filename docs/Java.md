# 计算机网络
## http/https
- Http1.0，1.1，2 区别
  ○ Http1.0 发送一次数据建立一次链接，http1.1 建立连接后可以多次发送数据，http2 使用头部压缩，多路复用，服务器推送等新特性
● http 和 https 区别
  ○ Https 使用对称/非对称加密，非对称加密用于证书验证，对称加密用来传输数据
● 对称加密和非对称加密
  ○ A给B发送消息，A使用 B 的公钥加密消息，用 A 的私钥签名，B 使用B 的私钥解密，使用 A 的公钥验签
● https 传输过程
  ○  客户端发起 https 请求
  ○ 服务端返回证书
  ○ 客户端验证证书，本地生成对称加密算法的随机数，并公钥加密传给服务端
  ○ 服务端用私钥解密得到随机数，之后双方使用对称加密进行加解密
● 为什么需要证书？
  ○ 防止中间人攻击，为网站提供身份证明
● 什么是中间人攻击
  ○ 客户端发起请求，获取到中间人的证书，与中间人传输数据时，中间人先服务端加密传输，再返回给客户端
● HTTP 报文
  ○ 请求报文
    ■ 请求行：方法 URL 版本
    ■ 请求头：
      ● Accept-Encoding:压缩编码
      ● Accept-Charset 字符集
      ● Accept-Language 语言
      ● Accept 数据类型
      ● Host 域名
      ● Connection: 
        ○ close，发送完请求的文档就可释放连接
        ○ keepalive：一个http产生的TCP连接在传送完最后一个响应后，还需要保持keepalive_timeout 时间后才开始关闭这个连接
      ● User-Agent：浏览器版本
      ● Referer：目前 url
      ● Content-type：发送信息格式
    ■ 请求主体
  ○ 响应报文
    ■ 响应行：HTTP的版本、状态码、以及解释状态码的简单短语
    ■ 响应头：
      ● Allow：GET, POST, HEAD, OPTIONS
      ● Content-Encoding：gzip
      ● Content-Type：application/json
      ● Date：Wed, 31 Jan 2024 09:29:39 GMT
      ● Server：nginx
    ■ 响应主体
● Thrift和 http 区别
  ○ Thrift二进制，http是文本
  ○ Thrift支持多种数据类型，http 仅只是文本，如 xml，json
TCP/IP
● 三次握手
  ○ 客户端发送SYN，初始序列号 x，进入 SYN_SEND状态
  ○ 服务端收到请求后，发送 SYN+AXK 标志位的包，初始序列号y，ackseq=x+1，进入 SYN_RECV 状态
  ○ 客户端发送 ACK标志为的包，ackseq = y + 1，进入ESTABLISHED 状态
● 四次挥手
  ○ 客户端发送 FIN 标志位的报文，进入 FIN-WAIT-1状态
  ○ 服务端收到 FIN 报文后，回复 ACK 标志的报文，进入 CLOSE-WAIT 状态
  ○ 服务端处理完所有数据后，发送 FIN 标志位报文，进入 LAST-ACK 状态
  ○ 客户端收到 FIN报文后，回应 ACK 报文，进入 TIME-WAIT 状态，等待 2MSL（最大报文生存时间） 时间后，进入 CLOSE 状态
● 为什么等待 2MSL？
  ○ 确保上一个连接的数据包都在网络中消失
● TCP的拥塞控制和慢启动？
  a. 慢启动（Slow Start）：
    ■ TCP连接刚开始时，拥塞窗口（Congestion Window，CWND）被初始化为一个较小的值，通常为1个报文段的大小。
    ■ 在慢启动阶段，每当收到一个确认（ACK）时，CWND的大小就会加倍。这意味着，每个往返时间（Round-Trip Time，RTT）内，发送方可以发送的数据量会指数增长。
    ■ 发送方持续增加CWND，直到达到一个阈值（拥塞避免阈值），或者发生超时。
  b. 拥塞避免（Congestion Avoidance）：
    ■ 一旦CWND达到拥塞避免阈值（通常为拥塞窗口的一半），TCP进入拥塞避免阶段。
    ■ 在拥塞避免阶段，CWND的增长速率变为线性增长，每收到一个确认就增加一个报文段的大小。
    ■ 这种线性增长的方式可以更稳健地探测网络的可用带宽，避免引起网络拥塞。
  c. 快速重传（Fast Retransmit）：
    ■ 当发送方连续收到重复的确认时，它会认为某个报文段丢失，并立即重传该报文段，而不是等待超时。
    ■ 这个机制可以更快地发现数据包的丢失，并加快恢复速度，从而减少对网络的影响。
  d. 快速恢复（Fast Recovery）：
    ■ 在快速重传后，发送方会将拥塞窗口设置为拥塞避免阈值的一半，并继续线性增长。
    ■ 此时，并不会像慢启动那样重新开始，而是进入快速恢复状态，维持当前的发送速率，直到发生超时或收到新的确认。
● 滑动窗口的理解
  ○ RTT：发送到收到 ACK 的时间
  ○ RTO：重传的时间间隔
  ○ Cwnd：拥塞窗口
  ○ MSS：最大报文段长度
  ○ flight: 已发送但尚未收到的数据包数量
  ○ cwnd < ssthresh，慢启动
  ○ Cond > qsthresh, 拥塞避免
● IP 报文内容
  ○ 
● websocket？
  ○ 双向通信协议
  ○ var ws = new WebSocket('ws://localhost:8080');
osi七层模型
● OSI 七层模型？
  ○ 物理层
  ○ 数据链路层
  ○ 网络层：路由器属于哪层？
  ○ 传输层：TCP/UDP 属于哪层？
  ○ 会话层：会话管理，身份验证和授权
  ○ 表示层：二进制
  ○ 应用层：FTP SMTP HTTPS TELNET
● TCP/IP 四层
  ○ 数据链路层
  ○ 网络层
  ○ 传输层
  ○ 应用层
● 五层
  ○ 物理层
  ○ 数据链路层
  ○ 网络层
  ○ 传输层
  ○ 应用层
操作系统
进程和线程
● 进程间的通信方式包括：
  ○ 管道：管道是一种半双工的通信方式，用于具有亲缘关系的进程间通信，通常是父子进程或兄弟进程之间使用。数据只能单向流动，且具有一定的容量限制。应用场景包括父子进程之间的通信或者在Shell脚本中使用。
  ○ 命名管道：命名管道也是一种半双工的通信方式，但可以在无亲缘关系的进程间进行通信。命名管道在文件系统中有对应的文件名，不受进程关系限制。应用场景包括进程间的数据交换、进程间通信等。
  ○ 信号：信号是一种异步的通信方式，用于通知接收进程发生了某种事件。常见的信号有SIGKILL、SIGTERM等。应用场景包括进程终止通知、进程状态变化通知等。
  ○ 消息队列：消息队列是一种通过消息实现进程间通信的方式，允许多个进程向同一个队列发送消息或接收消息。消息队列可以实现进程之间的异步通信，提供了一种高效的通信方式。应用场景包括进程解耦、进程间数据交换等。
  ○ 共享内存：共享内存是一种高效的进程间通信方式，多个进程可以访问同一块内存区域。可以实现进程间数据共享，但需要考虑同步和互斥问题。应用场景包括大数据量的数据交换、共享缓存等。
  ○ 信号量：信号量是一种用于进程间同步和互斥的机制，可以控制多个进程对共享资源的访问。通过信号量可以实现进程间的同步和互斥操作。应用场景包括进程同步、资源互斥访问等。
  ○ Socket通信：Socket通信是一种全双工的通信方式，通过网络套接字（socket）在不同主机上的进程进行通信。可以实现客户端与服务器之间的通信，包括TCP和UDP两种协议。应用场景包括网络编程、客户端与服务器之间的数据传输、实时通讯等。
● 线程如何通信
  ○ 共享内存： synchronized volatile Lock
  ○ 消息传递： join() notify() countDownLatch,CyclicBarrier
  ○ 管道 ：pipedReadStream
● 有哪些进程调度算法？
  a. 先来先服务调度算法，缺点是长时间进程耽误短时间进程
  b. 最短优先作业调度算法，缺点是长时间进程一直轮不上
  c. 高响应比优先调度算法，优先级=(等待时间+要求时间)/要求时间，优先级高的先执行，平衡了短时间进程和长时间进程
  d. 时间片轮转调度算法，每个进程分配一个时间片，时间片内结束，立即切换进程；否则进行下一个进程，保证了公平
  e. 最高优先级调度算法
    ⅰ. 静态优先级：进程创建时就指定了优先
    ⅱ. 动态优先级：进程在执行时动态调整优先级
    ⅲ. 抢占式：就绪队列中出现优先级高的进程，则切换至优先级高的进程
    ⅳ. 非抢占式：就绪队列中出现优先级高的进程，当前进程执行完后再执行优先级高的进程
  f. 多级反馈队列调度算法
    ⅰ. 多个队列，优先级从高到低，时间片从小到大
    ⅱ. 结合了时间片轮转算法和最高优先级调度算法
● 线程切换的过程
  ○ 保存当前线程上下文
  ○ 切换到调度器
  ○ 恢复下一个线程上下文
  ○ 切换到新线程
● 什么是分支预测？
  ○ 用于提高处理器执行分支指令时的效率。执行分支指令时，对结果进行预测，包括 1 位预测，2位预测。如果预测错误，重新预测，预测正确，继续执行。
● poll 和 epoll 区别
  ○ 都是用于I/O多路复用的机制
  ○ 事件分发器，epoll_create，epoll_ctl,epoll_wait
  ○ Select和poll 都轮询注册的文件描述符，epoll 基于回调
  ○ Select 限制最大连接数1024，poll无限制；select 可跨系统，poll 不可
● 分段和分页的区别
  ○ 单位性质。页是信息的物理单位，分页的目的是为了实现离散分配方式，以消减内存的外零头，提高内存的利用率，主要满足系统管理的需要；段是信息的逻辑单位，段的大小不固定，由用户编写的程序决定，主要为了满足用户的需要。
  ○ 大小固定性。页的大小固定且由系统决定，页只能以页大小的整数倍地址开始；段的大小不固定，由用户在编程时确定，或由编译程序在对源程序进行编译时根据信息的性质来划分。
  ○ 地址空间维度。分页的作业地址空间是一维的，即单一的线性地址空间，程序员只需利用一个记忆符，即可表示一个地址；分段的作业地址空间是二维的，程序员在标识一个地址时，既需给出段名，又需给出段内地址。	
Linux
命令
● 常用命令
  ○  top、netstat(监控TCP/IP网络)、grep、sed、awk
● sed和awk有什么区别？
  ○ 都是用于文本处理命令的工具，区别在于：
  ○ sed，主要用于对文本进行替换、删除、插入等操作。它适合对整行文本进行处理，可以通过正则表达式匹配文本进行操作。
  ○ awk，可以实现更复杂的文本处理逻辑，包括对字段的操作、条件判断、循环等。它适合处理结构化的文本数据，可以按列对数据进行处理。
  ○ sed适合简单的文本替换和编辑操作，而awk适合处理结构化的文本数据并实现更复杂的处理逻辑
● 如何查看进程，按 CPU 占用排序，按内存排序？
  ○ top按P(cpu)，按M(内存)，H线程
● ps aux 和 ps -elf 的区别？
  ○ aux 表示所有关联到终端的进程
  ○ elf 表示列出所有进程
● wk 和 grep 的区别是？
  ○ awk 是可以对文本和数据进行处理，也是一种编程语言
  ○ grep 仅支持搜索，不能处理
    ■ -i 忽略大小写
    ■ -v 反向匹配
    ■ -r 递归搜索
    ■ -n 显示行号
MySQL
索引
● MySQL 有哪些日志？
  ○ 慢查询日志：记录执行时间超过阈值的SQL查询语句。用于优化查询性能，找出慢查询并进行优化。
  ○ redo log：用于InnoDB存储引擎的崩溃恢复。
  ○ binlog：记录所有对数据的更改操作的日志。它包含了对数据库进行插入、更新和删除操作的详细信息，以二进制形式记录，可以用于数据的备份和恢复、主从复制等场景
  ○ Undo Log：实现事务的原子性，记录事务对数据的修改操作，方便在事务失败或回滚时恢复数据到事务前的状态，确保数据的一致性。支持事务的回滚和MVCC（多版本并发控制）机制。
● 聚簇索引与非聚簇索引
  ○ 在 InnoDB 中，非聚簇索引就是二级索引，两者用的都是 B+树，聚簇索引是一种特殊的索引，它的叶子节点包含了整个数据行，而非聚簇索引的叶子节点只包含相应行的索引值。
  ○ 聚簇索引只能有一个，由主键构建的；非聚簇索引可以有多个
● 不同引擎对索引的支持
  ○ InnoDB 支持聚簇索引和非局促索引，MyISAM 仅支持非聚簇索引
● B+树树高怎么算
  ○ h=log_m(N) 表示以 m 为底 N 的对数，
● B+树高为4能支持多少数据量
  ○ InnoDB 页的大小默认是 16k
  ○ 假设主键 bigInt，主键+指针=8+6=14 字节，度是 16*1024/14=1170，一条数据 1k
  ○ 1170^3*16~百亿
● B 树和 B+树的区别
  ○ B树中同一键值不会出现多次,并且它有可能出现在叶结点,也有可能出现在非叶结点中。
  ○ 而B+树的键一定会出现在叶结点中,并且有可能在非叶结点中也有可能重复出现,以维持B+树的平衡。
● InnoDB 和 MyISAM区别
  ○ MyISAM 不支持事务和外键
● Binlog 记录的是什么
  ○ log_name
  ○ pos
  ○ event_type
  ○ server_id
  ○ end_log_pos
  ○ info
● 主键的含义？
  ○ 唯一性
  ○ 非空性
  ○ 索引性：主键列通常会自动创建一个索引
  ○ 数据完整性：防止重复数据插入，
● 事务有哪些特性？
  ○ 原子性：要么全部成功，要么全部失败
  ○ 隔离性：多个并发事务之间不会交叉影响
  ○ 持久性：事务处理结束后，数据修改是永久的
  ○ 一致性：事务前后的数据完整性
● 隔离性的级别有哪些？
  ○ 读未提交（隔离级别最低）：脏读+不可重复读+幻读
  ○ 读提交：公司默认，不可重复读+幻读
  ○ 可重复读，InnoDB 默认，幻读
  ○ 串行化 （隔离级别最高）：会对记录加读写锁，多个事务修改同一个数据，一个事务需要等另一个执行完才能执行
● 不同隔离级别可能发生的情况有什么？
  ○ 读未提交：脏读，不可重复读，幻读
  ○ 读提交：不可重复读，幻读
  ○ 可重复读：幻读
  ○ 串行化：无，但是会影响性能
● 数据隔离级别可能产生的问题
  ○ 脏读：事务 T1 修改了数据，在 commit 或 rollback 前，事务 T2 读到了修改后的数据，如果 T1rollback，T2 读到了未曾改变的数据
  ○ 不可重复读：事务 T1 读取了数据，事务 T2 修改或删除了数据，如果 T1 试图重新读取数据，两次的结果不一致，即不可重复读
  ○ 幻读：事务 T1 查询范围内的数据，事务 T2 新增了范围内的数据，T1 重复读时，读到的数据不一致
● 游标
  ○ declare 游标名 cursor for 数据
  ○ 打开游标：open 游标名
  ○ 遍历游标： fetch 游标名 into 变量名
  ○ 关闭游标： close 游标名
  ○ 可以对数据集的没一行进行操作
● 脏读、不可重复读、幻读的区别？
  ○ 脏读：两次查询同一数据时，读到了未提交事务的数据
  ○ 不可重复读：两次查询同一数据时，期间提交了另一个事务，修改了数据——加行锁解决
  ○ 幻读：两次查询同一数据时，期间提交了另一个事务，增加或删除了范围内的数据，需要加间隙锁
● 间隙锁的原理？
  ○ 解决了不可重复读下的幻读现象，属于 InnoDB 中的行锁。有两个场景会出现间隙锁：
  ○ 有 ID 为 0～101 的数据，如果查询 id > 100，除了对 101 的数据加行锁，还会对大于 101 的间隙加间隙锁，这时如果插入 id=102，会阻塞
  ○ 插入一条不存在的记录，比如 102，其他 session 插入 201 的记录，也会出现锁等待
● 事务的状态有哪些？
  ○ 活动、部分提交、中止、失败、提交
● MySQL 如何保证原子性？
  ○ 通过 undo log 保证，提交事务前，将数据保存在 undo log，回滚后利用 undo log 恢复
● 如何通过 undo log 撤销数据？
  ○ 插入、更新、删除，把主键值、旧记录、记录值都记录下来，回滚时进行相应回滚
● 建立索引的原则？
  ○ 唯一 ID
  ○ 用在 where、order by、group by 里的
● 什么情况下不用建立索引？
  ○ 经常更新的
  ○ 表中存在大量重复数据
  ○ 表数据太少时
● 最左匹配是什么？例子？
  ○ abc联合索引，a，ab，abc 都可走索引，b，c 都不走索引
● abc 联合索引，bc 走索引吗？
  ○ 如果 select 【字段】，字段为 abc、b、c...都会走索引，因为只查索引就能找到所要的字段，否则不会走
● 什么是索引下推ICP？
  ○ 是一种数据库查询优化技术，在执行查询时，将一部分过滤条件下推到存储引擎层面进行处理，减少了数据库返回的数据量，提高了查询性能。
  ○ 之前通过索引获取到数据，再通过 where 条件筛选，索引下推就是将 where 条件也交给引擎处理，减少回表次数，聚簇索引不存在回表一说
  ○ 如果 where 仅可使用索引中的列，则会下推
  ○ 可用于 InnoDB 和 MyISAM
  ○ 子查询的条件无法下推
● 什么情况下索引会失效？
  ○ 违反最左前缀法则
  ○ 类型转化
  ○ != > <
  ○ like %xx
  ○ or
● not in 会走索引吗？
  ○ 5.7 前，not in 中的子查询不走索引，not in 根据实际情况评估，如果索引+回表效率高，就走索引，如果全表扫描效率高，则不走，8.0后都走
● MySQL有哪些优化方式？
  ○ 针对慢SQL可以进SQL优化，优化的方式可以建立联合索引进行索引覆盖优化，减少回表；对排序字段增加索引避免 file sort的问题；使用小表驱动大；or改成union；大偏移量的limit，先通过>id找到第一页的 id，再limit。
  ○ 针对热点数据可以构建缓存，查询的时候先查缓存，减少对 mysql 的访问提高查询效率。
  ○ 读请求过大的时候，可以构建 mysql 主从架构，进行读写分离，由多个从机来承接读请求的流量。
  ○ 写请求过大的时候，可以进行分库，由多个 mysql 机器来承接写请求的流量。
  ○ 客户端可以增加客户端连接池，减少客户端与 mysql 连接的建立和释放的开销，复用连接。
● char和varchar的区别？
  ○ char 定长
  ○ varchar 变长
● 主从复制的过程？
  ○ 主库生成二进制日志：主库在每个事务更新数据完成之前，会在二进制日志中记录这些改变。
  ○ 从库复制中继日志：从库将主库的这些二进制日志复制到自己的中继日志中。34
  ○ 从库重做中继日志：从库的SQL线程从中继日志中读取事件，并执行这些事件以更新从库的数据，使其与主库的数据保持一致。123
设计模式
● 常见的设计模式有？
  ○ 单例
  ○ 迭代器
  ○ 观察者
  ○ 模板方法
  ○ 策略模式
  ○ 适配器模式
  ○ 装饰器
  ○ 工厂方法
  ○ 责任链模式
● 责任链模式常用实现场景有哪些?
  ○ netty, Tomcat Filter 链
  ○ 解决了什么问题？解耦请求的发送者和处理者
Java
垃圾回收
● 垃圾回收机制？
  ○ 分为新生代和老年代
  ○ 新生代分为Eden和serviver区，代码中创建的对象保存在 Eden 和 serviver1，如果超过容量，触发Minor GC，剩余对象放到 Serviver2
  ○ 如果剩余对象大于 serviver2，直接放到老年代
  ○ Serviver2回收超过 15 次还存活的对象放到老年代
  ○ 还可设置为动态次数，规则是当前 serviver2 超过 50% 的对象时？中？超过这些对象的最大年龄，会直接到老年代
  ○ minorGC前，会先检查老年代最大连续空间是否大于新生代所有对象总空间。如果不成立，检查 handlePromotionFailure 这个参数，Jdk6之前，有一个担保参数handlePromotionFailure，
  ○ 如果允许，则用老年代最大连续可用空间和历次晋升到老年代的平均大小对比，如果大于，则尝试 minorGC，尽管有风险
  ○ 如果为不允许，或小于，则触发 FullGC
  ○ 如果 FullGC 后，仍然放不下，触发 OOM
  ○ Jkd6 update24 之后，参数失效，规则为只要老年代的连续空间大于新生代对象的总大小或历次晋升的平均大小，就进行 minorGC，否则 FullGC
● 垃圾收集算法
  ○ 标记-清除，基础，基于分代收集理论
    ■ 缺点1：执行效率不稳定，执行效率随着对象数量增长而降低
    ■ 缺点 2：内存碎片化
  ○ 标记-复制，回收新生代，IBM说新生代的98% 熬不过第一轮收集，因此不用按照 1:1 分配，HotSpot默认Eden 和 survivor 8:1
    ■ 优点：解决了内存碎片化的问题
    ■ 缺点：可用内存降为原来一半，浪费多
  ○ 标记-整理：基于标记清除，增加内存移动，对象整理为连续，老年代使用
● 垃圾收集器
  ○ CMS(Concurrent Mark Sweep)：大多数采用标记-清除，直到碎片化程度影响对象分配置，触发一次标记-整理
    ■ 缺点 1：并发程序，对资源比较敏感，cpu 跟不上影响性能
    ■ 缺点 2：无法处理浮动垃圾，可能出现concurrentModeFailure导致FullGC
    ■ 缺点 3：碎片化
  ○ G1(Garbage First)：
    ■ 优点，推出 mixedGC，不在基于分代收集的思想，而是任何代有垃圾都可以收集，建立了可预测的停顿时间模型，
    ■ 按照 region 划分，跟踪 region 的价值，优先处理回收价值大的 region
    ■ 低延迟
  ○ Shenandoah
  ○ ZGC 收集器
● 垃圾回收机制？
  ○ Eden Serviver1，Serviver2
  ○ 创建的对象放在 Eden 区和 Serviver 区，如果都快满了，触发 Monitor GC，将没被回收的放到 Serviver2 区
  ○ 躲过 15 次 GC 的对象会转移到老年代，大对象也会转移到老年代
  ○ 15 次可以设置 -XX:MAX TenuringThreshold
  ○ 除了15规则，还有动态年龄判断的规则，如果 Serviver 里的对象容量加起来超过 50%，那么超过这些最大年龄的对象直接进入老年代
  ○ 大对象直接进入老年代，通过参数 -XX:PretenureSizeThreshold 控制
  ○ 如果 Minor GC 后剩余的对象大于 serviver，直接放到老年代
  ○ 如果老年代也不够？看参数-XX:-HandlePromotionFailure是否设置，如果设置且判断失败，或者没有设置，则触发 Full GC
  ○ 如果还是不够，则触发 OOM
● 标记清除中“标记”的过程
  ○ 遍历对象并标记，使用深度优先和广度优先
● 新生代，老年代的回收算法？
  ○ 新生代标记复制，老年代标记清除+标记整理
● 强、弱、软、虚分别都是什么时候用，用给谁，怎么用
  ○ 强引用不会垃圾回收
  ○ 软引用如果内存够不会回收，不够会回收，回收后仍不足，抛 OOM
  ○ 弱引用生命周期更短，一旦发现弱引用对象就会回收
  ○ 虚引用最弱，不能通过 get 方法获取对象，主要用于跟踪垃圾回收的活动，必须和队列一起用
数据结构
● HashMap
  ○ 为什么是线程不安全的？
    ■ Resize 可能获取到空值
    ■ 多线程put可能数据覆盖
  ○ 数据结构？
    ■ Jdk1.8之前数组+链表
    ■ Jdk1.8开始数据+链表+红黑树，链表长度大于8时转为红黑树
  ○ resize
    ■ 数组扩容-每个元素重新 hash，如果是链表，先保存头结点，然后依次计算后续的节点
    ■ Jdk1.8 之前头插法，1.8 开始尾插法
● HashTable 和ConcurrentHashMap 区别
  ○ A 整个方法加锁，B 在每个链表头结点加锁，不会产生锁冲突
  ○ A 通过 synchronized加锁，B 使用 cas，效率更高
  ○ A resize 旧元素搬运到新空间，然后释放旧空间，大量拷贝，效率低； B 每次拷贝一部分，扩容时新旧空间同时存在
  ○ A get 加锁，B get 不加锁，原因是Node 中的 val(和next) 使用 volatile 修饰
● Hashmap扩容过程？
  ○ 判断老表容量是否超过上限，是修改为 Integer.MAX_VALUE
  ○ 否将容量和阈值都修改为原来 2 倍，遍历老数组，如果索引位置有一个节点，直接迁移到新位置，如果大于 1 个，则判断是红黑树节点还是链表节点，然后分别处理
类
● 类加载的五个阶段？
  ○ 加载
  ○ 验证
    ■ 文件格式验证
    ■ 元数据验证
    ■ 字节码验证
    ■ 符号引用验证
  ○ 准备
    ■ 分配变量内存并设置初始值
  ○ 解析
    ■ 符号引用转化为直接引用
  ○ 初始化
    ■ 初始化变量的初始值
●  双亲委派模式？
  ○ 当Java虚拟机需要加载一个类时，它会先委派给当前类加载器的父类加载器去加载。如果父类加载器也无法加载该类，就会依次向上委派，直到达到顶层的启动类加载器（Bootstrap ClassLoader）。如果顶层类加载器也无法加载该类，就会抛出类找不到的异常（ClassNotFoundException）。
  ○ 优势是避免重复加载同一个类，并且可以确保核心Java类库不会被恶意修改或替换。
● Java 异常机制？
  ○ Throwable
  ○ Exception
    ■ RuntimeException
    ■ NullPointerException
  ○ Error
    ■ OutOfMemoryError
● ASCII/Unicode/UTF-8区别？
  ○ ASCII是最初的字符编码方案，
  ○ Unicode是一个字符集，
  ○ 而UTF-8是一种Unicode字符编码方案，用于在计算机系统中表示和传输文本。
线程
● 有哪些线程状态？
  ○ NEW：线程创建，还未调用 start() 方法	
  ○ RUNNABLE：就绪状态+正在运行
  ○ BLOCKED：等待监视器锁时，陷入阻塞状态
  ○ WAITING：等待另一线程执行其他操作
    ■ Object.wait()
    ■ Thread.join()
    ■ LockSupport.park()
  ○ TIMED_WAITING：具有指定等待时间的等待状态
    ■ Thread.sleep()
    ■ Object.wait() with timeout
    ■ Thread.join with timeout
    ■ LockSupport.parkNanos
    ■ LockSupport.parkUntil
  ○ TERMINATED：线程执行完成，终态
● waiting 状态下的线程如何恢复到 running 状态？
  ○ 等待的线程被其他线程唤醒，notify()，notifyAll()
  ○ LockSupport.unpark()
● sleep和wait 的区别
  ○ Sleep 是 Thread 方法，wait 是 Obect 方法
  ○ sleep 不用和 synchronized 配合使用，wait 需要和 synchronized 一起用
  ○ Sleep 在睡眠同时，不会释放对象锁，wait 在等待时会释放对象锁
● notify, notifyAll 区别？
  ○ notify 唤醒一个线程
  ○ otifyAll 唤醒所有线程，所有线程退出 waiting 状态，开始竞争锁，只有一个线程获取到锁，这个线程执行完后其他线程继续竞争
● notify 选择哪个线程？
  ○ 随机的，具体依赖 jvm 的实现，如果在 hotspot，则是先进先出
● 如何停止运行中的线程？
  ○ Thread.interrupt()，线程内判断 interrunpted，如果为 true 抛出异常
  ○ 先将线程 sleep，然后调用 interrupt 方法
  ○ stop() 暴力终止
  ○ Thread.interrupt()，线程内判断 interrupted，如果为 true 直接 return
● 调用 interrupt 如何让线程抛出异常的？
  ○ 每个线程都有一个 interrupted 状态，默认为 false
  ○ 其他线程执行了 interrupt，如果当前线程处于阻塞状态，则会直接中断，抛出异常
  ○ 否则只做标记，后面根据轮询中断状态判断是否要停止任务
● 如果靠变量停止线程，弊端是什么？
  ○ 缺点是中断不及时，轮询是需要等下下一个循环才能判断出来
多线程
● 说说 volatile？
  ○ 保证可见性
  ○ 不保证原子性
  ○ 禁止指令重排
● 如何保证原子性？
  ○ 使用 synchronized
● synchronized 支持重入吗？如何实现的？
  ○ 重入：synchronized 内部可再次用 synchronized 请求对象锁
  ○ 底层调用 Mutex Lock，每个线程有一个状态和线程 ID
  ○ 线程请求方法时，检查锁状态
  ○ 锁状态是 0，代表锁没有被占用，使用 CAS 获取锁，把线程 ID 替换为自己的线程 ID
  ○ 锁状态不是 0，代表有线程正在访问该方法。如果线程 ID 是自己的线程 ID，如果是可重入锁，status 自增 1，然后获取到锁，执行方法；如果是非重入锁，则进入阻塞状态
  ○ 释放锁时，可重入锁，status 每次减 1，直到为 0，释放该锁；不可重入锁，线程退出方法，直接释放锁
● Sychorinize 用来干什么的？原理是什么？
  ○ 加锁的，分为类锁和对象锁
  ○ 对象加锁，使用 monitor，每个对象就是一个监视器，加锁就是获取该对象的监视器锁，monitorenter 和 moniterexit，方法加锁通过一个标识位，叫 ACC_SYNCHRONIZED
● synchronized 和 Lock 区别
  ○ synchronized 悲观锁，Lock 乐观锁
● Java 线程池有 7 个核心参数有什么？
  ○ corePoolSize 核心线程大小
  ○ maximumPoolSize 最大线程数量
  ○ keepAliveTime 空闲线程存活时间
  ○ unit 存活时间单位
  ○ workQueue 工作队列
  ○ threadFactory 线程工厂
  ○ handler 拒绝策略
● java 线程池调度过程？
  ○ 提交一个任务，先将线程数和 corePoolSize 比较，如果小于，创建线程；如果大于，放在队列中，队列满后，如果 maximumPoolSize > corePoolSize,创建新线程
  ○ 线程执行完，会从队列中取任务执行
  ○ 线程数量达到 maxmumPoolSize，且队列也满时，新提交的任务执行 rejectExecutinoHandler
  ○ 如果线程空闲时间达到 keepAliveTime时，如果 allowCoreThreadTimeOut =true，核心线程永远不会被销毁；如果为 false，则任何时间超过keepAliveTime 的线程都会被销毁
● 线程池什么时候会做等待？什么时候会做丢弃？
  ○ allowCoreThreadTimeOut=true会做等待，allowCoreThreadTimeOut=false 会做丢弃
● ThreadLocal
  ○ ThreadLocal用于创建线程局部变量。每个 ThreadLocal 对象都可以维护一个线程本地变量，可以使线程间的数据隔离，以此来解决多线程同时访问共享变量的安全性。
  ○ get 方法的主要流程为：
    ■ 先获取到当前线程的引用
    ■ 获取当前线程内部的 ThreadLocalMap
    ■ 如果 map 存在，则获取当前 ThreadLocal 对应的 value 值
    ■ 如果 map 不存在或者找不到 value 值，则调用 setInitialValue() 进行初始化
    ■ 其中每个 Thread 的 ThreadLocalMap 以 threadLocal 作为 key，保存自己线程的 value 副本，也就是保存在每个线程中，并没有保存在 ThreadLocal 对象中。
  ○ set 方法的作用是把我们想要存储的 value 给保存进去。set 方法的流程主要是：
    ■ 先获取到当前线程的引用
    ■ 利用这个引用来获取到 ThreadLocalMap
    ■ 如果 map 为空，则去创建一个 ThreadLocalMap
    ■ 如果 map 不为空，就利用 ThreadLocalMap 的 set 方法将 value 添加到 map 中
jvm
● 有用过堆外缓存吗？
  ○ 解决了 GC 频繁的问题，用在数据量大的场景，需要手动回收
  ○ ohc 包可以用，底层调用了 Unsafe 包
● java 内存模型
● 
Redis
数据结构
● 数据结构有哪些？对应应用场景？
  ○ String：(SDS(simple dynamic string))缓存对象，常规计数，分布式锁，共享 session
    ■ SDS 和 C 字符串的区别，为什么不用 C 字符串
    ■ 常数复杂度获取字符串长度
    ■ 杜绝缓冲区溢出
    ■ 减少内存重新分配带来的内存重分配次数
    ■ 二进制安全
    ■ 兼容部分 C 字符串函数
  ○ Hash：(listpack, 哈希表)缓存对象，购物车
  ○ List：(quicklist)消息队列（两个问题：1.生产者自行实现全局唯一 ID 2.不能以消费组形式消费数据）
  ○ Set：(哈希表, 整数集合)聚合计算场景（并集、交集、差集），点赞，共同关注，抽奖活动等
  ○ Zset：(listpack,skiplist)排序场景，排行榜，电话，姓名排序
● 新增的四种数据结构？
  ○ BitMap：2.2版本新增，二值统计场景，签到、用户登陆状态，连续签到用户总数
  ○ HyperLogLog: 2.8版本新增，海量数据基数统计场景，百万级网页 UV 统计
  ○ GEO：3.2版本新增，存储地理位置，滴滴
  ○ Stream: 5.0版本新增，消息队列，相比于 List 实现的消息队列有两个好处：1自动生成全局唯一 ID，2支持以消费组形式消费数据
● Zset 使用了什么数据结构？
  ○ 压缩列表ziplist或跳表。有序集合的元素个数小于 128 个，且每个元素大小小于 64 字节时，使用压缩列表，否则使用跳表。Redis7.0压缩列表已经废弃，改用 listpack 实现
● 跳表 Skiplist 了解吗？
  ○ 是由链表改进过来的，链表的查询复杂度是 O(n)，为了提高查询效率，采用多层结构，查询效率为O(logN)
● 定义 Redis 跳跃表的结构，再实现一版它的插入方法。
public class Skiplist {
    static final int MAX_LEVEL = 32;
    static final double P_FACTOR = 0.25;
    private Random random;
    private int level;
    SkipListNode head;

    public Skiplist() {
        this.head = new SkipListNode(-1, MAX_LEVEL);
        this.level = 0;
        this.random = new Random();
    }

    public boolean search(int target) {
        SkipListNode curr = this.head;
        for (int i = level - 1; i >= 0; i--) {
            while (curr.forwords[i] != null && curr.forwords[i].val < target) {
                curr = curr.forwords[i];
            }
        }
        curr = curr.forwords[0];
        if (curr != null && curr.val == target) {
            return true;
        }
        return false;
    }

    public void add(int num) {
        SkipListNode[] update = new SkipListNode[MAX_LEVEL];
        Arrays.fill(update, head);
        SkipListNode curr = this.head;
        for (int i = level - 1; i >= 0; i--) {
            while (curr.forwords[i] != null && curr.forwords[i].val < num) {
                curr = curr.forwords[i];
            }
            update[i] = curr;
        }
        int lv = randomLevel();
        SkipListNode newNode = new SkipListNode(num, lv);
        for (int i = 0; i < lv; i++) {
            newNode.forwords[i] = update[i].forwords[i];
            update[i].forwords[i] = newNode;
        }
    }

    public boolean erase(int num) {
        SkipListNode[] update = new SkipListNode[MAX_LEVEL];
        SkipListNode curr = this.head;
        for (int i = level - 1; i >= 0; i--) {
            while (curr.forwords[i] != null && curr.forwords[i].val < num) {
                curr = curr.forwords[i];
            }
            update[i] = curr;
        }
        curr = curr.forwords[0];
        if (curr == null || curr.val != num) {
            return false;
        }
        for (int i = 0; i < level; i++) {
            if (update[i].forwords[i] != curr) {
                break;
            }
            update[i].forwords[i] = curr.forwords[i];
        }
        while (level > 1 && head.forwords[level - 1] == null) {
            level--;
        }
        return true;
    }
    private int randomLevel() {
        int level = 1;
        while (random.nextDouble() < P_FACTOR && level < MAX_LEVEL) {
            level++;
        }
        return level;
    }
    class SkipListNode {
        int val;
        SkipListNode[] forwords;

        public SkipListNode(int val, int level) {
            this.val = val;
            this.forwords = new SkipListNode[level];
        }
    }
}

● 什么是压缩列表？为什么在元素少时使用压缩列表？
  ○ 字节数组，压缩列表占用一块连续的内存，如果元素多时，创建和拓展时需要操作更多内存，所以只在元素少时才能提升效率
  ○ zlbytes：占用内存字节数；zltail：尾节点的偏移量；zllen：节点个数；entry[]：具体节点信息；zlend：特殊值，用于标记末端
● 什么是 listPack？listpack 和 ziplist 区别是什么？
  ○ listpack 是 ziplist 优化版本，为了解决 ziplist 的连锁更新问题，连锁更新是因为 ziplist 每个元素保存了上一个元素的长度信息
  ○ listpack 有四个部分：总空间大小+元素个数+每个元素+末尾标识符，每个元素由编码方式+数据+前两部分长度三个部分组成，元素的第三部分长度信息保证了可以倒序遍历
● 为什么 MySQL 不用 SkipList
  ○ B+树三层，跳表可能大于三层，磁盘 io 次数更多
单机数据库
● RDB 和 AOF 的区别
  ○ RDB 是快照，AOF 是日志追加形式
  ○ RDB 可能会数据丢失，AOF 不易丢失，但性能可能受影响，如果可容忍部分丢失，选择 RDB
  ○ 先复制文件，然后删除旧文件
多机数据库
● redis如何保证高可用？
  ○ 主从复制，哨兵模式（实现高可用。主节点挂了可以把从节点转为主节点），集群模式（集群模式，分布式数据库方案。sharding（分片））
● 主从复制流程？实现数据一致性
  ○ 从服务器发送 sync 命令
  ○ 主执行 bgsave 命令，生成 rdb 文件
  ○ 发送 rdb 文件
  ○ 发送缓冲区写命令
  ○ 如果主删除 key，命令传播
  ○ 旧版流程缺点
  ○ 如果主从连接断开，触发同步，主节点会同步全量数据，实际只需同步断开链接时的那些数据
  ○ 新版流程
  ○ 分为完整同步和部分同步
  ○ 初次复制-完整同步，断线重连-部分同步
  ○ 队列实现Offset、复制积压缓冲区
使用场景
● redis 使用场景？
  ○ 分布式锁，缓存，消息队列，对于数据操作都是原子性的
● redis 快的原因？
  ○ 单线程模型，避免线程切换带来的开销
  ○ 在内存中操作，采用高效的数据结构
  ○ I/O 多路复用，select/epoll 机制允许内核中存在多个已连接socket 和监听 socket，一旦有请求到达，就交给 redis 线程处理
  ○ 数据结构丰富
● redis 和 MySQL 如何保证数据一致性
  ○ 先更新数据库，再更新缓存+过期时间
● Redis如何实现分布式锁？用什么命令实现
  ○ setnx expire
  ○ SET key value PX 30000 NX
  ○ redission
  ○ Redlock 
    ■ redis集群解决方案,使用redlock解决:
    ■ 顺序向5个节点请求加锁（5个节点相互独立，没任何关系）
    ■ 根据超时时间来判断是否要跳过该节点
    ■ 如果大于等于3节点加锁成功，并且使用时间小于锁有效期，则加锁成功，否则获取锁失败，解锁
● redis怎么实现限制用户请求的？怎么计数+1的？如果多条线程过来怎么保证线程安全？
  ○ 计数器
  ○ 滑动窗口
  ○ 令牌桶
  ○ Incr 实现+1，原子操作
中间件
RocketMQ
● rocketMq的延迟消息原理？
  ○ 共有 18 个延迟等级，每个等级有一个延迟队列
  ○ broker 在启动时，对于延迟消息，根据daleyTimeLevel 存入特定的队列中，topic也修改 topic 为 schedule_topic_xxxx,每个线程处理某个等级的消息，当消息过期时，放入消费队列中用于消费
● broker里有哪些存储的文件？
  ○ store/commitlog：消息的物理文件
  ○ store/consumequeue 消费队列的索引信息
  ○ store/index 文件 commitLog 的索引信息
  ○ store/checkpoint 记录 CommitLog 和 ConsumeQueue 文件的刷盘点，保证消息的可靠性
NoSQL
● 各个NOSQL之间的区别
  ○ redis：快，数据类型丰富
  ○ Hbase：
    ■ 存储容量大，一个表可以容纳上亿行，上百万列；
    ■ 可通过版本进行检索，能搜到所需的历史版本数据；
    ■ 负载高时，可通过简单的添加机器来实现水平切分扩展，跟Hadoop的无缝集成保障了其数据可靠性（HDFS）和海量数据分析的高性能（MapReduce）
  ○ mongoDB
    ⅰ. 强大的自动化 shading 功能
    ⅱ. 全索引支持，查询非常高效；
    ⅲ. 面向文档（BSON）存储，数据模式简单而强大。
    ⅳ. 支持动态查询，查询指令也使用JSON形式的标记，可轻易查询文档中内嵌的对象及数组。
    ⅴ. 支持 javascript 表达式查询，可在服务器端执行任意的 javascript函数。
    ⅵ. 缺点
      1. 单个文档大小限制为16M，32位系统上，不支持大于2.5G的数据；
      2. 对内存要求比较大，至少要保证热数据（索引，数据及系统其它开销）都能装进内存；
      3. 非事务机制，无法保证事件的原子性。



