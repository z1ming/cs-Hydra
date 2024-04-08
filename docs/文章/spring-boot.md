# Spring Boot 面试题

## 介绍一下 Springboot Starter 的工作机制？


1. 在依赖的 Starter 包中寻找 resources/META-INF/spring.factories 文件，然后根据文件中配置的 Jar 包去扫描所依赖的 Jar 包
2. 根据 spring.factories 配置加载 AutoConfigure 类
3. 根据 @Conditional 注解的条件，进行自动配置并将 Bean 注入 Spring Context

简单来讲，先读取 Spring Boot Starter 的配置信息，在根据配置信息进行资源初始化，并注入到 Spring 容器中。这样 Spring Boot 启动完毕后，就已经准备好了一切资源，使用过程中直接注入对应的 Bean 资源即可。

## Spring 循环依赖怎么解决？

循环依赖是指 A 类中有属性 B，B 类中有属性 A，主要有三种情况：

1. 通过构造方法依赖注入时产生循环依赖问题
2. 通过 setter 方法依赖注入且是多例模式下产生循环依赖问题
3. 通过 setter 方法依赖注入且是单例模式下产生循环依赖问题

只有第三种循环依赖方式被 Spring 解决了，其他两种方式都会产生异常。Spring 解决单例模式下的 setter 循环依赖问题是通过三级缓存解决的。三级缓存是指在初始化过程中，通过三级缓存来缓存正在创建的 Bean，以及创建完成的 Bean。具体步骤如下：

- 实例化 Bean：先创建一个空对象，并将其放入一级缓存中
- 属性赋值：Spring 开始对 Bean 进行属性赋值，如果发现循环依赖，会将当前 Bean 对象提前暴露给后续需要依赖的 Bean
- 初始化 Bean：完成属性赋值后，Spring 将 Bean 进行初始化，并将其放入二级缓存中
- 注入依赖：Spring 继续对 Bean 进行依赖注入，如果发现循环依赖，会从二级缓存中获取已经完成初始化的 Bean 实例

通过三级缓存机制能确保 Spring 在处理循环依赖时，将正在创建的 Bean 对象及时暴露出来，并能够正确地注入已经初始化的 Bean 实例，从而解决循环依赖问题。

## Spring 三级缓存的数据结构是什么？

都是 Map 类型的缓存，比如 `Map<k:name; v:Bean>`

1. 一级缓存（Singleton Objects）：存储的是完全初始化好的 Bean，即完全准备好可以使用的 Bean 实例。结构是“bean 名称： bean 实例”，这个缓存在 `DefaultSingletonBeanRegistry.singletonObjects` 中
2. 二级缓存（Early Singleton Objects）：存储的是早期的 Bean 引用，即已经实例化但还未完全初始化的 Bean。这些 Bean 已经被实例化，但还没有进行属性注入等操作, 在`DefaultSingletonBeanRegistry.earlySingletonObjects` 属性中
3. 三级缓存（Singleton Factories）：存储的是 ObjectFactory 对象，这些对象可以生成早期的 bean 引用。当一个 Bean 正在创建过程中，如果它被其他 Bean 依赖，那么这个正在创建的 Bean 就会通过这个 ObjectFactory 来创建一个早期引用，从而解决循环依赖问题。这个缓存在 `DefaultSingletonBeanRegistry.singletonFactories` 属性中
