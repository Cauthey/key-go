# 防火墙

> 基于 GIN + Gorm 2.0 + CASBIN + WIRE 实现的RBAC权限管理脚手架，目标是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。

## 特性

- 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
- 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestLogger、RequestRateLimiter、TraceID、CasbinEnforce、Recover、GZIP）
- 基于 `Casbin` 的 RBAC 访问控制模型 -- **权限控制可以细粒度到按钮 & 接口**
- 基于 `Gorm 2.0` 的数据库访问层 - 全功能 ORM
- 基于 `WIRE` 的依赖注入 -- 依赖注入本身的作用是解决了各个模块间层级依赖繁琐的初始化过程
- 基于 `Logrus & Context` 实现了日志输出，通过结合 Context 实现了统一的 TraceID/UserID 等关键字段的输出(同时支持日志钩子写入到`Gorm`)
- 基于 `JWT` 的用户认证 -- 基于 JWT 的黑名单验证机制
- 基于 `Swaggo` 自动生成 `Swagger` 文档 -- 独立于接口的 mock 实现
- 基于 `net/http/httptest` 标准包实现了 API 的单元测试
- 基于 `go mod` 的依赖管理(国内源可使用：<https://goproxy.cn/>)

## 依赖工具

```
go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag
```

- [air](https://github.com/cosmtrek/air) -- Live reload for Go apps
- [wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go
- [swag](https://github.com/swaggo/swag) -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## 依赖框架

- [Gin](https://gin-gonic.com/) -- The fastest full-featured web framework for Go.
- [GORM](https://gorm.io/) -- The fantastic ORM library for Golang
- [Casbin](https://casbin.org/) -- An authorization library that supports access control models like ACL, RBAC, ABAC in Golang
- [Wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go

## 快速开始


```bash
$ git clone 

$ cd key-go

# 基于Makefile运行
$ make start

# OR 使用go命令运行
$ go run cmd/key-go/main.go web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml
```

> 启动成功之后，可在浏览器中输入地址进行访问：[http://127.0.0.1:10088](http://127.0.0.1:10088)

## 生成`swagger`文档

```bash
# 基于Makefile
make swagger

# OR 使用swag命令
swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/app/swagger
```

## 生成依赖注入文件

```bash
# 基于Makefile
make wire

# OR 使用wire命令
wire gen ./internal/app
```

#### 创建模板 task.yaml

```yaml
name: Task
comment: 任务管理
fields:
  - name: Code
    type: string
    comment: 任务编号
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Name
    type: string
    comment: 任务名称
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Memo
    type: string
    comment: 任务备注
    required: false
    binding_options: ""
    gorm_options: "size:1024;"
```

#### 执行生成命令并运行

```bash
cd test-gin-admin
gin-admin-cli g -d . -p test-gin-admin -f ./task.yaml

# 生成 swagger
make swagger

# 生成依赖项
make wire

# 运行服务
make start
```

## 前端工程

- 基于[Ant Design React](https://ant.design/index-cn)版本的实现：[key-go-react](https://github.com/gin-admin/gin-admin-react)

## Questions

### OSX 下基于 `sqlite3` 运行出现：`'stdlib.h' file not found`

```bash
export CGO_ENABLED=1; export CC=gcc; go get -v -x github.com/mattn/go-sqlite3
```

## 目录结构

```text
├── cmd
│   └── key-go
│       └── main.go       # 入口文件
├── configs
│   ├── config.toml       # 配置文件
│   ├── menu.yaml         # 菜单初始化配置
│   └── model.conf        # casbin 策略配置
│   └── config.xml        # 系统全局配置文件
├── docs                  # 文档
├── internal
│   └── app
│       ├── api           # API 处理层
│       ├── config        # 配置文件映射
│       ├── contextx      # 统一上下文处理
│       ├── dao           # 数据访问层
│       ├── ginx          # gin 扩展模块
│       ├── middleware    # gin 中间件模块
│       ├── module        # 通用业务处理模块
│       ├── router        # 路由层
│       ├── schema        # 统一入参、出参对象映射
│       ├── service       # 业务逻辑层
│       ├── swagger       # swagger 生成文件
│       ├── test          # 模块单元测试
├── pkg
│   ├── auth              
│   │   └── jwtauth       # jwt 认证模块
│   ├── errors            # 错误处理模块
│   ├── gormx             # gorm 扩展模块
│   ├── logger            # 日志模块
│   │   ├── hook
│   └── util              # 工具包
│       ├── conv         
│       ├── hash         
│       ├── json
│       ├── snowflake
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts               # 统一处理脚本
```

## MIT License

    Copyright (c) 2021 Lyric
