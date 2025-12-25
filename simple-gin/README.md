# Simple Gin - 规范的Gin项目结构示例

这是一个基于Gin框架的规范Web应用项目，展示了企业级Go项目的标准架构和最佳实践。

## 项目结构

```
simple-gin/
├── config/           # 配置模块
│   └── config.go     # 应用程序配置（服务器、数据库等）
├── models/           # 数据模型层
│   ├── user.go       # 用户模型 + 请求体定义
│   └── product.go    # 产品模型 + 请求体定义
├── database/         # 数据层（模拟数据存储）
│   └── db.go         # 模拟数据库实现，实现 Database 接口
├── service/          # 业务逻辑层（新增）
│   ├── database.go   # Database 接口定义（依赖倒置）
│   ├── user_service.go   # 用户服务实现（使用context）
│   └── product_service.go # 产品服务实现（使用context）
├── handlers/         # HTTP请求处理层（控制器）
│   ├── helper.go     # Context 辅助函数
│   ├── user.go       # 用户处理器（依赖UserService）
│   └── product.go    # 产品处理器（依赖ProductService）
├── middleware/       # 中间件层
│   └── middleware.go # 日志、错误恢复、CORS等中间件
├── routes/           # 路由配置层
│   └── routes.go     # 定义所有API路由、依赖注入
├── main.go           # 应用程序入口
├── go.mod            # Go Module 依赖声明
└── go.sum            # Go Module 依赖哈希值

```

## 架构设计特点

### 企业级分层架构
- **Models 层**: 定义数据结构和请求/响应体
- **Database 层**: 数据存储和访问逻辑（这里模拟）
  - 实现 `service.Database` 接口
  - 编译时验证接口实现
- **Service 层** (新增): 业务逻辑层
  - `UserService` 和 `ProductService` 接口定义
  - 使用 Go `context.Context` 处理超时和取消
  - 包含业务规则校验
  - 依赖 Database 接口（依赖倒置原则）
- **Handlers 层** (控制器): HTTP 请求处理
  - `UserHandler` 和 `ProductHandler` 控制器
  - 依赖注入 Service 接口
  - 创建具有超时的 context 传递给 Service
- **Routes 层**: 路由定义和依赖注入
- **Middleware 层**: 横切关注点（日志、错误处理、CORS等）
- **Config 层**: 应用程序配置管理

### 核心设计原则
- ✅ **依赖倒置原则 (DIP)**: Handler → Service → Database，通过接口解耦
- ✅ **接口隔离**: Database 接口定义在 Service 层，实现在 Database 层
- ✅ **控制反转 (IoC)**: 在 main.go 中进行依赖注入
- ✅ **Context 传播**: 从 HTTP 请求到业务逻辑的超时控制
- ✅ **RESTful API** 设计
- ✅ 请求参数验证（使用 Gin 的 binding tag）
- ✅ 统一的 JSON 响应格式
- ✅ 中间件支持（日志、错误恢复、CORS、请求ID）
- ✅ 线程安全的数据存储（使用 sync.RWMutex）
- ✅ 完整的 CRUD 操作

## API 接口

### 用户接口
```
GET    /api/v1/users           # 获取所有用户
POST   /api/v1/users           # 创建新用户
GET    /api/v1/users/:id       # 获取指定用户
PUT    /api/v1/users/:id       # 更新指定用户
DELETE /api/v1/users/:id       # 删除指定用户
```

### 产品接口
```
GET    /api/v1/products        # 获取所有产品
POST   /api/v1/products        # 创建新产品
GET    /api/v1/products/:id    # 获取指定产品
PUT    /api/v1/products/:id    # 更新指定产品
DELETE /api/v1/products/:id    # 删除指定产品
POST   /api/v1/products/:id/reduce-stock  # 减少产品库存（业务操作示例）
```

### 健康检查
```
GET    /ping                   # 健康检查
```

## 运行项目

### 前置条件
- Go 1.25 或更高版本
- 安装依赖: `go mod download`

### 启动服务器
```bash
go run main.go
```

输出示例：
```
2024-12-26 10:15:30 Config loaded. Server will run on port 8080
2024-12-26 10:15:30 Database initialized. Connection: postgres://postgres:password123@localhost:5432/simple_gin_db
2024-12-26 10:15:30 Starting server on :8080
```

### 编译构建
```bash
go build
./simple-gin
```

## API 使用示例

### 获取所有用户
```bash
curl http://localhost:8080/api/v1/users
```

### 创建新用户
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "王五",
    "email": "wangwu@example.com",
    "phone": "13800138002"
  }'
```

### 获取指定用户
```bash
curl http://localhost:8080/api/v1/users/1
```

### 更新用户
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三（更新）",
    "email": "zhangsan_new@example.com"
  }'
```

### 删除用户
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 数据模型

### User (用户)
```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Product (产品)
```go
type Product struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Stock    int     `json:"stock"`
    Category string  `json:"category"`
}
```

## 配置说明

所有配置在 `config/config.go` 中定义：

```go
Config struct {
    Server ServerConfig        // 服务器配置 (端口、模式)
    DB     DatabaseConfig      // 数据库配置 (模拟)
}
```

**注意**: 这个项目的数据库配置是模拟的，使用内存存储。实际生产项目中应该连接真实的数据库（PostgreSQL、MySQL 等）。

## 中间件列表

1. **LoggingMiddleware** - 请求/响应日志记录
2. **RecoveryMiddleware** - Panic 恢复处理
3. **CORSMiddleware** - 跨域资源共享
4. **RequestIDMiddleware** - 请求ID追踪

## 项目初始化数据

项目启动时会自动初始化以下数据：

**用户**:
- ID: 1, Name: 张三, Email: zhangsan@example.com
- ID: 2, Name: 李四, Email: lisi@example.com

**产品**:
- ID: 1, Name: iPhone 15, Price: 5999, Stock: 50
- ID: 2, Name: MacBook Pro, Price: 12999, Stock: 30

## 从这个模板扩展

### 添加新的模型
1. 在 `models/` 下创建新文件（如 `order.go`）
2. 在 `database/` 中添加 CRUD 方法
3. 在 `handlers/` 下创建处理器文件
4. 在 `routes/routes.go` 中注册路由

### 连接真实数据库
1. 更新 `config/config.go` 的数据库配置
2. 使用 ORM（如 GORM）替换 `database/db.go` 的内存存储
3. 更新处理器以使用真实数据库查询

### 添加身份验证
1. 在 `middleware/` 中创建认证中间件
2. 在需要保护的路由上应用该中间件

## 最佳实践

✅ **遵循的原则**:
- 清晰的分层架构
- 单一职责原则
- 一致的错误处理
- 统一的 API 响应格式
- 请求验证和类型安全
- 线程安全的数据操作
- RESTful API 设计

## Service 层使用 Context 的示例

Service 层的每个方法都接受 `context.Context` 参数，用于控制超时和取消：

```go
// UserService 接口定义
type UserService interface {
    GetUsers(ctx context.Context) ([]*models.User, error)
    GetUserByID(ctx context.Context, id int) (*models.User, error)
    CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
    UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error)
    DeleteUser(ctx context.Context, id int) error
}

// Handler 中的使用
func (h *UserHandler) GetUsers(c *gin.Context) {
    // 创建 5 秒超时的 context
    ctx, cancel := createContextWithTimeout(c, 5*time.Second)
    defer cancel()

    // Service 会监听 context 的取消信号
    users, err := h.userService.GetUsers(ctx)
    // ...
}
```

**Context 的作用**:
- ✅ 请求级别的超时控制（5秒）
- ✅ 监听客户端连接中断（context 取消）
- ✅ 跨层级传播超时信息
- ✅ 支持链路追踪的请求 ID 传递

## 依赖注入流程

```
main.go
  ↓
创建 Database 实例
  ↓
创建 UserService 和 ProductService（注入 Database）
  ↓
routes.SetupRoutes(router, userService, productService)
  ↓
创建 UserHandler 和 ProductHandler（注入 Service）
  ↓
注册路由处理函数
```

这种设计使得：
- 各层完全解耦，可独立测试
- 容易替换实现（如将模拟 Database 替换为真实数据库）
- 遵循 SOLID 原则

## 生成的依赖

主要依赖:
- `github.com/gin-gonic/gin` - Gin Web 框架

所有完整依赖见 `go.mod` 和 `go.sum` 文件。

---

## 关键文件说明

### service/user_service.go
- 实现 `UserService` 接口
- 所有方法都使用 `context.Context`
- 包含业务规则校验（如 email 格式、ID 合法性）
- 监听 context 取消信号

### handlers/user.go
- 实现 `UserHandler` 控制器
- 依赖注入 `UserService` 接口
- 创建超时 context 并传递给 Service
- 错误映射为 HTTP 状态码

### handlers/helper.go
- 定义 `createContextWithTimeout` 函数
- 将 Gin 的 `Request.Context()` 转换为带超时的 context

### routes/routes.go
- 配置所有 API 路由
- 进行依赖注入（创建 Handler 实例）
- 路由分组和中间件应用

---

这个项目展示了一个规范、可扩展的**企业级** Go Web 应用架构，涵盖了：
- 完整的分层设计
- 接口设计与依赖注入
- Context 的正确使用
- 错误处理与日志
- RESTful API 设计

可以直接用作新项目的基础模板！
