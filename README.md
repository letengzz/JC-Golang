# 目录
## 开发准备

- [搭建Go语言开发环境](/403.md)

- [配置Go语言开发环境](/403.md)
- [Go依赖管理及Go module使用](/403.md)
- [使用Go module导入本地包](/403.md)

## Go语言基础

- [Go程序的结构及用法](/403.md)
- [变量和常量](https://github.com/letengzz/JC-Golang/tree/main/Basics/Variables_Constants) 
- [基本数据类型](https://github.com/letengzz/JC-Golang/tree/main/Basics/Basic_Data_Type) 
- [运算符](https://github.com/letengzz/JC-Golang/tree/main/Basics/Operator) 
- [流程控制](https://github.com/letengzz/JC-Golang/tree/main/Basics/Process_Control) 
- [数组](https://github.com/letengzz/JC-Golang/tree/main/Basics/Arrays) 
- [切片](https://github.com/letengzz/JC-Golang/tree/main/Basics/Slice) 
- [map](https://github.com/letengzz/JC-Golang/tree/main/Basics/map) 
- [指针](https://github.com/letengzz/JC-Golang/tree/main/Basics/pointer) 
- [函数](https://github.com/letengzz/JC-Golang/tree/main/Basics/Func) 
- [自定义类型和类型别名](https://github.com/letengzz/JC-Golang/tree/main/Basics/struct/CustomTypes_TypeAliases) 
- [结构体](https://github.com/letengzz/JC-Golang/tree/main/Basics/struct) 
- [包和依赖管理]() 
- [接口](https://github.com/letengzz/JC-Golang/tree/main/Basics/interface) 
- [反射](https://github.com/letengzz/JC-Golang/tree/main/Basics/Reflect) 
- [并发](https://github.com/letengzz/JC-Golang/tree/main/Basics/Concurrency) 
- [网络编程](/403.md)  
- [单元测试](/403.md)

## 标准库

- [官方标准库文档](https://pkg.go.dev/std)
- [标准库文档中文版](https://studygolang.com/pkgdoc)

### 常用标准库

- [fmt及格式化占位符](/403.md)
- [time](/403.md)
- [flag](/403.md)
- [log](/403.md)
- [文档操作](/403.md)
- [strconv](/403.md)
- [net/http](/403.md)
- [context](/403.md)
- [template]()

## Web相关

**Web方面**：

- [HTML/CSS]()
- [JavaScript]()
- 1

**Go Web相关**：

- [RESTful API]()

**Web框架**：

### [Gin](https://gin-gonic.com/zh-cn/docs/)

- [Gin框架安装与使用]()
- [模板渲染]()
- [返回json]()
- [获取querystring参数]()
- [获取form参数]()
- [获取URL路径参数]()
-  [参数绑定]()
- [文件上传]()
- [请求重定向]()
- [路由和路由组]()



- [gin框架中使用zap日志库](https://www.liwenzhou.com/posts/Go/use_zap_in_gin/)
- [gin框架源码解析](https://www.liwenzhou.com/posts/Go/read_gin_sourcecode/)
- [gin框架中使用validator若干实用技巧](https://www.liwenzhou.com/posts/Go/validator_usages/)
- [优雅的关机或重启gin项目](https://www.liwenzhou.com/posts/Go/graceful_shutdown/)
- [gin框架路由拆分与注册](https://www.liwenzhou.com/posts/Go/gin_routes_registry/)
- [在gin框架中使用JWT认证](https://www.liwenzhou.com/posts/Go/jwt_in_gin/)

### [beego]()

### [iris]()

### [Echo]()

### [martini]()

### [Fiber]()

## 数据库相关

[Go操作MySQL——database/sql使用指南](https://www.liwenzhou.com/posts/Go/go_mysql/)

[更强大、更好用的sqlx库使用指南](https://www.liwenzhou.com/posts/Go/sqlx/)

[Go操作Redis——go-redis库使用指南](https://www.liwenzhou.com/posts/Go/go_redis/)

[Go操作MongoDB](https://www.liwenzhou.com/posts/Go/go_mongodb/)

[usql](https://link.zhihu.com/?target=https%3A//github.com/xo/usql) - 几乎支持全部 SQL 与 NoSQL 数据库的命令行工具

[GORM](https://link.zhihu.com/?target=https%3A//github.com/go-gorm/gorm) - GORM V2

- [GORM V1](https://link.zhihu.com/?target=https%3A//github.com/jinzhu/gorm)
- [gorm2sql](https://link.zhihu.com/?target=https%3A//github.com/liudanking/gorm2sql) - 根据 Model Struct 生成建表语句

[Xorm](https://link.zhihu.com/?target=https%3A//gitea.com/xorm/xorm)

[XormPlus](https://link.zhihu.com/?target=https%3A//github.com/xormplus/xorm) - Xorm 的定制增强版本

[GoRose](https://link.zhihu.com/?target=https%3A//github.com/gohouse/gorose)

[sqlx](https://link.zhihu.com/?target=https%3A//github.com/jmoiron/sqlx) - `database/sql` 扩展包

[dbq](https://link.zhihu.com/?target=https%3A//github.com/rocketlaunchr/dbq) - 数据库操作

[gendry](https://link.zhihu.com/?target=https%3A//github.com/didi/gendry) - 滴滴开源的SQL Builder

[ozzo-dbx](https://link.zhihu.com/?target=https%3A//github.com/go-ozzo/ozzo-dbx)

[Squirrel](https://link.zhihu.com/?target=https%3A//github.com/Masterminds/squirrel) - Fluent SQL Builder

[qb](https://link.zhihu.com/?target=https%3A//github.com/aacanakin/qb) - the database toolkit for go

[redigo](https://link.zhihu.com/?target=https%3A//github.com/gomodule/redigo) - Redis 客户端

[go-redis](https://link.zhihu.com/?target=https%3A//github.com/go-redis/redis)

[mgo](https://link.zhihu.com/?target=http%3A//labix.org/mgo)

- [globalsign/mgo](https://link.zhihu.com/?target=https%3A//github.com/globalsign/mgo) - The MongoDB driver for Go
- [mgo使用指南](https://link.zhihu.com/?target=https%3A//studygolang.com/articles/3485)

[kingshard](https://link.zhihu.com/?target=https%3A//github.com/flike/kingshard) - MySQL Proxy

[SOAR](https://link.zhihu.com/?target=https%3A//github.com/XiaoMi/soar) - 对SQL进行优化和改写的自动化工具

[SQLE](https://link.zhihu.com/?target=https%3A//github.com/actiontech/sqle) - SQL 审核工具

[Vitess](https://link.zhihu.com/?target=https%3A//github.com/vitessio/vitess) - 用于部署、扩展和管理大型MySQL实例集群的数据库解决方案

[gh-ost](https://link.zhihu.com/?target=https%3A//github.com/github/gh-ost) - GitHub 开源的在线更改 MySQL 表结构的工具

[SQLer](https://link.zhihu.com/?target=https%3A//github.com/alash3al/sqler) - write APIs using direct SQL queries with no hassle, let's rethink about SQL

[gocraft/dbr](https://link.zhihu.com/?target=https%3A//github.com/gocraft/dbr)

[Gaea](https://link.zhihu.com/?target=https%3A//github.com/XiaoMi/Gaea) - 小米开源的基于 MySQL 协议的数据库中间件

[OctoSQL](https://link.zhihu.com/?target=https%3A//github.com/cube2222/octosql) - 支持多数据库的 SQL 查询工具

[goose](https://link.zhihu.com/?target=https%3A//github.com/pressly/goose) - 数据库迁移工具

[migrate](https://link.zhihu.com/?target=https%3A//github.com/golang-migrate/migrate) - 数据库迁移工具

[dbmate](https://link.zhihu.com/?target=https%3A//github.com/amacneil/dbmate) - 数据库迁移工具

[ent](https://link.zhihu.com/?target=https%3A//github.com/facebook/ent) - An Entity Framework For Go

[godb](https://link.zhihu.com/?target=https%3A//github.com/samonzeweb/godb) - a Go query builder and struct mapper

[go-nulltype](https://link.zhihu.com/?target=https%3A//github.com/mattn/go-nulltype)

[go-mysql](https://link.zhihu.com/?target=https%3A//github.com/siddontang/go-mysql) - MySQL [工具集](https://www.zhihu.com/search?q=工具集&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A2336661362})

[SQLittle](https://link.zhihu.com/?target=https%3A//github.com/alicebob/sqlittle) - 纯读取 SQLite 文件

[Bifrost](https://link.zhihu.com/?target=https%3A//github.com/brokercap/Bifrost) - MySQL 同步到 Redis、ClickHouse 等服务的异构中间件

[elasticsql](https://link.zhihu.com/?target=https%3A//github.com/cch123/elasticsql) - 转换 SQL 成 Elasticsearch DSL

[POP](https://link.zhihu.com/?target=https%3A//github.com/gobuffalo/pop) - 基于 [sqlx](https://link.zhihu.com/?target=https%3A//github.com/jmoiron/sqlx) 封装的数据库 ORM 工具

[REL](https://link.zhihu.com/?target=https%3A//github.com/go-rel/rel) - Modern Database Access Layer for Go

## 单元测试

- [单元测试基础]()
- [网络测试]()
- [MySQL和Redis测试]()
- [mock接口测试]()
- [monkey打桩测试]()
- [goconvey的使用]()
- [编写可测试的代码]()

## 分布式系统

- [go-kit](https://link.zhihu.com/?target=https%3A//github.com/go-kit/kit) **star:22226** 支持服务发现、负载平衡、插件式传输、请求跟踪等功能的Microservice toolkit。 
- [go-micro](https://link.zhihu.com/?target=https%3A//github.com/micro/go-micro) **star:17587** 分布式系统开发框架。 
- [Kratos](https://link.zhihu.com/?target=https%3A//github.com/go-kratos/kratos) **star:16321** 一个模块化设计的易于使用的微服务框架。 
- [grpc-go](https://link.zhihu.com/?target=https%3A//github.com/grpc/grpc-go) **star:15266** gRPC的Go语言实现。 
- [go-zero](https://link.zhihu.com/?target=https%3A//github.com/tal-tech/go-zero) **star:14247** 一个web和rpc框架。它的诞生是为了确保繁忙场地的稳定性，弹性设计。内置的goctl大大提高了开发效率。 
- [micro](https://link.zhihu.com/?target=https%3A//github.com/micro/micro) **star:10867** 一个分布式系统运行时，用于云及其他。
- [NATS](https://link.zhihu.com/?target=https%3A//github.com/nats-io/gnatsd) **star:10404** 轻量级、高性能消息传递系统，可用于微服务、物联网(IoT)和云。 
- [rpcx](https://link.zhihu.com/?target=https%3A//github.com/smallnest/rpcx) **star:6651** 分布式可插拔的RPC服务框架，如阿里巴巴Dubbo。

## 代码分析

- [reviewdog](https://link.zhihu.com/?target=https%3A//github.com/reviewdog/reviewdog) - Code Review 机器人
- [revive](https://link.zhihu.com/?target=https%3A//github.com/mgechev/revive) - 代码检查分析
- **[GolangCI-Lint](https://link.zhihu.com/?target=https%3A//github.com/golangci/golangci-lint)** - 代码质量检查分析工具
- [errcheck](https://link.zhihu.com/?target=https%3A//github.com/kisielk/errcheck) - 检测未处理的错误(errors)
- [Staticcheck](https://link.zhihu.com/?target=https%3A//github.com/dominikh/go-tools) - 一系列的 Go 代码静态分析工具
- [Golint](https://link.zhihu.com/?target=https%3A//github.com/golang/lint) - Google 官方出品的代码质量检测工具
- [GoReporter](https://link.zhihu.com/?target=https%3A//github.com/360EntSecGroup-Skylar/goreporter)
- [go-critic](https://link.zhihu.com/?target=https%3A//github.com/go-critic/go-critic)
- [gocloc](https://link.zhihu.com/?target=https%3A//github.com/hhatto/gocloc) - 分语言代码行数统计
- [coca](https://link.zhihu.com/?target=https%3A//github.com/phodal/coca) - 代码统计分析
- **[Go Report Card](https://link.zhihu.com/?target=https%3A//github.com/gojp/goreportcard)** - Go 项目质量分析报告工具
- [ddsv-go](https://link.zhihu.com/?target=https%3A//github.com/y-taka-23/ddsv-go) - 死锁检测工具
- [golang/perf](https://link.zhihu.com/?target=https%3A//github.com/golang/perf) - 官方性能[量化分析](https://www.zhihu.com/search?q=量化分析&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A2336661362})工具
- [GoPlantUML](https://link.zhihu.com/?target=https%3A//github.com/jfeliu007/goplantuml) - 为 Go 项目生成 PlantUML 类图
- [gosize](https://link.zhihu.com/?target=https%3A//github.com/bradfitz/gosize) - 分析Go二进制文件大小
- [shotizam](https://link.zhihu.com/?target=https%3A//github.com/bradfitz/shotizam) - 分析 Go 二进制文件的大小并输出到 SQLite3
- [goconst](https://link.zhihu.com/?target=https%3A//github.com/jgautheron/goconst) - 查找可以被常量替换的重复字符串
- [sploit](https://link.zhihu.com/?target=https%3A//github.com/zznop/sploit) - 帮助二进制分析和开发的库
- [perf](https://link.zhihu.com/?target=https%3A//github.com/hodgesds/perf-utils) - Perf Utilities for Go
- [fgprof](https://link.zhihu.com/?target=https%3A//github.com/felixge/fgprof) - Go 性能分析工具
- [conprof](https://link.zhihu.com/?target=https%3A//github.com/conprof/conprof) - 协程分析
- [statsview](https://link.zhihu.com/?target=https%3A//github.com/go-echarts/statsview) - 实时 Go 运行时统计数据可视化分析器
- [codesearch](https://link.zhihu.com/?target=https%3A//github.com/google/codesearch) - 代码搜索工具
- [Pyroscope](https://link.zhihu.com/?target=https%3A//github.com/pyroscope-io/pyroscope) - 可视化程序性能监控工具，支持多种语言
- [gosec](https://link.zhihu.com/?target=https%3A//github.com/securego/gosec) - 代码安全性检查工具
- [gokart](https://link.zhihu.com/?target=https%3A//github.com/praetorian-inc/gokart) - 代码静态分析工具
- [gofumpt](https://link.zhihu.com/?target=https%3A//github.com/mvdan/gofumpt) - `gofmt` 增强版代码格式化工具

## 组件及技巧

- [Go语言json技巧](https://www.liwenzhou.com/posts/Go/json_tricks_in_go)
- [option选项模式](https://www.liwenzhou.com/posts/Go/functional_options_pattern/)
- [Go语言中的单例模式（翻译）](https://www.liwenzhou.com/posts/Go/singleton_in_go/)
- [结构体转map的若干方法](https://www.liwenzhou.com/posts/Go/struct2map/)
- [Go语言配置管理神器——Viper中文教程](https://www.liwenzhou.com/posts/Go/viper_tutorial/)
- [protobuf初识](https://www.liwenzhou.com/posts/Go/protobuf/)
- [gRPC初识](https://www.liwenzhou.com/posts/Go/gRPC/)
- [Go操作NSQ](https://www.liwenzhou.com/posts/Go/go_nsq/)
- [Go操作kafka](https://www.liwenzhou.com/posts/Go/go_kafka/)
- [Go操作etcd](https://www.liwenzhou.com/posts/Go/go_etcd/)
- [Go语言获取系统性能数据gopsutil库](https://www.liwenzhou.com/posts/Go/go_gopsutil/)
- [二进制协议gob及msgpack介绍](https://www.liwenzhou.com/posts/Go/gob_msgpack/)
- [influxDB](https://www.liwenzhou.com/posts/Go/go_influxdb/)
- [Elasticsearch](https://www.liwenzhou.com/posts/Go/go_elasticsearch/)
- [Go第三方日志库logrus](https://www.liwenzhou.com/posts/Go/go_logrus/)
- [Go语言项目中使用zap日志库（翻译）](https://www.liwenzhou.com/posts/Go/zap/)
- [Go pprof性能调优](https://www.liwenzhou.com/posts/Go/performance_optimisation/)
- [为Go项目编写Makefile](https://www.liwenzhou.com/posts/Go/makefile/)
- [在select语句中实现优先级](https://www.liwenzhou.com/posts/Go/priority_in_go_select/)
- [在关于切片操作的技巧（翻译）](https://www.liwenzhou.com/posts/Go/slice_tricks/)
- [Go语言结构体的内存布局](https://www.liwenzhou.com/posts/Go/struct_memory_layout/)

## 其他

[使用Air实现Go程序实时热重载](https://www.liwenzhou.com/posts/Go/live_reload_with_air/)

[如何使用docker部署Go Web程序](https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/)

[Cookie和Session](https://www.liwenzhou.com/posts/Go/Cookie_Session/)

[使用swagger生成接口文档](https://www.liwenzhou.com/posts/Go/gin_swagger/)

[HTTP Server常用压测工具介绍](https://www.liwenzhou.com/posts/Go/benchmark_tool/)

[漏桶和令牌桶限流策略介绍及使用](https://www.liwenzhou.com/posts/Go/ratelimit/)

[部署Go语言程序的N种方法](https://www.liwenzhou.com/posts/Go/deploy_go_app/)

### 爬虫

### GitHub



### RabbitMQ

- [RabbitMQ Go客户端教程1——HelloWorld（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_01/)
- [RabbitMQ Go客户端教程2——任务队列（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_02/)
- [RabbitMQ Go客户端教程3——发布/订阅（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_03/)
- [RabbitMQ Go客户端教程4——路由（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_04/)
- [RabbitMQ Go客户端教程5——topic（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_05/)
- [RabbitMQ Go客户端教程6——RPC（翻译）](https://www.liwenzhou.com/posts/Go/go_rabbitmq_tutorials_06/)

### 工具使用

- [Goland]()

- [VS Code]()

- [postman]()

  
