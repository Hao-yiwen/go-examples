# Simple Gin - 标准 Go 项目结构示例

基于 Gin 框架的 Web 应用，遵循 [golang-standards/project-layout](https://github.com/golang-standards/project-layout) 标准目录结构。

## 项目结构

```
simple-gin/
├── cmd/                         # 主程序入口
│   └── simple-gin/
│       └── main.go              # 应用程序入口
├── configs/                     # 配置文件
│   └── config.yaml              # YAML 配置
├── internal/                    # 私有代码（Go 强制禁止外部导入）
│   ├── config/                  # 配置加载
│   │   └── config.go
│   ├── container/               # 依赖注入容器
│   │   └── container.go
│   ├── handler/                 # HTTP 处理器（控制器）
│   │   ├── helper.go
│   │   ├── product.go
│   │   └── user.go
│   ├── middleware/              # HTTP 中间件
│   │   └── middleware.go
│   ├── model/                   # 数据模型
│   │   ├── product.go
│   │   └── user.go
│   ├── repository/              # 数据访问层
│   │   └── db.go
│   ├── router/                  # 路由配置
│   │   └── router.go
│   └── service/                 # 业务逻辑层
│       ├── database.go          # 接口定义
│       ├── product_service.go
│       └── user_service.go
├── pkg/                         # 公共库（可被外部项目导入）
│   ├── response/                # 统一响应格式
│   │   └── response.go
│   ├── utils/                   # 通用工具
│   │   ├── string.go
│   │   └── string_test.go
│   └── validator/               # 数据验证
│       ├── validator.go
│       └── validator_test.go
├── test/                        # 测试文件
│   ├── integration/             # 集成测试
│   │   └── api_test.go
│   └── testdata/                # 测试数据
│       ├── products.json
│       └── users.json
├── go.mod
├── go.sum
└── README.md
```

## 目录说明

| 目录 | 用途 | 导入限制 |
|------|------|----------|
| `cmd/` | 可执行程序入口，支持多个应用 | - |
| `internal/` | 项目私有代码 | **Go 编译器强制禁止外部导入** |
| `pkg/` | 公共库代码 | 任何项目都可以导入使用 |
| `configs/` | 配置文件模板 | - |
| `test/` | 集成测试和测试数据 | - |

## 快速开始

### 前置条件
- Go 1.21+

### 安装依赖
```bash
go mod download
```

### 运行
```bash
# 方式一：直接运行
go run ./cmd/simple-gin

# 方式二：编译后运行
go build -o bin/simple-gin ./cmd/simple-gin
./bin/simple-gin
```

### 测试
```bash
# 运行所有测试
go test ./...

# 运行单元测试（pkg）
go test -v ./pkg/...

# 运行集成测试
go test -v ./test/integration/...

# 测试覆盖率
go test -cover ./...

# 性能测试
go test -bench=. ./pkg/validator/
```

## API 接口

### 健康检查
```
GET /ping
```

### 用户接口
```
GET    /api/v1/users           # 获取所有用户
POST   /api/v1/users           # 创建用户
GET    /api/v1/users/:id       # 获取指定用户
PUT    /api/v1/users/:id       # 更新用户
DELETE /api/v1/users/:id       # 删除用户
```

### 产品接口
```
GET    /api/v1/products                  # 获取所有产品
POST   /api/v1/products                  # 创建产品
GET    /api/v1/products/:id              # 获取指定产品
PUT    /api/v1/products/:id              # 更新产品
DELETE /api/v1/products/:id              # 删除产品
POST   /api/v1/products/:id/reduce-stock # 减少库存
```

## 使用示例

### API 调用
```bash
# 健康检查
curl http://localhost:8080/ping

# 获取所有用户
curl http://localhost:8080/api/v1/users

# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "王五", "email": "wangwu@example.com", "phone": "13800138002"}'

# 获取单个用户
curl http://localhost:8080/api/v1/users/1

# 更新用户
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "张三（已更新）"}'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### 在其他项目中使用 pkg
```go
import (
    "example/simple-gin/pkg/response"
    "example/simple-gin/pkg/validator"
    "example/simple-gin/pkg/utils"
)

// 统一响应
response.Success(c, data)
response.BadRequest(c, "invalid input")
response.NotFound(c, "user not found")

// 数据验证
validator.IsValidEmail("test@example.com")  // true
validator.IsValidPhone("13800138000")       // true
validator.IsNotEmpty("hello")               // true

// 工具函数
utils.Contains([]int{1, 2, 3}, 2)           // true
utils.Unique([]int{1, 1, 2, 2, 3})          // [1, 2, 3]
utils.Filter(slice, func(x int) bool { return x > 0 })
```

## 架构设计

### 分层架构
```
HTTP Request
    ↓
┌─────────────────┐
│   Middleware    │  ← 日志、CORS、Recovery、RequestID
└────────┬────────┘
         ↓
┌─────────────────┐
│    Handler      │  ← HTTP 请求处理、参数绑定、响应
└────────┬────────┘
         ↓
┌─────────────────┐
│    Service      │  ← 业务逻辑、Context 超时控制
└────────┬────────┘
         ↓
┌─────────────────┐
│   Repository    │  ← 数据访问（当前为内存模拟）
└─────────────────┘
```

### 依赖注入
```
main.go
  ↓
Container（依赖注入容器）
  ├── Config
  ├── Repository (Database)
  ├── Services (UserService, ProductService)
  └── Handlers (UserHandler, ProductHandler)
  ↓
Router（路由注册）
```

### 核心设计原则
- **依赖倒置 (DIP)**: Handler → Service → Repository，通过接口解耦
- **接口隔离**: Database 接口定义在 Service 层
- **控制反转 (IoC)**: Container 管理所有依赖
- **Context 传播**: 请求超时控制贯穿所有层

## 数据模型

### User
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

### Product
```go
type Product struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Stock    int     `json:"stock"`
    Category string  `json:"category"`
}
```

## 中间件

| 中间件 | 功能 |
|--------|------|
| LoggingMiddleware | 请求日志记录 |
| RecoveryMiddleware | Panic 恢复 |
| CORSMiddleware | 跨域资源共享 |
| RequestIDMiddleware | 请求 ID 追踪 |

## 配置

配置文件: `configs/config.yaml`

支持的配置项:
- Server: 端口、模式、超时
- Database: 连接信息
- Logger: 日志级别、格式
- Cache: 缓存类型、TTL
- Middleware: CORS、超时

配置优先级: 环境变量 > 配置文件 > 默认值

环境变量前缀: `SIMPLE_GIN_`

## 扩展指南

### 添加新模型
1. `internal/model/` 创建模型文件
2. `internal/repository/` 添加数据访问方法
3. `internal/service/` 创建服务接口和实现
4. `internal/handler/` 创建处理器
5. `internal/router/` 注册路由
6. `internal/container/` 添加依赖注入

### 添加公共库
1. `pkg/` 下创建包目录
2. 编写代码和单元测试
3. 其他项目可直接导入使用

### 连接真实数据库
1. 修改 `internal/repository/db.go`
2. 使用 GORM 或其他 ORM
3. 更新 `configs/config.yaml`

## 初始数据

项目启动时自动初始化:

**用户:**
- 张三 (zhangsan@example.com)
- 李四 (lisi@example.com)

**产品:**
- iPhone 15 (5999 元)
- MacBook Pro (12999 元)

## 依赖

- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Web 框架
- [spf13/viper](https://github.com/spf13/viper) - 配置管理

---

这是一个标准的 Go 项目模板，可直接用于新项目开发。
