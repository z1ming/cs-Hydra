# 计算机网络面试题

## 输入网址后发生了什么？

1. 应用层 DNS 解析：将域名转化为 ip 地址
2. 传输层 TCP 连接：浏览器通过 DNS 获取到真实的 ip 地址之后，便向 web 服务器发起 TCP 连接请求，通过 TCP 三次握手建立连接
3. 网络层 IP：建立 TCP 连接时，需要发送数据，发送数据在网络层使用 IP 协议，通过 IP 协议将 IP 地址封装为 IP 数据报，此时会用到 ARP 协议，主机发送信息时将包含目标 IP 地址的 ARP 请求广播到网络上的所有主机，并接收返回消息，以此确定目标的物理地址，找到目的 MAC 地址
4. IP 数据包在路由器之间，路由选择使用 OPSF 协议，采用 Dijkstra 算法来计算最短路径树，抵达服务器
5. 发送 HTTP 请求：建立 TCP 连接之后，浏览器向 Web 服务器发起一个 HTTP 请求；如果是 HTTPS 协议，发送 HTTP 请求之前还要完成 TLS 四次握手
6. 处理请求并返回：服务器获取到客户端的 HTTP 请求之后，会根据 HTTPS 请求中的内容来决定如何获取相应的文件，并将文件发送给浏览器
7. 浏览器渲染：浏览器根据响应开始显示页面，首先解析 HTML 文件构建 DOM 树，然后解析 CSS 文件构建渲染树，等到渲染树构建完成后，浏览器开始布局渲染树并将其绘制到屏幕

## HTTP 为什么是无状态的？

HTTP 请求无状态是因为每个请求之间是相互独立的，服务器不会在不同请求之间保留客户端的状态信息。

无状态的原因：

1. 简单性：无状态使得 HTTP 协议涉及更加简单和灵活，每个请求都可以独立处理，不需要维护复杂的状态信息
2. 可伸缩性：因为不需要保存状态信息，更容易进行水平扩展，处理更多的并发请求
3. 基于缓存：无状态使得缓存更加有效，可以缓存响应并在多个请求之间共享，提高性能和减少网络流量

尽管 HTTP 是无状态的，浏览器通常会通过 Cookies，Session 等技术来耿总用户状态。

## 网络常见的通信协议？

- HTTP：应用层协议。在 web 浏览器和 web 服务器之间的超文本传输协议，是目前最常见的应用层协议
- HTTPS：在 HTTP 基础上添加了 SSL/TLS 加密层，用于在不安全的网络上安全地传输数据
- TCP：面向连接的传输层协议，提供可靠的数据传输服务，保证数据的顺序和完整性
- UDP：无连接的传输层协议，提供了数据包传输的简单服务，用于实时性要求高的应用
- IP：网络层协议，用于在网络中传输数据包，定义了数据包的格式和传输规则

## 前后端交互用的什么协议？

用 HTTP 和 HTTPS。

## TCP 三次握手介绍一下

- 首先，服务器处于 LISTEN 状态，监听某个端口
- 客户端初始化随机序列号 `client_isn`，将 SYN 标志为置为 1，将报文发送给服务端，该报文不包含应用层数据，之后客户端处于 SYN-SEND 状态
- 服务段收到 SYN 报文后，也初始化自己的序列号 (server_isn)，将此序号填入 TCP 首部的序号字段中，把 TCP 首部的确认应答号字段填入 `client_isn + 1`，将 SYN 和 ACK 标志位置为 1，最后发给客户端，处于 SYN—RCVD 状态
- 客户端收到服务端报文后，还要向服务端回应最后一个应答报文，将 ACK 标志位置为 1，确认应答号字段填入 server_isn + 1，
将报文发给服务端，之后客户端处于 ESTABLISHED 状态
- 服务端收到客户端的应答报文后，也进入 ESTABLISHED 状态

## TCP 四次挥手说一下

- 客户端主动关闭连接，发送 FIN 报文，FIN 代表客户端不会在发送数据了，进入 FIN_WAIT_1 状态
- 服务端收到 FIN 报文后，马上回复一个 ACK 报文，服务端进入 CLOSE_WAIT 状态。在收到 FIN 报文之后，TCP 协议栈会为 FIN 包插入一个文件结束符 EOF 到接收缓冲区中，服务端应用程序可以通过 read 调用来感知这个 FIN 包，这个 EOF 会被放到已排队的其他已接收的数据之后，此时 read 还在进行
- 服务端继续读取数据，知道读到 EOF 的数据，接着 read 返回 0，这时服务端有数据，就发完数据再调用关闭连接的函数，否则可以直接调用关闭连接的函数，这时服务端会发一个 FIN 包，这个 FIN 报文代表服务端不会再发送数据了，之后处于 LAST_ACK 状态
- 客户端接受到服务端的 FIN 包，并发送 ACK 给服务端，客户端进入 TIME_WAIT 状态
- 服务端收到 ACK 后，进入 CLOSE 状态
- 客户端经过 2MSL 时间之后，也进入 CLOSE 状态

## TCP 和 UDP 的区别是什么？

- 连接：TCP 有连接，UDP 无连接
- 可靠性：TCP 是可靠交付的，数据可以无差错，不丢失，不重复，按序到达；UDP 尽最大努力交付，不保证可靠数据交付。但是我们可以基于 UDP 传输协议实现可靠的传输协议，如 QUIC 协议
- 拥塞控制、流量控制：TCP 有拥塞控制和流量控制，保证数据传输的安全性；UDP 即使网络非常堵了，发送速率不会变
- 首部开销：TCP 首部没有使用“选项”字段的开销是 20 字节，如果使用“选项”字段则变长。UDP 首部只有 8 字节，并且是固定不变的，开销较小
- 服务对象：TCP 是一对一，UDP 支持一对一、一对多、多对多通信
- 传输方式：TCP 是面向字节流的，没有边界，但保证数据的顺序和可靠。UDP 是一个包一个包发送，有边界的，但是可能会乱序和丢包
- 应用场景：TCP 是面向连接的可靠协议，经常用于 HTTPS/HTTP，FTP。UDP 是面向无连接，随时可以发送数据，UDP 本身处理既简单又高效，经常用于视频、音频等多媒体信息

## TCP 如何保证数据可靠？

- 连接：三次握手建立连接，四次挥手关闭连接
- 丢包
    - 丢数据包：超时重传
    - 丢 ACK 包：由于 ACK 包丢失，发送端认为数据没有发过去，所以会再次发送，因此对方有可能收到重复数据，这时通过序列号判断数据是否重复，并重新发送 ACK 包
- 顺序：通过序列号保证数据有序，比如每次发送时，带一个序列号 x，对方会回复 x+1 的 ACK 包，这是发送方就知道从哪个序列号开始发送数据

## 拥塞控制的流程？

- 慢启动
- 拥塞避免
- 拥塞发生
- 快速恢复

## DNS 解析的具体过程

DNS 域名解析就是把域名翻译成 ip 地址的过程，例如将 www.baidu.com 翻译成 1.2.3.4 这个 IP，总共分为以下步骤：

- 浏览器缓存检查：浏览器首先会检查自身 DNS 缓存，这个缓存只能容纳 1000 条，且缓存时间比较短，只有一分钟。如果没有查到缓存，则进行下一步
- 操作系统缓存检查 + hosts 解析：如果浏览器缓存没有查到，则会查找操作系统缓存的解析结果。Linux 可以通过 /etc/hosts 文件来设置域名对应的 ip，如果在这个文件查到了域名，会优先访问这里配置的 ip，并将这个结果缓存起来。缓存时间受域名失效时间和缓存大小控制的
- dns 解析 1：客户端通过浏览器访问域名为 www.baidu.com 的网站，发起查询该域名 ip 的 dns 请求。该请求发送到本地 dns 服务器，本地 dns 服务区首先会检查自己的缓存记录，如果缓存中有此记录，就可以直接返回结果，否则本地 dns 服务器向根服务器发起请求
- dns 解析 2：本地 dns 服务器向根服务器发送 dns 请求，请求域名为 www.baidu.com 的 ip 地址
- dns 解析 3：根服务器经过查询，没有找到该域名及 ip 的对应关系，告诉本地 dns 服务器，可以到 .com 域名服务器上继续查询
- dns 解析 4：本地 dns 服务器向 .com 服务器发送 dns 请求，请求域名为 www.baidu.com 的 ip 地址
- dns 解析 5：.com 服务器接收到请求后，不会立即返回对应 ip，而是告诉本地 dns 服务器，可以在 baidu.com 的域名服务器上请求获取 ip 地址
- dns 解析 6：本地 dns 服务器向 baidu.com 服务器发送 dns 请求，请求域名为 www.baidu.com 的 ip 地址
- dns 解析 7：baidu.com 服务器在自己的缓存表中发现了该域名和 ip 地址的对应关系，并将 ip 地址返回给本地 dns 服务器
- dns 解析 8：本地 dns 服务器将获取到域名与对应的 ip 地址返回给客户端，并将结果缓存供下次查询

## DNS 基于 TCP 还是 UDP？为什么？

基于 UDP 实现域名解析和数据传输。因为 UDP 实现 DNS 能够提供低延迟、简单快速、轻量级的特性，更适合 DNS 这种需要快速响应的域名解析服务。

- 低延迟：UDP 是一种无连接的协议，适合 DNS 这种需要快速响应的场景
- 简单快速：UDP 相比于 TCP 更简单，没有 TCP 的连接管理和流量控制机制，传输效率更高，适合 DNS 这种需要快速传输数据的场景
- 轻量级：UDP 头部较小，占用较少的网络资源，对于小型请求和响应来说更加轻量级，适合 DNS 这种频繁且短小的数据交换

## 了解过 DNS 劫持吗？

DNS 劫持的原理是攻击者在查询 DNS 服务器时篡改响应，将用户请求的域名映射到攻击者控制的虚假 IP 上，使用户误以为是正常的网站，实际上被攻击者操控的恶意网站。这种劫持可以通过植入恶意的 DNS 记录或劫持用户的 DNS 流量实现。

## 子网掩码的作用是什么？

它的作用是用来指明一个 IP 地址的哪些位标识的是主机所在的网络地址一级哪些位标识的是主机地址的位掩码[1]。有两个方面：

1. 划分网络和主机：子网掩码通过指示 IP 地址中的网络部分和主机部分，帮助路由器确定网络内部和网络间的通信
2. 确定网络范围：通过与 IP 地址进行逻辑与操作，子网掩码帮助确定一个 IP 地址所在的网络范围，以便正确路由数据包

## TCP 拆包沾包的原因

- TCP 协议是基于字节流的传输协议，没有固定的分包边界。发送方将数据拆分成多个小的数据包进行传输，接收方再将这些数据包组合成完整的数据。在这个过程中可能会出现拆包和沾包现象
- 网络传输中的延迟和拥塞会影响数据包发送的速度和到达接收方的顺序，可能导致数据包的拆分和组合不规律
- 接收方缓冲区大小限制。当接收方缓冲区大小小于完整数据包大小，可能会发生拆包

解决方式：

- 在应用层实现数据包的边界识别，添加包头，包头中添加数据长度等信息
- 使用固定长度的数据包或特殊的分隔符，便于数据包识别边界
- 使用更高级的传输层协议，如 WebSocket，它在 TCP 的基础上添加了数据帧的概念，可以更好地解决拆包和沾包问题

## 对称加密和非对称加密的区别？

- 对称加密是双方使用同一个密钥加解密，优点是运算速度快，缺点是安全性不高，一旦密钥泄露则有风险
- 非对称加密的安全性高，但计算速度低，一般非对称加密用在验证证书阶段，对称加密用在数据传输阶段，非对称加密的基本路程如下[1]：
	- 假设爱丽丝和鲍勃事先互不认识，也没有可靠安全的沟通通道，爱丽丝现在要通过不安全的互联网向鲍勃发送消息
	- 爱丽丝写好原文，设为 x
	- 鲍勃使用密码学安全伪随机数生成一对密钥，一个作为公钥 c，一个作为私钥 d
	- 鲍勃将公钥 c 通过不安全的互联网发送给爱丽丝，即使窃听者伊夫在中间窃听到 c 也没问题
	- 爱丽丝用公钥 c 对原文加密，得到密文 c(x)
	- 爱丽丝可以用任何方法传输密文 c(x) 给鲍勃，即使窃听者伊夫在中间窃听到 c(x) 也没问题
	- 鲍勃收到密文 c(x)，用私钥 d 对 c(x) 进行解密，即 d(c(x))，得到原始明文 x
	- 由于伊夫没有得到私钥 d，所以无法获取明文 x
	- 如果爱丽丝丢失了原文 x，在没有得到鲍勃的私钥 d 的情况下，她和伊夫一样无法通过鲍勃的公钥 c 和密文 c(x) 重新得到原文 x

## 有哪些对称加密和非对称加密的算法？

- 对称加密：
	- AES：最流行和广泛使用的算法之一，支持不同的密钥长度。如 AES-128，AES-256
- 非对称加密算法：
	- RSA：最常见的非对称加密算法，用于数据加密和数字签名
	- ECC：基于椭圆曲线的非对称加密算法，具有较高的安全性和效率，使用于移动设备等资源受限的环境

## 在信息传输过程中，HTTPS用的是对称加密还是非对称加密

HTTPS 使用的混合加密方式，对称加密和非对称加密都有。

- 在通信建立前采用非对称加密的方式交换会话密钥，后续就不再使用非对称加密
- 在通信过程中使用对称加密的会话密钥方式加密会话数据

## HTTPS 四次握手过程说一下

SSL/TLS 协议的基本流程：

- 客户端向服务端索要并验证服务器的公钥
- 双方协商生成会话密钥
- 双方使用会话密钥进行加密通信

前两步即 SSL 建立握手的阶段。TLS 协议建立的详细流程：

1. ClientHello

客户端向服务端发送加密通信请求，即 ClientHello 请求，客户端发送以下信息：

- 客户端支持的 TLS 版本，如 TLS1.2 版本
- 客户端生产的随机数（Client Random），后面用于生成“会话密钥”的条件之一
- 客户端支持的密码套件列表，如 RSA 加密算法

2. ServerHello

服务端收到客户端请求之后，向客户端发出响应，也就是 ServerHello，服务器回应以下内容：

- 确认 TLS 协议版本，如果浏览器不支持，则关闭加密通信
- 服务器生产的随机数（Server Random），也是后面用于生产会话密钥的条件之一
- 确认的密码套件列表，如 RSA 加密算法
- 服务器的数字证书

3. 客户端回应

客户端收到服务器的回应之后，首先通过浏览器或者操作系统中的 CA 公钥，确认服务器的数字证书的真实性。

如果证书没有问题，客户端会从数字证书中取出服务器的公钥，然后用它加密报文，向服务器发送如下信息：

- 一个随机数（pre-master-key），该随机数会被服务器加密
- 加密算法改变通知，表示随后的信息都将用“会话密钥”加密通信
- 客户端握手结束通知，表示客户端的握手阶段已经结束，同时把之前所有的内容做个摘要，供服务器校验

客户端和服务端有了 Client Random，Server Random，pre-master-key 之后，接着就用双方协商的加密算法，各自生成本次通信的会话密钥。

4. 服务端回应

服务器收到第三个随机数 pre-master-key 之后，通过协商的加密算法，计算出本次通信的会话密钥

然后，向客户端发送最后的信息：

- 加密通信算法改变通知，表示随后的信息都将用于“会话密钥”加密通信
- 服务器握手结束通知，表示服务器的握手阶段已经结束，这项会把之前的内容做个摘要，用来供客户端校验

至此，TLS 握手结束，接下来使用“会话密钥”加密通信内容。

## 中间人劫持了会怎么样？

中间人服务器介于客户端和服务器之间，客户端以为自己在于真正的服务器 TLS 握手，实际上与中间人服务器进行 TLS 握手，中间人服务器与真正服务器进行 TLS 握手。与客户端通信时，中间人服务器使用自己的证书，客户端在浏览时一般浏览器会做提醒，提示此网站的安全证书存在问题。

## CA 签发流程是什么？

- 服务方向第三方机构提交公钥、组织信息、个人信息（域名）等信息并申请认证（不交私钥）
- CA 通过线上、线下等多种手段验证申请者提供信息的真实性，如组织是否存在，企业是否合法，是否拥有域名的所有权等
- 如果审核通过，CA 会向申请者签发认证文件-证书
  - 证书包含以下信息：申请者公钥、申请者的组织和个人信息、签发机构 CA 的信息、有效时间、证书序列号等信息的明文，同时包含一个签名
 
## 浏览器缓存内 CA 证书哪里来的？

浏览器缓存中的 CA 证书通常来自于浏览器预装的根证书颁发机构，这些根证书是浏览器或操作系统内置的信任的 CA 机构颁发的证书。

当用户访问一个使用 SSL/TLS 加密的网站时，浏览器会检查网站提供的数字证书是否由根证书中的 CA 机构颁发，从而验证网站的身份和安全性。如果数字证书的颁发机构是根证书的 CA 之一，浏览器会信任该证书，建立安全连接。

## 最后加密的时候，是怎么加密的？是用三个随机数加密吗？

客户端和服务器手里有了三个随机数 Client Random、Server Random 和 Pre-Master，用这三个作为原始材料，就可以生成用于加密会话的主密钥，叫 Master Secret。而黑客因为拿不到 Pre-Master，也就得不到主密钥。

为了保证真正的完全随机和不可预测，把三个不可靠的随机数混合起来，那么随机的程度就非常高了，足够让黑客难以预测。

RFC 文档里，Master Secret 的计算方式：

```
master_secret = PRF(pre_master_secret, "master secret",
  		    ClientHello.random + ServerHello.random)
```

PRF 是伪随机函数，它基于密码套件里的最后一个参数，比如这次的 SHA384，通过摘要算法再次强化 Master Secret 的随机性。

主密钥有 48 字节，但它也不是最终用于通信的会话密钥，还会再用 PRF 拓展出更多的密钥，比如客户端发送用的会话密钥（client_write_key），服务器发送用的会话密钥（server_write_key）等等，避免只用一个密钥带来安全隐患。

## HTTPS客户端验证证书的细节？[2]

CA 机构首先进行证书签名，然后客户端进行校验。

证书签名过程：

- CA 将持有者的公钥、用途、颁发者、有效时间等信息打成一个包，然后对这些信息进行 Hash 计算，得到一个 Hash 值
- CA 使用私钥将 Hash 值加密，形成 Certificate Signature，即对证书进行签名
- 将 Certificate Signature 添加在文件证书上，形成电子证书

客户端校验证书过程：

- 客户端使用同样的 Hash 算法得到该证书的 Hash 值 H1
- 浏览器和操作系统中集成了 CA 的公钥信息，浏览器收到证书后可以使用 CA 的公钥对 Certificate Signature 解密，得到 Hash 值 H2
- 比较 H1 和 H2，如果值相同，则为可信赖的证书，否则认为证书不可信

实际上证书的签发还有一个证书信任链的问题，我们向 CA 申请的证书一般不是根证书签发的，而是由中间证书签发的。假如客户端收到的是 baidu.com 证书，层级关系如下：

- 客户端收到 baidu.com 的证书后，发现这个证书的签发者不是根证书，就无法根据本地已有的根证书中的公钥去验证 baidu.com 证书是否可信。于是，客户端根据 baidu.com 证书中的签发者，找到该证书的颁发机构是 “GlobalSign Organization Validation CA - SHA256 - G2”，然后向 CA 请求该中间证书
- 请求到证书后发现 “GlobalSign Organization Validation CA - SHA256 - G2” 证书是由 “GlobalSign Root CA” 签发的，由于 “GlobalSign Root CA” 已经是根签发机构，应用软件会检查此证书是否已预加载于跟证书清单上，如果有，则可以利用根证书中的公钥去验证“GlobalSign Organization Validation CA - SHA256 - G2”，如果验证通过，认为中间证书是可信的
- “GlobalSign Organization Validation CA - SHA256 - G2” 证书被信任后，可以使用 “GlobalSign Organization Validation CA - SHA256 - G2” 证书中的公钥去验证 baidu.com 证书的可信性，如果验证通过，就可以新增 baidu.com 证书

整个信任链路如下：

- 服务器证书的颁发者 -> 中间证书的颁发者 -> 根证书的颁发者
- 根证书自验证 -> 中间证书验证 -> 服务器证书验证

## 说说 HTTP2

HTTP/2 相比 HTTP/1.1性能上的改进：

- 头部压缩：HTTP/2会压缩头。如果你同时发出多个请求，他们的头是一样或类似的，协议会帮你消除重复的部分，这就是所谓的 HPACK 算法：在客户端和服务器同时维护一张头信息表，所有字段都会存入这张表，生成一个索引号，以后就不发送同样字段了，只发送索引号，这样能提高传输效率
- 二进制格式：HTTP/2 不再像 HTTP1.1 里的纯文本形式的报文，而是全面采用了二进制格式，头信息和数据体都是二进制，并且统称为帧：头信息帧（Headers Frame）和数据帧（Data Frame），虽然对人不友好，但对计算机友好，因为计算机只懂二进制，收到报文后，无需再将明文的报文转成二进制，而是直接解析二进制报文，这增加了数据传输的效率
- 并发传输：引出了 Stream 的概念，多个 Stream 复用一条 TCP 连接，解决了 HTTP/1.1 队头阻塞的问题
- 服务器主动推送资源：可以主动向客户端推送资源

## HTTP 断点续传是什么？

断点续传是 HTTP/1.1 协议支持的特性。实现断点续传的功能，需要客户端记录下当前的下载进度，并在需要续传的时候通知服务端本次需要下载的内容片段。一个简单的断点续传流程：

1. 客户端开始下载一个 1024K 的文件，服务端发送 Accept-Ranges: bytes 来告诉客户端，其支持带 Range 请求
2. 假如客户端下载了其中 512K 的时候网络断开了，过了一会网络可以了，客户端再下载的时候，需要在 HTTP 头中声明本次需要续传的片段：Range:bytes=512000- 这个头通知服务端从头文件的 512K 位置开始传输文件，直到文件内容结束
3. 服务端收到断点续传请求，从文件的 512K 位置开始传输，并且在 HTTP 头中增加：Content-Range:bytes 512000-1024000, Content-Length: 512000。并且此时服务端返回的 HTTP 状态码应该是 206 Partial Content。如果客户端传递过来的 Range 超过资源的大小，则响应 416 Requested Range Not Satisfiable

综上，断点续传有如下 4 个头：

- Range: bytes=开始位置-结束位置：浏览器告诉服务器所需分部分内容范围的消息头
- Content-Range
- Accept-Range: bytes 这个值声明了可被接收的每一个范围的请求，大多数情况下是字节数 bytes
- Content-Length

## TLS 握手里面的哈希函数用来干什么？

用来计算握手消息的摘要，确保消息的完整性。

## TLS 握手过程中如何确定对方的身份

通过数字证书。服务器会向客户端发送自己公钥的数字证书，客户端通过验证数字证书的有效性、颁发者和域名等信息来确认服务器的身份。另外，TLS 握手还包括双方随机数的生成、密钥协商等步骤来确保通信的安全并确认对方的身份。

## 如何防止下载的文件被劫持和篡改？

- 使用 HTTPS 下载：确保下载链接使用 HTTPS 协议，通过 SSL 加密传输协议，防止中间人攻击和数据篡改
- 验证文件完整性：在下载文件后使用 Hash 算法（如 MD5，SHA-256）计算文件的哈希值，于官方提供的哈希值进行比对，验证文件是否被篡改

## 一台机器理论上能创建多少条 TCP 连接？

如果在不考虑服务器内存和文件句柄资源的条件下，理论上一个服务器进程最多支持 2 的 48 次方个连接，约等于 200 万亿！

- 理论上讲，2 的 32 次方（ip 数）* 2 的 16次方（端口数）[4]
- 实际中能支持 1000 条连接已经很好了

## 参考文章

1. 子网掩码 wiki：https://zh.wikipedia.org/wiki/%E5%AD%90%E7%BD%91
2. 公开密钥加密：https://zh.wikipedia.org/wiki/%E5%85%AC%E5%BC%80%E5%AF%86%E9%92%A5%E5%8A%A0%E5%AF%86
3. 签发证书的过程：https://mp.weixin.qq.com/s/bygsszqpqCuWdC0yxd4jMA
4. 一台服务器最大能支持多少条TCP连接: https://juejin.cn/post/7162824884597293086
