# gin-api-http
GIN HTTP API 服务


## 关于

`gin-api-http` 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发; 使用wire依赖注入,约束项目组开发，规避混乱无序及自由随意的编码。

供参考学习，线上使用请谨慎！

集成组件：
1. 支持 [sentinel](https://sentinelguard.io/zh-cn/docs/golang/quick-start.html) 接口限流/熔断降级/系统自适应过载保护 等服务治理
2. 支持 [cors](https://github.com/rs/cors) 接口跨域、zaplog、recover 等中间件
3. 支持 [wire](https://github.com/google/wire) 依赖注入 [使用介绍](https://lailin.xyz/post/go-training-week4-wire.html)
4. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录
```
#Prometheus 支票查看配置如下：
monitor:
  enable: true
  #host: 127.0.0.1 #只允许本机访问,注释可所有主机访问、docker容器中需注释
  host:
  port: 6013
  
#2、访问:    http://localhost:6013/metrics
```
5. 支持 [Swagger](https://github.com/swaggo/swag) 接口文档生成 [使用示例](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
```
   1、执行命令：make doc 即可生成swagger 文档；
   2、访问: http://localhost:3013//goodsCenterLogic/swagger/index.html 可查看swagger文档
```
6. 支持 [OpenApi Generator](https://openapi-generator.tech/docs/generators) 接口文档生成go等语言sdk
7. 支持 RESTful API 返回值规范
8. 支持 [Jaeger](https://www.jaegertracing.io/)  链路追踪
```
#jaeger 指标采集配置
xTrace:
  SamplerType: const
  SamplerParam: 1
  ReporterLocalAgentHostPort: '127.0.0.1:6831'
```
10. 支持 [zap](https://go.uber.org/zap) 日志收集、traceId 日志链路
```
#日志配置 在 env.yaml中
xLogger:
  serviceName: *name
  level: debug
  LevelColor: true
  format: plain  #日志格式，支持 plain 和 json
#  LevelColor: false
#  format: json
  stdout: false #是否输出到控制台
  file:
    filename: /apps/log/app/gin-api-http/app.log #日志文件路径
    maxSize: '500' #单个文件最大M
    maxDays: 3  #做多保留多少天
    compress: true #是否开启压缩
```
11. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
12. 支持 [viper](https://github.com/spf13/viper) 配置文件解析、支持配置文件热更新
13. 支持 [gorm](https://gorm.io/gorm) 数据库组件 v1 v2都支
14. 支持 [gorm-gen](https://github.com/go-gorm/gen) **生成数据表 CURD方法 等**代码生成器
15. 支持 [go-redis](https://github.com/redis/go-redis) 组件 
16. 支持 [sling](https://github.com/dghubble/sling) http client组件\支持重试 退避等
17. 支持 [asynq](https://github.com/hibiken/asynq)  异步任务框架、同时支持cron任务管理;[使用教程](https://www.tizi365.com/topic/14001.html)
18. 支持 [snowflake](https://github.com/bwmarrin/snowflake)   [gosnowflake](https://github.com/snowflakedb/gosnowflake) 雪花ID生成器
19. 支持 [bigcache](https://github.com/allegro/bigcache)  | [剖析Golang Bigcache的极致性能优化
    ](https://www.cyhone.com/articles/bigcache/)  | [使用教程](https://www.cyhone.com/articles/bigcache/) | [bigcache/freecache/fastcache对比](https://zhuanlan.zhihu.com/p/487455942)
```
#asynq 异步任务配置
asynq:
  # asynq monitoring 监控dashbord服务端口
  port: 7013
  #是否开启 数据面板服务
  enable: true
  #是否开启调度注册; 当多台机器时会重复注册，导致任务执行多次；多个机器只需要开启一台服务的注册即可；执行任务时分布式的
  enableRegisterSched: true
  #动态调度任务 配置文件路径
  periodicTaskConfig: '/apps/conf/gin-api-http/periodic_task_config.yaml'
  # 是否开启动态任务调度 注册,分布式服务保持一台可以调度即可； 否则出现重复任务
  enablePeriodicTaskSched: true
#dashbord访问地址: http://localhost:7013/monitoring
```


## 接口文档
Chrome浏览器可通过装  [插件](https://chromewebstore.google.com/detail/swagger-preview/jmgefnoecdghkgopjefhfjiaeckpcobk) swagger-preview 直接预览 swagger | openapi 文件 </br>
🌍  [markdown-文档](./gen/go/README.md) </br>
🌍  [swagger-文档](./gen/swagger/swagger.json) </br>
🌍  [openapi-文档](./gen/go/api/openapi.yaml) </br>
  

## 运行
```
#启动服务, envPath 动态参数可指定运行的配置文件，不指定默认 env.yaml
go run cmd/server/main.go -envPath=test

#数据库配置文件：
将目录config下的配置文件 db_goods_center.sql 导入配置文件的数据库中

# gorm-gen 自动生成 生成数据表实体 CURD
go run cmd/gen/main.go

#gorm 使用示例如下接口有使用
http://localhost:3013//goodsCenterLogic/v1/userScore/list

```



## 其他
配置文件：env.yaml

##todo 支持
1. [casbin](https://casbin.org/zh/) 权限管理; 一个支持如ACL, RBAC, ABAC等访问模型
2. 其他...



## 联系作者
    smganoabc@163.com