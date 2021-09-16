## 友情提示

> 1. **快速体验项目**：[在线访问地址](http://mall.example.com)。

## 项目介绍

`mall` 是一套商城全栈学习项目，项目出自 [uni-app实战商城类app和小程序](https://study.163.com/course/introduction/1209401825.htm)
，原项目使用（`uni-app + php`），此项目使用 `golang + vant` 重写了整个项目，支付使用以太币（`ethereum`）

## Usage

> 需要在浏览器中安装 Metamask 插件 https://metamask.io/

- 效果演示

## 📗 目录结构

```lua
├── Makefile                     # 项目管理文件
├── admin                        # 管理后台
├── app                          # 业务目录
│   ├── cache                    # 缓存封装
│   ├── conf                     # 配置
│   ├── constvar                 # 常量工具
│   ├── ecode                    # 错误码定义
│   ├── eth                      # 操作以太坊合约
│   ├── handler                  # http 接口
│   ├── idl                      # 数据转换输出
│   ├── middleware               # http 中间件
│   ├── model                    # 数据库model
│   ├── repo                     # 数据库访问层
│   ├── routers                  # 路由定义
│   ├── server                   # 服务
│   ├── service                  # 业务逻辑层
├── cmd                          # 命令行工具
├── config                       # 配置文件统一存放目录
├── deploy                       # 部署相关
├── docs                         # 文档
├── eth                          # 以太坊合约类
├── frontend                     # 前端界面
├── logs                         # 日志目录
├── pkg                          # 公共的 package
├── seed                         # 数据填充
├── template                     # 模板
└── target                       # 运行时目录
├── main.go                      # 项目入口文件
```

### 后端技术（golang）

- 脚手架 [snake](https://github.com/1024casts/snake)
- http框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://gorm.io)
- 以太坊客户端 [go-ethereum](https://github.com/ethereum/go-ethereum)
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

### 数据抓取（python3）

- 爬虫框架 [scrapy](https://github.com/scrapy/scrapy)

### 管理后台（php）

- [入口](./admin)
- laravel框架(5.5.*) [laravel文档](https://learnku.com/docs/laravel/5.5/installation/1282)
- laravel-admin后台框架 [laravel-admin文档](https://laravel-admin.org/)

### 前端技术（vue）

- [入口](./frontend)
- 移动端 Vue 组件库 [vant](https://youzan.github.io/vant/#/zh-CN/)
- 脚手架 [vue-cli4 vant rem 移动端框架方案](https://github.com/sunniejs/vue-h5-template)

### 开发环境

| 工具           | 版本号 | 下载                                                            |
| ------------- | ------ | ------------------------------------------------------------ |
| golang        | 1.15   | https://golang.org/dl/                                       |
| nodejs        | 14.16  | https://nodejs.org/zh-cn/download/                           |
| php           | 7.3    | https://www.php.net/downloads.php                            |
| python        | 3.9    | https://www.python.org/downloads                             |
| solidity      | 0.8.0  | http://remix.ethereum.org/                                   |
| mysql         | 5.7    | https://www.mysql.com/                                       |
| redis         | 6.0    | https://redis.io/download                                    |
| nginx         | 1.19   | http://nginx.org/en/download.html                            |
| minio         | latest | http://docs.minio.org.cn/docs/                               |

### 项目部署

### 手动编译部署

TIPS: 需要本地安装MySQL数据库和 Redis

```bash
# 下载安装
git clone 

# 进入项目目录
cd mall

# 编译
make build

# 修改配置
cd target/config

# 运行
make run
```

### docker

[docker安装文档](https://docs.docker.com/engine/install/)

```shell
cd mall
# 1. build image: 
docker build -t mall:latest -f Dockerfile .
# 2. start: 
docker run --rm -it -p 9052:9052 mall:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统
```

### docker-compose

[docker-compose安装文档](https://docs.docker.com/compose/install/)
组件清单:

- mall_frontend 前端UI
- mall_admin 管理后台
- mall_api 接口
- db mysql数据库
- redis 缓存
- minio 私有对象存储

```shell
# 部署
cd mall
# 前端项目默认api host：127.0.0.1，如需修改，请 vim frontend/src/config/env.production.js
docker-compose up -d
```

访问 [http://127.0.0.1](http://127.0.0.1)


## 常用命令

- make help 查看帮助
- make build 编译项目
- make run 运行项目
- make test 运行测试用例
- make clean 清除编译文件
- make doc 生成接口文档
- make lint 代码检查
- make graph 生成交互式的可视化Go程序调用图
- make docker 生成docker镜像，确保已安装docker
- make abi 生成以太坊合约类

## 📝 接口文档

- [chat接口文档](http://127.0.0.1:9050/swagger/index.html)
- [管理后台](http://127.0.0.1:8000)
- [前端界面](http://127.0.0.1)

## 其他

- 开发规范: [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)