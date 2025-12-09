# TodoList-DDD 项目说明

## 项目概述

这是一个基于 Go 语言和领域驱动设计（DDD）架构实现的 TodoList 备忘录项目。该项目采用清晰的分层架构，实现了用户管理和任务管理功能。

### 主要技术栈

- **语言**: Go 1.25.1
- **Web框架**: Gin
- **数据库**: MySQL (GORM)
- **缓存**: Redis
- **配置管理**: Viper
- **认证**: JWT
- **架构模式**: 领域驱动设计 (DDD)

### 项目结构

```
├── application/          # 应用层
│   ├── task/            # 任务应用服务
│   └── user/            # 用户应用服务
├── cmd/                 # 命令行入口
├── conf/                # 配置文件
├── consts/              # 常量定义
├── domain/              # 领域层
│   ├── task/            # 任务领域
│   └── user/            # 用户领域
├── infrastructure/      # 基础设施层
│   ├── auth/            # 认证相关
│   ├── container/       # 依赖注入容器
│   ├── encrypt/         # 加密工具
│   └── persistence/     # 数据持久化
└── interfaces/          # 接口层
    ├── controller/      # 控制器
    ├── initilaize/      # 初始化
    └── types/           # 类型定义
```

## 构建和运行

### 环境准备

1. 启动依赖服务（MySQL 和 Redis）：
```bash
docker-compose up -d
```

2. 确保数据库配置正确（参考 conf/config.yaml）：
   - MySQL: 端口 3307
   - Redis: 端口 6379

### 运行项目

```bash
# 运行主程序
go run cmd/main.go

# 或者构建后运行
go build -o todolist cmd/main.go
./todolist
```

服务将在 `:8080` 端口启动。

### 开发相关命令

```bash
# 格式化代码
go fmt ./...

# 检查代码
go vet ./...

# 运行测试（如果有）
go test ./...
```

## 核心功能

### 用户模块
- 用户注册
- 用户登录
- JWT 认证

### 任务模块
- 创建任务（已实现）
- 任务列表（已实现）
- 任务详情（已实现）
- 任务更新（已实现）
- 任务搜索（已实现）
- 任务删除（已实现）

## 开发约定

### 代码结构
- 严格遵循 DDD 分层架构
- 每个领域包含 entity、repository、service 三个子包
- 应用层负责业务流程编排
- 基础设施层负责技术实现细节

### 命名规范
- 包名使用小写字母
- 接口名以大写字母开头
- 实现体以 Impl 结尾
- 常量使用全大写字母

### 数据流转
1. Controller 接收 HTTP 请求
2. Types 层进行数据转换
3. Application 层处理业务逻辑
4. Domain 层定义核心业务规则
5. Infrastructure 层处理数据持久化

## 配置说明

主要配置文件位于 `conf/config.yaml`，包含：
- 服务器配置（端口、版本、JWT密钥）
- MySQL 数据库配置
- Redis 缓存配置

## 项目状态

这是一个学习项目，基于 [CocaineCong/todolist-ddd](https://github.com/CocaineCong/todolist-ddd) 进行开发。用户模块基本完成，任务模块的部分功能仍在开发中。