## 友情提示

> 1. **快速体验项目**：[在线访问地址](http://chat.example.com)。

## 项目介绍

`chat` 是一套仿微信ui的即时通讯学习项目，项目出自 [uni-app实战仿微信app开发](https://study.163.com/course/introduction/1209487898.htm)
，学习中，用 `golang + vue` 微服务思想，重写了整个项目，功能点如下
![功能点](./deploy/img/app.png)

## 系统架构图

![系统架构图](./deploy/img/system.jpeg)
参考：goim [项目地址](https://github.com/Terry-Mao/goim) [文章地址](https://zhuanlan.zhihu.com/p/128941542)

## 📗 目录结构

```lua
├── Makefile                     # 项目管理文件
├── admin                        # 管理后台
├── app                          # 业务目录
│   ├── chat                     # chat核心逻辑业务层
│   ├── connect                  # 连接层，管理http,ws,tcp连接
│   ├── constvar                 # 常量定义
│   ├── task                     # 任务层,处理消息队列消息
├── cmd                          # 支持的命令
├── config                       # 配置文件统一存放目录
├── deploy                       # 部署相关
├── dict                         # 敏感词
├── docs                         # 框架相关文档
├── frontend                     # 前端界面
├── internal                     # 公共业务目录
├── logs                         # 存放日志的目录
├── pkg                          # 公共的 package
├── proto                        # 协议
└── target                       # 运行时目录
├── main.go                      # 项目入口文件
```

### 后端技术

- 脚手架 [snake](https://github.com/1024casts/snake)
- 消息推送架构 [goim](https://github.com/Terry-Mao/goim)
- 轻量级并发服务器框架 [zinx](https://github.com/aceld/zinx)
- http框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- grpc框架 [grpc](https://google.golang.org/grpc)
  consul服务注册中心 [consul](https://github.com/hashicorp/consul)
- websocket使用 [Websocket](https://github.com/gorilla/websocket)
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://gorm.io)
- 命令行工具 [Cobra](https://github.com/spf13/cobra)
- 文档使用 [Swagger](https://swagger.io/) 生成
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 使用 [JWT](https://jwt.io/) 进行身份鉴权认证
- 校验器使用 [validator](https://github.com/go-playground/validator)  也是 Gin 框架默认的校验器
- 任务调度 [cron](https://github.com/robfig/cron)
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 测试框架 [GoConvey](http://goconvey.co/)
- CI/CD [GitHub Actions](https://github.com/actions)
- 使用 [GolangCI-lint](https://golangci.com/) 进行代码检测
- 使用 make 来管理 Go 工程
- 使用 YAML 文件进行多环境配置

### 管理后台
- laravel框架(5.5.*) [laravel文档](https://learnku.com/docs/laravel/5.5/installation/1282)
- laravel-admin后台框架 [laravel-admin文档](https://laravel-admin.org/)

### 前端技术
- 移动端 Vue 组件库 [vant](https://youzan.github.io/vant/#/zh-CN/)
- 脚手架 [vue-cli4 vant rem 移动端框架方案](https://github.com/sunniejs/vue-h5-template)

### 开发环境

| 工具           | 版本号 | 下载                                                            |
| ------------- | ------ | ------------------------------------------------------------ |
| golang        | 1.15   | https://golang.org/dl/                                       |
| nodejs        | 14.16  | https://nodejs.org/zh-cn/download/                           |
| php           | 7.3    | https://www.php.net/downloads.php                            |
| mysql         | 5.7    | https://www.mysql.com/                                       |
| redis         | 6.0    | https://redis.io/download                                    |
| nginx         | 1.19   | http://nginx.org/en/download.html                            |
| consul        | 1.9    | https://github.com/hashicorp/consul                          |
| protobuf      | 3.14   | https://github.com/protocolbuffers/protobuf                  |
| minio         | latest | http://docs.minio.org.cn/docs/                  |
| go-fastdfs    | latest | https://github.com/sjqzhang/go-fastdfs                  |

### 项目部署

### 手动编译部署

TIPS: 需要本地安装MySQL数据库和 Redis Consul go-fastdfs
```bash
# 下载安装
git clone 

# 进入项目目录
cd chat

# 编译
make build

# 修改配置
cd target/config

# 运行
make run
```

### docker

[docker安装文档](https://docs.docker.com/engine/install/)
```
cd chat
# 1. build image: 
docker build -t chat:latest -f Dockerfile .
# 2. start: 
docker run --rm -it -p 9050:9050 -p 9070:9070 chat:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统
```

### docker-compose
[docker-compose安装文档](https://docs.docker.com/compose/install/)
```
cd chat
部署前端，记得修改frontend/src/config/env.production.js下的配置，默认本机127.0.0.1
docker-compose up -d
```
访问 [http://127.0.0.1](http://127.0.0.1)

### 连接层多开动态负载部署使用 nginx + consul-template
[文档地址](https://learn.hashicorp.com/tutorials/consul/load-balancing-nginx?in=consul/load-balancing)

## 常用命令

- make help 查看帮助
- make build 编译项目
- make gen-docs 生成接口文档
- make run 运行项目

## 📝 接口文档

- [chat接口文档](http://127.0.0.1:9050/swagger/index.html)
- [Protobuf学习](https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/)

## 开发规范

遵循: [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)