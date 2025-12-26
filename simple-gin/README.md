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
├── docs/                        # Swagger 文档（自动生成）
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/                    # 私有代码（Go 强制禁止外部导入）
│   ├── config/                  # 配置加载
│   ├── container/               # 依赖注入容器
│   ├── handler/                 # HTTP 处理器（含 Swagger 注释）
│   ├── middleware/              # HTTP 中间件
│   ├── model/                   # 数据模型
│   ├── repository/              # 数据访问层
│   ├── router/                  # 路由配置
│   └── service/                 # 业务逻辑层
├── pkg/                         # 公共库（可被外部项目导入）
│   ├── response/                # 统一响应格式
│   ├── utils/                   # 通用工具
│   └── validator/               # 数据验证
├── test/                        # 测试文件
│   ├── integration/             # 集成测试
│   └── testdata/                # 测试数据
├── go.mod
├── go.sum
└── README.md
```

## 目录说明

| 目录 | 用途 | 导入限制 |
|------|------|----------|
| `cmd/` | 可执行程序入口 | - |
| `internal/` | 项目私有代码 | **Go 编译器强制禁止外部导入** |
| `pkg/` | 公共库代码 | 任何项目都可以导入使用 |
| `docs/` | Swagger 文档 | 自动生成，勿手动修改 |
| `configs/` | 配置文件 | - |
| `test/` | 集成测试和测试数据 | - |

## 快速开始

### 前置条件
- Go 1.21+
- Make（可选，用于简化命令）
- swag CLI（可选，用于更新文档）

### 安装依赖
```bash
# 使用 Make
make deps

# 或手动安装
go mod download

# 安装开发工具（swag, lint, air）
make tools
```

### 运行
```bash
# 使用 Make
make run          # 直接运行
make dev          # 开发模式（生成文档 + 运行）
make build        # 编译到 bin/

# 或手动运行
go run ./cmd/simple-gin
go build -o bin/simple-gin ./cmd/simple-gin
./bin/simple-gin
```

启动后访问：
- API: http://localhost:8080
- Swagger UI: http://localhost:8080/swagger/index.html

### 测试
```bash
# 使用 Make
make test              # 运行所有测试
make test-v            # 详细输出
make test-cover        # 显示覆盖率
make test-cover-html   # 生成 HTML 覆盖率报告
make test-unit         # 只运行单元测试
make test-integration  # 只运行集成测试
make bench             # 运行性能测试

# 或手动运行
go test ./...
go test -v ./pkg/...
go test -v ./test/integration/...
go test -cover ./...
```

## Swagger 文档

### 访问方式
启动服务后访问：http://localhost:8080/swagger/index.html

### 更新文档
修改 handler 中的注释后，重新生成文档：
```bash
# 使用 Make
make docs

# 或手动运行
swag init -g cmd/simple-gin/main.go -o docs
```

### 注释格式
```go
// CreateUser godoc
// @Summary      创建用户
// @Description  创建一个新用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.CreateUserRequest  true  "用户信息"
// @Success      201   {object}  response.Response{data=model.User}
// @Failure      400   {object}  response.Response
// @Router       /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
```

### 常用注释标签

| 标签 | 说明 | 示例 |
|------|------|------|
| `@Summary` | 接口简述 | `@Summary 创建用户` |
| `@Description` | 详细描述 | `@Description 创建一个新用户` |
| `@Tags` | 分组标签 | `@Tags users` |
| `@Param` | 参数定义 | `@Param id path int true "用户ID"` |
| `@Success` | 成功响应 | `@Success 200 {object} Response` |
| `@Failure` | 失败响应 | `@Failure 400 {object} Response` |
| `@Router` | 路由路径 | `@Router /api/v1/users [post]` |

## API 接口

### 健康检查
```
GET /ping
```

### Swagger 文档
```
GET /swagger/*any
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
│    Handler      │  ← HTTP 请求处理、参数绑定、Swagger 注释
└────────┬────────┘
         ↓
┌─────────────────┐
│    Service      │  ← 业务逻辑、数据验证、Context 超时
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
Router（路由注册 + Swagger）
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

### 添加新接口
1. `internal/model/` 创建模型文件（添加 `example` 标签）
2. `internal/repository/` 添加数据访问方法
3. `internal/service/` 创建服务接口和实现
4. `internal/handler/` 创建处理器（添加 Swagger 注释）
5. `internal/router/` 注册路由
6. `internal/container/` 添加依赖注入
7. 运行 `swag init -g cmd/simple-gin/main.go -o docs` 更新文档

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
- [swaggo/swag](https://github.com/swaggo/swag) - Swagger 文档生成
- [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) - Gin Swagger 中间件

## Makefile 命令

项目提供了完整的 Makefile 简化开发流程。查看所有可用命令：

```bash
make help
```

### 命令速查表

| 分类 | 命令 | 说明 |
|------|------|------|
| **构建** | `make build` | 编译项目到 `bin/` |
| | `make run` | 直接运行项目 |
| | `make dev` | 开发模式（生成文档 + 运行） |
| **测试** | `make test` | 运行所有测试 |
| | `make test-v` | 运行测试（详细输出） |
| | `make test-cover` | 显示测试覆盖率 |
| | `make test-cover-html` | 生成覆盖率 HTML 报告 |
| | `make test-unit` | 只运行单元测试（pkg） |
| | `make test-integration` | 只运行集成测试 |
| | `make bench` | 运行性能测试 |
| **文档** | `make docs` | 生成 Swagger 文档 |
| | `make docs-fmt` | 格式化 Swagger 注释 |
| **依赖** | `make deps` | 下载依赖 |
| | `make deps-update` | 更新依赖 |
| | `make deps-tidy` | 清理依赖 |
| **代码质量** | `make fmt` | 格式化代码 |
| | `make lint` | 代码检查（需要 golangci-lint） |
| | `make vet` | go vet 检查 |
| **清理** | `make clean` | 清理构建产物 |
| | `make clean-all` | 清理所有（包括文档） |
| **工具** | `make tools` | 安装开发工具（swag, lint, air） |

### 常用开发流程

```bash
# 首次使用
make deps           # 下载依赖
make tools          # 安装开发工具

# 日常开发
make dev            # 生成文档并运行

# 提交前检查
make fmt            # 格式化代码
make lint           # 代码检查
make test           # 运行测试

# 构建发布
make build          # 编译项目
```

---

这是一个标准的 Go 项目模板，可直接用于新项目开发。
