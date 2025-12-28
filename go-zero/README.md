# Go-Zero 微服务用户管理系统

基于 go-zero 框架构建的微服务用户管理系统，用于学习微服务架构的核心概念和实践。

## 目录

- [项目介绍](#项目介绍)
- [技术栈](#技术栈)
- [系统架构](#系统架构)
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [API文档](#api文档)
- [配置说明](#配置说明)
- [学习指南](#学习指南)

---

## 项目介绍

本项目是一个完整的微服务架构示例，包含以下核心功能：

- **用户管理**：用户注册、登录、信息查询与更新
- **身份认证**：JWT Token 生成、验证、刷新、注销
- **角色权限**：角色 CRUD、用户角色分配、RBAC 基础实现

项目采用典型的微服务分层架构，API 网关负责接收 HTTP 请求，通过 gRPC 协议调用后端 RPC 服务。

## 技术栈

| 组件 | 技术选型 | 说明 |
|------|----------|------|
| 框架 | go-zero v1.9 | 高性能微服务框架 |
| API网关 | go-zero rest | HTTP RESTful 服务 |
| RPC通信 | go-zero zrpc | 基于 gRPC 的 RPC 框架 |
| 数据库 | MySQL 8.0 | 关系型数据库 |
| 缓存 | Redis 7 | 内存缓存，Token 存储 |
| 认证 | JWT | JSON Web Token |
| 代码生成 | goctl | go-zero 官方 CLI 工具 |

## 系统架构

```
                         ┌────────────────────────────────────┐
                         │            Client                  │
                         │      (Web/Mobile/Postman)          │
                         └──────────────┬─────────────────────┘
                                        │ HTTP
                                        ▼
                         ┌────────────────────────────────────┐
                         │         user-api (网关)            │
                         │           :8888                    │
                         │  ┌─────────────────────────────┐   │
                         │  │    JWT 认证中间件           │   │
                         │  │    路由分发                 │   │
                         │  │    参数校验                 │   │
                         │  └─────────────────────────────┘   │
                         └───────┬──────────┬─────────┬───────┘
                                 │          │         │
                    gRPC         │          │         │        gRPC
             ┌───────────────────┘          │         └───────────────────┐
             │                              │                             │
             ▼                              ▼                             ▼
┌────────────────────────┐   ┌────────────────────────┐   ┌────────────────────────┐
│     user-rpc           │   │     auth-rpc           │   │     role-rpc           │
│       :9001            │   │       :9002            │   │       :9003            │
│  ┌──────────────────┐  │   │  ┌──────────────────┐  │   │  ┌──────────────────┐  │
│  │ 用户注册/登录    │  │   │  │ Token 生成       │  │   │  │ 角色 CRUD        │  │
│  │ 用户信息 CRUD    │  │   │  │ Token 验证       │  │   │  │ 用户角色分配     │  │
│  │ 密码加密验证     │  │   │  │ Token 刷新/注销  │  │   │  │ 权限查询         │  │
│  └──────────────────┘  │   │  └──────────────────┘  │   │  └──────────────────┘  │
└───────────┬────────────┘   └───────────┬────────────┘   └───────────┬────────────┘
            │                            │                            │
            ▼                            ▼                            ▼
     ┌──────────┐                 ┌──────────┐                 ┌──────────┐
     │  MySQL   │                 │  Redis   │                 │  MySQL   │
     │  users   │                 │  tokens  │                 │  roles   │
     └──────────┘                 └──────────┘                 └──────────┘
```

### 服务职责

| 服务 | 端口 | 职责 |
|------|------|------|
| user-api | 8888 | HTTP 网关，JWT 认证，路由转发 |
| user-rpc | 9001 | 用户核心业务：注册、登录、CRUD |
| auth-rpc | 9002 | 认证服务：Token 生成、验证、刷新 |
| role-rpc | 9003 | 角色服务：角色管理、用户角色绑定 |

## 项目结构

```
go-zero/
├── user-api/                    # API 网关服务
│   ├── etc/
│   │   └── user-api.yaml        # 配置文件
│   ├── internal/
│   │   ├── config/              # 配置结构体
│   │   ├── handler/             # HTTP 处理器
│   │   ├── logic/               # 业务逻辑层
│   │   ├── svc/                 # 服务上下文（依赖注入）
│   │   └── types/               # 请求/响应类型
│   ├── user.api                 # API 定义文件
│   └── user.go                  # 入口文件
│
├── user-rpc/                    # 用户 RPC 服务
│   ├── etc/
│   │   └── user.yaml
│   ├── internal/
│   │   ├── config/
│   │   ├── logic/               # RPC 业务逻辑
│   │   ├── model/               # 数据模型（MySQL 操作）
│   │   ├── server/              # gRPC Server 实现
│   │   └── svc/
│   ├── pb/                      # protobuf 生成的 Go 代码
│   ├── user/                    # RPC Client 封装
│   ├── user.proto               # protobuf 定义文件
│   └── user.go
│
├── auth-rpc/                    # 认证 RPC 服务
│   ├── etc/
│   │   └── auth.yaml
│   ├── internal/
│   │   ├── config/
│   │   ├── logic/
│   │   ├── server/
│   │   └── svc/
│   ├── pb/
│   ├── auth/
│   ├── auth.proto
│   └── auth.go
│
├── role-rpc/                    # 角色 RPC 服务
│   ├── etc/
│   │   └── role.yaml
│   ├── internal/
│   │   ├── config/
│   │   ├── logic/
│   │   ├── model/
│   │   ├── server/
│   │   └── svc/
│   ├── pb/
│   ├── role/
│   ├── role.proto
│   └── role.go
│
├── common/                      # 公共模块
│   ├── errorx/                  # 统一错误码定义
│   │   └── baseerror.go
│   ├── response/                # 统一响应格式
│   │   └── response.go
│   └── utils/                   # 工具函数
│       ├── jwt.go               # JWT 工具
│       └── password.go          # 密码加密工具
│
├── deploy/                      # 部署配置
│   ├── sql/
│   │   └── init.sql             # 数据库初始化脚本
│   └── docker-compose.yaml      # Docker 编排
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 快速开始

### 环境要求

- Go 1.21+
- Docker & Docker Compose
- goctl（go-zero CLI 工具）

### 1. 安装 goctl

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### 2. 克隆项目

```bash
cd /path/to/your/workspace
git clone <repository-url>
cd go-zero
```

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 启动基础设施

```bash
# 启动 MySQL 和 Redis
make docker-up

# 等待服务启动（约 10 秒）
# MySQL 会自动执行 deploy/sql/init.sql 初始化数据库
```

### 5. 启动微服务

需要在 **4 个不同的终端** 中分别运行：

```bash
# 终端 1：启动 user-rpc
make run-user-rpc

# 终端 2：启动 auth-rpc
make run-auth-rpc

# 终端 3：启动 role-rpc
make run-role-rpc

# 终端 4：启动 user-api
make run-user-api
```

### 6. 验证服务

```bash
# 测试注册接口
curl -X POST http://localhost:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456","email":"test@example.com"}'

# 测试登录接口
curl -X POST http://localhost:8888/api/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

### 7. 停止服务

```bash
# 停止基础设施
make docker-down

# 在各终端中按 Ctrl+C 停止微服务
```

## API文档

### 基础信息

- **Base URL**: `http://localhost:8888/api`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token（在 Header 中添加 `Authorization: Bearer <token>`）

### 响应格式

所有接口返回统一格式：

```json
{
  "code": 0,
  "msg": "success",
  "data": { ... }
}
```

### 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 未授权 |
| 1005 | 内部错误 |
| 2001 | 用户不存在 |
| 2002 | 用户已存在 |
| 2003 | 密码错误 |
| 2004 | 用户已禁用 |
| 3001 | Token 无效 |
| 3002 | Token 已过期 |
| 4001 | 角色不存在 |
| 4002 | 角色已存在 |

---

### 用户接口

#### 1. 用户注册

**POST** `/user/register`

**请求参数**：

```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "phone": "13800138000"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名，唯一 |
| password | string | 是 | 密码 |
| email | string | 否 | 邮箱 |
| phone | string | 否 | 手机号 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1,
    "username": "testuser"
  }
}
```

---

#### 2. 用户登录

**POST** `/user/login`

**请求参数**：

```json
{
  "username": "testuser",
  "password": "123456"
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
    "expiresAt": 1703836800
  }
}
```

---

#### 3. 获取用户信息

**GET** `/user/info`

**需要认证**：是

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "phone": "13800138000",
      "avatar": "",
      "status": 1,
      "createdAt": 1703750400
    }
  }
}
```

---

#### 4. 更新用户信息

**PUT** `/user/update`

**需要认证**：是

**请求参数**：

```json
{
  "email": "newemail@example.com",
  "phone": "13900139000",
  "avatar": "https://example.com/avatar.jpg"
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "success": true
  }
}
```

---

#### 5. 删除用户

**DELETE** `/user/:id`

**需要认证**：是

**路径参数**：

| 参数 | 类型 | 说明 |
|------|------|------|
| id | int64 | 用户 ID |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "success": true
  }
}
```

---

#### 6. 用户列表

**GET** `/user/list`

**需要认证**：是

**查询参数**：

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| page | int64 | 1 | 页码 |
| pageSize | int64 | 10 | 每页条数 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "users": [
      {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "phone": "",
        "avatar": "",
        "status": 1,
        "createdAt": 1703750400
      }
    ],
    "total": 1
  }
}
```

---

### 认证接口

#### 1. 刷新 Token

**POST** `/auth/refresh`

**请求参数**：

```json
{
  "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
    "expiresAt": 1703836800
  }
}
```

---

#### 2. 退出登录

**POST** `/auth/logout`

**需要认证**：是

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "success": true
  }
}
```

---

### 角色接口

#### 1. 创建角色

**POST** `/role/create`

**需要认证**：是

**请求参数**：

```json
{
  "name": "编辑员",
  "code": "editor",
  "description": "内容编辑权限"
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 3
  }
}
```

---

#### 2. 角色列表

**GET** `/role/list`

**需要认证**：是

**查询参数**：

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| page | int64 | 1 | 页码 |
| pageSize | int64 | 10 | 每页条数 |

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "roles": [
      {
        "id": 1,
        "name": "管理员",
        "code": "admin",
        "description": "系统管理员，拥有所有权限",
        "status": 1,
        "createdAt": 1703750400
      },
      {
        "id": 2,
        "name": "普通用户",
        "code": "user",
        "description": "普通用户，拥有基本权限",
        "status": 1,
        "createdAt": 1703750400
      }
    ],
    "total": 2
  }
}
```

---

#### 3. 更新角色

**PUT** `/role/:id`

**需要认证**：是

**请求参数**：

```json
{
  "name": "超级管理员",
  "description": "拥有最高权限",
  "status": 1
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "success": true
  }
}
```

---

#### 4. 分配角色

**POST** `/role/assign`

**需要认证**：是

**请求参数**：

```json
{
  "userId": 1,
  "roleIds": [1, 2]
}
```

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "success": true
  }
}
```

---

#### 5. 获取用户角色

**GET** `/role/user`

**需要认证**：是

**响应示例**：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "roles": [
      {
        "id": 1,
        "name": "管理员",
        "code": "admin",
        "description": "系统管理员",
        "status": 1,
        "createdAt": 1703750400
      }
    ]
  }
}
```

---

## 配置说明

### user-api 配置

```yaml
# user-api/etc/user-api.yaml
Name: user-api
Host: 0.0.0.0
Port: 8888

# JWT 配置
Auth:
  AccessSecret: "go-zero-user-access-secret-key"  # JWT 密钥
  AccessExpire: 7200                               # 过期时间（秒）

# RPC 服务连接配置
UserRpc:
  Target: 127.0.0.1:9001    # user-rpc 地址
  NonBlock: true
  Timeout: 5000

AuthRpc:
  Target: 127.0.0.1:9002    # auth-rpc 地址
  NonBlock: true
  Timeout: 5000

RoleRpc:
  Target: 127.0.0.1:9003    # role-rpc 地址
  NonBlock: true
  Timeout: 5000
```

### user-rpc 配置

```yaml
# user-rpc/etc/user.yaml
Name: user.rpc
ListenOn: 0.0.0.0:9001

# MySQL 配置
Mysql:
  DataSource: root:123456@tcp(127.0.0.1:3306)/go_zero_user?charset=utf8mb4&parseTime=True&loc=Local

# Redis 缓存配置
CacheRedis:
  - Host: 127.0.0.1:6379
    Type: node
```

### auth-rpc 配置

```yaml
# auth-rpc/etc/auth.yaml
Name: auth.rpc
ListenOn: 0.0.0.0:9002

# JWT 配置
Jwt:
  AccessSecret: "go-zero-user-access-secret-key"
  AccessExpire: 7200      # Access Token 2小时
  RefreshSecret: "go-zero-user-refresh-secret-key"
  RefreshExpire: 604800   # Refresh Token 7天

# Redis 配置
Redis:
  Host: 127.0.0.1:6379
  Type: node
  Pass: ""
```

### role-rpc 配置

```yaml
# role-rpc/etc/role.yaml
Name: role.rpc
ListenOn: 0.0.0.0:9003

# MySQL 配置
Mysql:
  DataSource: root:123456@tcp(127.0.0.1:3306)/go_zero_user?charset=utf8mb4&parseTime=True&loc=Local
```

---

## 数据库设计

### users 表

```sql
CREATE TABLE users (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) DEFAULT '',
    phone VARCHAR(20) DEFAULT '',
    avatar VARCHAR(255) DEFAULT '',
    status TINYINT UNSIGNED DEFAULT 1,  -- 0:禁用 1:启用
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);
```

### roles 表

```sql
CREATE TABLE roles (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255) DEFAULT '',
    status TINYINT UNSIGNED DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status)
);
```

### user_roles 表

```sql
CREATE TABLE user_roles (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    role_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_role (user_id, role_id),
    INDEX idx_user_id (user_id),
    INDEX idx_role_id (role_id)
);
```

---

## 学习指南

### 1. goctl 代码生成

go-zero 提供 goctl 工具自动生成代码框架：

```bash
# 根据 .proto 文件生成 RPC 代码
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# 根据 .api 文件生成 HTTP API 代码
goctl api go -api user.api -dir .
```

**核心文件**：
- `*.proto` - 定义 RPC 服务接口
- `*.api` - 定义 HTTP API 接口

### 2. 服务间通信

API 网关通过 gRPC 调用 RPC 服务：

```go
// user-api/internal/svc/servicecontext.go
type ServiceContext struct {
    Config   config.Config
    UserRpc  user.User      // RPC 客户端
    AuthRpc  auth.Auth
    RoleRpc  role.Role
}

// user-api/internal/logic/loginlogic.go
func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
    // 1. 调用 user-rpc 验证用户
    userResult, err := l.svcCtx.UserRpc.Login(l.ctx, &userpb.LoginReq{...})

    // 2. 调用 auth-rpc 生成 Token
    tokenResult, err := l.svcCtx.AuthRpc.GenerateToken(l.ctx, &authpb.GenerateTokenReq{...})

    return &types.LoginResp{...}, nil
}
```

### 3. JWT 认证流程

```
┌──────────┐     POST /login      ┌──────────┐
│  Client  │ ──────────────────▶  │ user-api │
└──────────┘                      └────┬─────┘
                                       │
     ┌─────────────────────────────────┼─────────────────────────────────┐
     │                                 ▼                                 │
     │  ┌──────────┐  验证用户   ┌──────────┐  生成Token   ┌──────────┐  │
     │  │ user-rpc │ ◀────────  │          │ ────────────▶ │ auth-rpc │  │
     │  └──────────┘            │          │               └──────────┘  │
     │                          │ user-api │                             │
     │                          │          │                             │
     │                          └────┬─────┘                             │
     │                               │                                   │
     └───────────────────────────────┼───────────────────────────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │   返回 Token  │
                              │ accessToken  │
                              │ refreshToken │
                              └──────────────┘
```

### 4. 分层架构

```
Handler (HTTP处理器)
    ↓
Logic (业务逻辑)
    ↓
ServiceContext (服务上下文/依赖注入)
    ↓
Model (数据访问层) / RPC Client (远程调用)
```

### 5. 统一错误处理

```go
// common/errorx/baseerror.go
type CodeError struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}

func NewCodeError(code int) *CodeError {
    return &CodeError{Code: code, Msg: codeMsg[code]}
}

// 使用示例
if err != nil {
    return nil, errorx.NewCodeError(errorx.CodeUserNotFound)
}
```

---

## Makefile 命令

| 命令 | 说明 |
|------|------|
| `make tidy` | 整理 Go 依赖 |
| `make build` | 构建所有服务 |
| `make docker-up` | 启动 MySQL 和 Redis |
| `make docker-down` | 停止 MySQL 和 Redis |
| `make run-user-rpc` | 运行用户 RPC 服务 |
| `make run-auth-rpc` | 运行认证 RPC 服务 |
| `make run-role-rpc` | 运行角色 RPC 服务 |
| `make run-user-api` | 运行 API 网关服务 |
| `make clean` | 清理构建产物 |
| `make help` | 显示帮助信息 |

---

## 扩展建议

1. **服务注册发现**：集成 Etcd 实现服务注册与发现
2. **链路追踪**：集成 Jaeger 实现分布式链路追踪
3. **监控告警**：集成 Prometheus + Grafana
4. **日志收集**：集成 ELK Stack
5. **API 文档**：集成 Swagger
6. **限流熔断**：使用 go-zero 内置的限流和熔断功能
7. **消息队列**：集成 Kafka 或 RabbitMQ 实现异步处理

---

## 参考资源

- [go-zero 官方文档](https://go-zero.dev/)
- [go-zero GitHub](https://github.com/zeromicro/go-zero)
- [goctl 使用指南](https://go-zero.dev/docs/tutorials)
