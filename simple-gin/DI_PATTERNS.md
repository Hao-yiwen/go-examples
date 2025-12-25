# Go 中的依赖注入模式

这个项目演示了 Go 中企业级项目推荐的依赖注入模式，从简单到高级。

## 📚 背景

**依赖注入 (DI)** 是 SOLID 原则中的重要部分，核心思想：
- ❌ 不要在组件内部创建依赖
- ✅ 让依赖从外部注入进来

这样做的好处：
1. **解耦** - 各层完全独立，可独立测试
2. **灵活** - 易于替换实现（如切换数据库）
3. **可测性** - 容易 Mock 依赖进行单元测试
4. **维护** - 修改一个实现不影响其他模块

---

## 方案 1️⃣：构造注入（Simple DI）

**适用场景**：小到中型项目，依赖不复杂

### 实现方式

```go
// 1. 定义接口
type UserService interface {
    GetUsers(ctx context.Context) ([]*models.User, error)
}

// 2. 实现接口
type userService struct {
    db Database  // 依赖其他接口
}

// 3. 构造函数注入依赖
func NewUserService(db Database) UserService {
    return &userService{db: db}
}

// 4. main.go 中初始化
func main() {
    db := database.Init(cfg)
    userService := service.NewUserService(db)
    handler := handlers.NewUserHandler(userService)
}
```

### 优点
- ✅ 简单直接，易于理解
- ✅ 无额外工具或框架
- ✅ 编译时检查，类型安全
- ✅ 适合快速开发

### 缺点
- ❌ 依赖多时 main 函数会很长
- ❌ 依赖关系散落在代码各处
- ❌ 难以管理生命周期

### 代码示例（旧版本）
```go
// main.go 版本 1
func main() {
    cfg := config.LoadConfig()
    db, _ := database.Init(cfg)
    userService := service.NewUserService(db)
    productService := service.NewProductService(db)

    userHandler := handlers.NewUserHandler(userService)
    productHandler := handlers.NewProductHandler(productService)

    router := gin.New()
    routes.SetupRoutes(router, userService, productService)

    // 随着项目增长，这个函数会越来越长...
}
```

---

## 方案 2️⃣：容器模式（推荐 ⭐）

**适用场景**：大多数生产项目

### 实现方式

```go
// container/container.go
type Container struct {
    Config *config.Config
    DB     service.Database

    UserService    service.UserService
    ProductService service.ProductService

    UserHandler    *handlers.UserHandler
    ProductHandler *handlers.ProductHandler
}

func NewContainer(cfg *config.Config) (*Container, error) {
    c := &Container{Config: cfg}

    // 分步骤初始化，清晰明了
    if err := c.initDatabase(); err != nil {
        return nil, err
    }
    c.initServices()
    c.initHandlers()

    return c, nil
}

// main.go 变得简洁
func main() {
    cfg := config.LoadConfig()
    container, _ := container.NewContainer(cfg)

    router := gin.New()
    routes.SetupRoutes(router, container.UserService, container.ProductService)
    router.Run(cfg.Server.GetServerAddr())
}
```

### 优点
- ✅ 所有依赖在一个地方管理
- ✅ main 函数保持简洁
- ✅ 易于扩展（添加新服务只需修改容器）
- ✅ 支持生命周期管理
- ✅ 支持优雅关闭
- ✅ 可以集中处理配置逻辑

### 缺点
- ⚠️ 容器层增加了一个抽象
- ⚠️ 需要谨慎避免循环依赖

### 适用场景示例
```go
// 添加日志系统
type Container struct {
    Logger logger.Logger
    // ...
}

func (c *Container) initLogger() error {
    c.Logger = logger.NewLogger(c.Config)
    return nil
}

// 添加缓存系统
type Container struct {
    Cache cache.Cache
    // ...
}

// 添加数据库连接池
type Container struct {
    DB   *sql.DB
    // ...
}

// 处理优雅关闭
func (c *Container) Close() error {
    if err := c.Logger.Close(); err != nil {
        return err
    }
    if err := c.Cache.Close(); err != nil {
        return err
    }
    // ...
    return nil
}
```

---

## 方案 3️⃣：使用 Google Wire（企业级）

**适用场景**：超大型项目，需要代码生成和完全的自动化

### 概念

Wire 是 Google 提供的依赖注入工具，使用代码生成方式：

```go
// wire.go
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*App, error) {
    wire.Build(
        database.NewDB,
        service.NewUserService,
        service.NewProductService,
        handlers.NewUserHandler,
        handlers.NewProductHandler,
        routes.NewRouter,
        wire.Struct(new(App), "*"),
    )
    return nil, nil
}

// main.go
func main() {
    cfg := config.LoadConfig()
    app, err := InitializeApp(cfg)
    if err != nil {
        log.Fatal(err)
    }
    app.Start()
}
```

### 优点
- ✅ 完全自动化，代码生成，零运行时开销
- ✅ 编译时检查所有依赖
- ✅ 支持复杂的依赖图
- ✅ 生成的代码可以查看和调试

### 缺点
- ❌ 学习曲线陡峭
- ❌ 需要额外的构建步骤 (`wire` 命令)
- ❌ 小项目可能过度设计
- ❌ 生成的代码可能难以理解

### 何时使用
- 项目有 20+ 个服务和组件
- 需要复杂的条件依赖
- 团队成员对 Wire 有经验

---

## 🔍 本项目的方案演进

### 版本 1：直接在 main 中初始化（最初）
```go
func main() {
    db, _ := database.Init(cfg)
    userService := service.NewUserService(db)
    handler := handlers.NewUserHandler(userService)
}
```

### 版本 2：构造注入（当前推荐用于理解）
```go
// 所有初始化逻辑明确可见
// 适合学习和理解依赖关系
```

### 版本 3：容器模式（已实现 ⭐ 推荐）
```go
container, _ := container.NewContainer(cfg)
routes.SetupRoutes(router, container.UserService, container.ProductService)
```

---

## 📊 选择指南

| 方案 | 项目规模 | 复杂度 | 推荐度 | 何时升级 |
|------|--------|------|------|--------|
| **构造注入** | 小 (<10 服务) | 低 | ⭐⭐⭐ | 项目初期 |
| **容器模式** | 中-大 (10-50 服务) | 中 | ⭐⭐⭐⭐⭐ | 有 3+ 个服务时 |
| **Wire** | 超大 (50+ 服务) | 高 | ⭐⭐⭐⭐ | 需要代码生成时 |

---

## 💡 最佳实践

### 1. 优先使用接口
```go
// ✅ Good: 依赖接口，不依赖实现
type UserHandler struct {
    userService service.UserService  // 接口，不是具体实现
}

// ❌ Bad: 直接依赖实现
type UserHandler struct {
    userService *userServiceImpl  // 紧耦合
}
```

### 2. 在顶层进行注入
```go
// ✅ Good: 在 main 或容器中进行注入
func main() {
    container := NewContainer(cfg)
    // 只在最顶层知道具体实现
}

// ❌ Bad: 在中间层进行注入
func (h *Handler) doSomething() {
    service := NewService()  // 不应该在这里创建！
}
```

### 3. 使用工厂函数
```go
// ✅ Good: 使用 New 前缀的工厂函数
func NewUserService(db Database) UserService {
    return &userService{db: db}
}

// ❌ Bad: 在内部创建依赖
func NewUserService() UserService {
    db := connectDatabase()  // 这就是服务定位器模式，不是 DI
}
```

### 4. 管理生命周期
```go
// ✅ Good: 提供关闭方法
func (c *Container) Close() error {
    c.Logger.Close()
    c.DB.Close()
    c.Cache.Close()
    return nil
}

// 在 main 中使用
defer container.Close()
```

---

## 🎯 本项目推荐

**我们建议使用：容器模式 (方案 2)**

### 原因
1. **平衡** - 既不过度设计，也不过于简单
2. **可维护** - 清晰集中的依赖管理
3. **可扩展** - 轻松添加新的服务和中间件
4. **实用** - 大多数生产项目都在用
5. **学习价值** - 理解原理后，升级到 Wire 很容易

### 使用方式（当前项目）
```go
// container/container.go
type Container struct {
    DB             service.Database
    UserService    service.UserService
    ProductService service.ProductService
    UserHandler    *handlers.UserHandler
    ProductHandler *handlers.ProductHandler
}

// main.go
func main() {
    cfg := config.LoadConfig()
    c, _ := container.NewContainer(cfg)

    router := gin.New()
    routes.SetupRoutes(router, c.UserService, c.ProductService)
    router.Run(cfg.Server.GetServerAddr())
}
```

---

## 📖 进阶阅读

- [SOLID 原则](https://en.wikipedia.org/wiki/SOLID)
- [Dependency Injection in Go](https://pkg.go.dev/github.com/google/wire)
- [Service Locator Pattern vs DI](https://martinfowler.com/articles/injection.html)

---

## 总结

| 特性 | 值 |
|------|-----|
| **推荐模式** | 容器模式 |
| **当前项目使用** | ✅ 已实现 |
| **编译状态** | ✅ 通过 |
| **测试状态** | ✅ 已验证 |

**下一步**：根据项目成长，可以考虑升级到 Wire 工具。
