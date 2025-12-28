# Wire + Gin 依赖注入示例项目

本项目演示如何在 Gin Web 框架中使用 Google Wire 进行依赖注入。

## 目录

- [什么是依赖注入](#什么是依赖注入)
- [什么是 Wire](#什么是-wire)
- [项目结构](#项目结构)
- [核心概念](#核心概念)
- [代码详解](#代码详解)
- [运行项目](#运行项目)
- [API 测试](#api-测试)
- [进阶用法](#进阶用法)

---

## 什么是依赖注入

依赖注入（Dependency Injection，DI）是一种设计模式，核心思想是：**对象不应该自己创建它所依赖的对象，而是由外部注入**。

### 没有依赖注入的代码

```go
type UserService struct {
    repo *UserRepository  // 直接依赖具体实现
}

func NewUserService() *UserService {
    return &UserService{
        repo: NewUserRepository(),  // 自己创建依赖，耦合度高
    }
}
```

**问题**：
- `UserService` 与 `UserRepository` 的具体实现紧耦合
- 难以进行单元测试（无法 mock）
- 难以替换实现

### 使用依赖注入的代码

```go
type UserService struct {
    repo UserRepository  // 依赖接口
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{
        repo: repo,  // 依赖由外部注入
    }
}
```

**优点**：
- 松耦合，依赖接口而非实现
- 易于测试，可以注入 mock 对象
- 易于替换实现

---

## 什么是 Wire

[Wire](https://github.com/google/wire) 是 Google 开源的 Go 依赖注入代码生成器。

### Wire 的特点

| 特性 | 说明 |
|------|------|
| **编译时注入** | 在编译时生成代码，无运行时反射开销 |
| **类型安全** | 编译时检查依赖关系，错误提前发现 |
| **代码生成** | 自动生成依赖注入代码，减少样板代码 |
| **易于调试** | 生成的代码是普通 Go 代码，可以直接阅读和调试 |

### Wire vs 其他 DI 框架

| 框架 | 类型 | 特点 |
|------|------|------|
| **Wire** | 编译时 | 代码生成，零运行时开销 |
| **dig** (Uber) | 运行时 | 基于反射，更灵活 |
| **fx** (Uber) | 运行时 | 基于 dig，提供应用生命周期管理 |

---

## 项目结构

```
go-wire/
├── main.go              # 应用入口
├── wire.go              # Wire 依赖定义（核心文件）
├── wire_gen.go          # Wire 生成的代码（自动生成，勿手动修改）
├── go.mod               # Go 模块定义
├── go.sum               # 依赖校验
│
├── config/
│   └── config.go        # 应用配置
│
├── model/
│   └── user.go          # 数据模型
│
├── repository/
│   └── user.go          # 数据访问层（DAO）
│
├── service/
│   └── user.go          # 业务逻辑层
│
└── handler/
    └── user.go          # HTTP 处理器（Controller）
```

### 分层架构

```
┌─────────────────────────────────────────────────────────┐
│                      HTTP Request                        │
└─────────────────────────┬───────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────┐
│                    Handler 层                            │
│              处理 HTTP 请求/响应                         │
│              参数校验、错误处理                          │
└─────────────────────────┬───────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────┐
│                    Service 层                            │
│                  业务逻辑处理                            │
│              事务管理、业务规则                          │
└─────────────────────────┬───────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────┐
│                   Repository 层                          │
│                   数据访问操作                           │
│              CRUD、数据库查询                            │
└─────────────────────────┬───────────────────────────────┘
                          ▼
┌─────────────────────────────────────────────────────────┐
│                    Data Store                            │
│              数据库 / 缓存 / 文件                        │
└─────────────────────────────────────────────────────────┘
```

### 依赖注入链

```
Config ─────────────────────────────────────┐
                                            │
Repository ◄── Service ◄── Handler ◄── App ─┘
     │              │           │
     └──────────────┴───────────┴──── Wire 自动注入
```

---

## 核心概念

### 1. Provider（提供者）

Provider 是一个普通的 Go 函数，用于创建某个类型的实例。

```go
// config/config.go
// NewConfig 是一个 Provider，返回 *Config
func NewConfig() *Config {
    return &Config{
        Server: ServerConfig{Port: ":8080"},
    }
}

// repository/user.go
// NewUserRepository 是一个 Provider，返回 UserRepository 接口
func NewUserRepository() UserRepository {
    return &userRepository{
        users:  make(map[int64]*model.User),
        nextID: 1,
    }
}

// service/user.go
// NewUserService 是一个 Provider，依赖 UserRepository
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}
```

**Provider 的规则**：
- 函数名通常以 `New` 开头
- 参数是该类型的依赖
- 返回值是创建的实例（可以返回 error 作为第二个返回值）

### 2. ProviderSet（提供者集合）

使用 `wire.NewSet()` 将多个 Provider 组合在一起。

```go
// wire.go
var ProviderSet = wire.NewSet(
    config.NewConfig,           // 配置 Provider
    repository.NewUserRepository,  // Repository Provider
    service.NewUserService,     // Service Provider
    handler.NewUserHandler,     // Handler Provider
)
```

**ProviderSet 的优点**：
- 模块化组织 Provider
- 可以嵌套其他 ProviderSet
- 便于复用和测试

### 3. Injector（注入器）

Injector 是一个函数，定义了最终要创建的对象。Wire 会根据这个函数生成实际的依赖注入代码。

```go
// wire.go
//go:build wireinject

func InitializeApp() (*App, error) {
    wire.Build(
        ProviderSet,
        NewApp,
    )
    return nil, nil  // 这行会被 Wire 替换
}
```

**Injector 的规则**：
- 必须在有 `//go:build wireinject` 构建标签的文件中
- 函数体只包含 `wire.Build()` 调用
- 返回值是最终要创建的对象

### 4. 生成的代码

运行 `wire` 命令后，会生成 `wire_gen.go`：

```go
// wire_gen.go (自动生成)
//go:build !wireinject

func InitializeApp() (*App, error) {
    configConfig := config.NewConfig()
    userRepository := repository.NewUserRepository()
    userService := service.NewUserService(userRepository)
    userHandler := handler.NewUserHandler(userService)
    app := NewApp(configConfig, userHandler)
    return app, nil
}
```

**注意**：
- `wire.go` 有 `//go:build wireinject` 标签，只在运行 `wire` 时编译
- `wire_gen.go` 有 `//go:build !wireinject` 标签，在正常编译时使用
- 两个文件互斥，不会同时编译

---

## 代码详解

### config/config.go - 配置层

```go
package config

// Config 应用配置
type Config struct {
    Server ServerConfig
}

type ServerConfig struct {
    Port string
}

// NewConfig 是一个 Provider
// 实际项目中可以从文件、环境变量读取配置
func NewConfig() *Config {
    return &Config{
        Server: ServerConfig{
            Port: ":8080",
        },
    }
}
```

### model/user.go - 数据模型

```go
package model

// User 用户模型
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// CreateUserRequest 创建用户请求
// binding 标签用于 Gin 的参数校验
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=150"`
}
```

### repository/user.go - 数据访问层

```go
package repository

// UserRepository 定义接口，面向接口编程
type UserRepository interface {
    FindAll() ([]*model.User, error)
    FindByID(id int64) (*model.User, error)
    Create(user *model.User) error
    Update(user *model.User) error
    Delete(id int64) error
}

// userRepository 私有实现（小写开头）
type userRepository struct {
    mu     sync.RWMutex
    users  map[int64]*model.User
    nextID int64
}

// NewUserRepository 返回接口类型
// 这样调用者只依赖接口，不依赖具体实现
func NewUserRepository() UserRepository {
    return &userRepository{
        users:  make(map[int64]*model.User),
        nextID: 1,
    }
}
```

### service/user.go - 业务逻辑层

```go
package service

// UserService 业务逻辑接口
type UserService interface {
    GetAllUsers() ([]*model.User, error)
    GetUserByID(id int64) (*model.User, error)
    CreateUser(req *model.CreateUserRequest) (*model.User, error)
    UpdateUser(id int64, req *model.UpdateUserRequest) (*model.User, error)
    DeleteUser(id int64) error
}

type userService struct {
    repo repository.UserRepository  // 依赖 Repository 接口
}

// NewUserService 接收 UserRepository 作为参数
// Wire 会自动注入 NewUserRepository 创建的实例
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}
```

### handler/user.go - HTTP 处理层

```go
package handler

type UserHandler struct {
    svc service.UserService  // 依赖 Service 接口
}

// NewUserHandler 接收 UserService 作为参数
// Wire 会自动注入 NewUserService 创建的实例
func NewUserHandler(svc service.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

// RegisterRoutes 注册路由
func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/users", h.GetUsers)
        api.GET("/users/:id", h.GetUser)
        api.POST("/users", h.CreateUser)
        api.PUT("/users/:id", h.UpdateUser)
        api.DELETE("/users/:id", h.DeleteUser)
    }
}
```

### wire.go - Wire 定义文件

```go
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    // ...
)

// ProviderSet 组合所有 Provider
var ProviderSet = wire.NewSet(
    config.NewConfig,
    repository.NewUserRepository,
    service.NewUserService,
    handler.NewUserHandler,
)

// InitializeApp 是 Injector 函数
func InitializeApp() (*App, error) {
    wire.Build(
        ProviderSet,
        NewApp,
    )
    return nil, nil
}

// App 应用主结构体
type App struct {
    Config  *config.Config
    Engine  *gin.Engine
    Handler *handler.UserHandler
}

// NewApp 也是一个 Provider
func NewApp(cfg *config.Config, h *handler.UserHandler) *App {
    engine := gin.Default()
    h.RegisterRoutes(engine)
    return &App{
        Config:  cfg,
        Engine:  engine,
        Handler: h,
    }
}
```

### main.go - 应用入口

```go
package main

func main() {
    // InitializeApp 由 Wire 生成
    // 会自动按正确顺序创建所有依赖
    app, err := InitializeApp()
    if err != nil {
        log.Fatalf("Failed to initialize app: %v", err)
    }

    log.Printf("Server starting on %s", app.Config.Server.Port)
    app.Engine.Run(app.Config.Server.Port)
}
```

---

## 运行项目

### 安装 Wire

```bash
go install github.com/google/wire/cmd/wire@latest
```

### 生成依赖注入代码

```bash
# 在项目根目录执行
wire
```

### 运行项目

```bash
go run .
```

### 重新生成

修改 `wire.go` 后，需要重新运行 `wire` 命令：

```bash
wire
```

---

## API 测试

### 创建用户

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhang@example.com","age":25}'
```

响应：
```json
{"data":{"id":1,"name":"张三","email":"zhang@example.com","age":25}}
```

### 获取用户列表

```bash
curl http://localhost:8080/api/users
```

响应：
```json
{"data":[{"id":1,"name":"张三","email":"zhang@example.com","age":25}]}
```

### 获取单个用户

```bash
curl http://localhost:8080/api/users/1
```

### 更新用户

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"李四","age":30}'
```

### 删除用户

```bash
curl -X DELETE http://localhost:8080/api/users/1
```

---

## 进阶用法

### 1. 使用 wire.Bind 绑定接口

当 Provider 返回具体类型，但需要绑定到接口时：

```go
var ProviderSet = wire.NewSet(
    NewMySQLUserRepository,  // 返回 *MySQLUserRepository
    wire.Bind(new(UserRepository), new(*MySQLUserRepository)),
)
```

### 2. 使用 wire.Struct 注入结构体字段

```go
type App struct {
    Config  *config.Config
    Handler *handler.UserHandler
}

var ProviderSet = wire.NewSet(
    // 自动注入 App 的所有字段
    wire.Struct(new(App), "*"),
    // 或指定字段
    wire.Struct(new(App), "Config", "Handler"),
)
```

### 3. 使用 wire.Value 注入固定值

```go
var ProviderSet = wire.NewSet(
    wire.Value(Config{Port: ":8080"}),
)
```

### 4. 使用 wire.InterfaceValue 注入接口值

```go
var ProviderSet = wire.NewSet(
    wire.InterfaceValue(new(io.Writer), os.Stdout),
)
```

### 5. 模块化 ProviderSet

```go
// repository/provider.go
var ProviderSet = wire.NewSet(
    NewUserRepository,
    NewOrderRepository,
)

// service/provider.go
var ProviderSet = wire.NewSet(
    NewUserService,
    NewOrderService,
)

// wire.go
var ProviderSet = wire.NewSet(
    repository.ProviderSet,
    service.ProviderSet,
    handler.ProviderSet,
)
```

### 6. 处理 cleanup 函数

```go
// Provider 返回 cleanup 函数
func NewDatabase() (*sql.DB, func(), error) {
    db, err := sql.Open("mysql", "...")
    if err != nil {
        return nil, nil, err
    }
    cleanup := func() {
        db.Close()
    }
    return db, cleanup, nil
}

// Injector 也需要返回 cleanup
func InitializeApp() (*App, func(), error) {
    wire.Build(ProviderSet, NewApp)
    return nil, nil, nil
}

// main.go 中调用 cleanup
func main() {
    app, cleanup, err := InitializeApp()
    if err != nil {
        log.Fatal(err)
    }
    defer cleanup()  // 程序退出时清理资源

    app.Engine.Run(":8080")
}
```

---

## 常见问题

### Q: wire: no provider found for XXX

**原因**：缺少某个类型的 Provider

**解决**：在 ProviderSet 中添加对应的 Provider

### Q: wire: multiple providers for XXX

**原因**：同一类型有多个 Provider

**解决**：移除重复的 Provider，或使用不同的类型

### Q: 修改代码后没有生效

**原因**：没有重新运行 `wire`

**解决**：修改 `wire.go` 后必须重新运行 `wire` 命令

---

## 参考资料

- [Wire GitHub](https://github.com/google/wire)
- [Wire 官方教程](https://github.com/google/wire/blob/main/docs/guide.md)
- [Wire 最佳实践](https://github.com/google/wire/blob/main/docs/best-practices.md)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
