```
           _                  __                                             _    
  ___  ___| |__   ___        / _|_ __ __ _ _ __ ___   _____      _____  _ __| | __
 / _ \/ __| '_ \ / _ \ _____| |_| '__/ _` | '_ ` _ \ / _ \ \ /\ / / _ \| '__| |/ /
|  __/ (__| | | | (_) |_____|  _| | | (_| | | | | | |  __/\ V  V / (_) | |  |   < 
 \___|\___|_| |_|\___/      |_| |_|  \__,_|_| |_| |_|\___| \_/\_/ \___/|_|  |_|\_\
                                                                                  
```

[echo-framework](https://github.com/nelsonkti/echo-framework) 是基于 echo 搭建用于快速开发的项目框架

## 安装
```
go get -u github.com/nelsonkti/echo-framework
```

## 功能叙述
- 支持 gorm、logger 日志、jwt、cron 定时任务、redis 等
- mysql 数据库读写分离、 负载均衡
- socket.io 通信协议
- nsq 消息队列
- 分布式部署

## 文件夹结构 
- config 文件配置和初始化配置数据
- cron 定时任务
- lib 日常使用的库
- logic logic 业务逻辑
- main 程序启动入口, 主要可以启动 http
- routes 包含应用的所有路由定义
- socket 通信相关的代码，以 socket.io 通信协议为主


- logic 目录
    - http 包含了控制器、中间件以及表单请求、验证器等
       - controllers 控制器层
       - middleware 中间件
       - models 模型层
       - repository 业务层调用数据访问层
       - responses 返回层
       - services 服务层主要处理业务逻辑
       - validators 表单验证器
      
    - mq nsq 生产者和消费者
       - producer 生产者


## 项目介绍
#### 项目默认支持 `nsq`、`memcache`、`redis`，如果不需要，可以再 `main` 文件夹下 注释以下代码

`memcache` 连接
```
//连接 memcache
db.ConnectMemcache(config.Memcache)
```

`redis` 连接
```
//连接redis
db.ConnectRedis(config.RedisIP, config.RedisPassword, 0, "default")
```

`nsq` 连接
```
//连接redis
go func() {
    defer helper.RecoverPanic()
    //producer.StartNsqProducer(config.NSQIP)
    mq.StartNsqServer(config.NSQIP, config.NSQConsumers)
}()
```

`cron` 本地默认不启动， 需要启动，去掉`if`就可以了
```
//启动定时任务
if config.Env != "local" {
    cron.RegisterCrons(config.RedisIP, config.RedisPassword)
}
```

`grom` 读写分离 [DBResolver](https://gorm.io/zh_CN/docs/dbresolver.html)
```
// 使用 Write 模式
User.Model().Clauses(dbresolver.Write).First(&User)
```

运行logic
```
cd main
go run logic.go
```

运行socket
```
cd main
go run socket.go
```

## 环境要求 

- go >= 1.13

