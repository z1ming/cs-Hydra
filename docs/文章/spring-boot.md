# Spring Boot 面试题

## 介绍一下Springboot Starter的工作机制？


1. 在依赖的 Starter 包中寻找 resources/META-INF/spring.factories 文件，然后根据文件中配置的 Jar 包去扫描所依赖的 Jar 包
2. 根据 spring.factories 配置加载 AutoConfigure 类
3. 根据 @Conditional 注解的条件，进行自动配置并将 Bean 注入 Spring Context

简单来讲，先读取 Spring Boot Starter 的配置信息，在根据配置信息进行资源初始化，并注入到 Spring 容器中。这样 Spring Boot 启动完毕后，就已经准备好了一切资源，使用过程中直接注入对应的 Bean 资源即可。

