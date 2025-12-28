# Go DDD 用户管理系统

基于领域驱动设计（Domain-Driven Design）架构的 Go 语言用户管理系统示例项目。

## 目录

- [项目简介](#项目简介)
- [技术栈](#技术栈)
- [项目结构](#项目结构)
- [DDD 分层架构详解](#ddd-分层架构详解)
- [快速开始](#快速开始)
- [API 接口](#api-接口)
- [核心概念示例](#核心概念示例)
- [依赖注入流程](#依赖注入流程)

---

## 项目简介

这是一个用于学习 DDD（领域驱动设计）的示例项目，实现了一个完整的用户管理系统，包括：

- 用户注册、登录
- JWT 认证授权
- 用户信息管理
- 角色权限控制

通过这个项目，你可以学习到：

1. DDD 四层架构的划分和职责
2. 实体（Entity）与值对象（Value Object）的区别
3. 聚合根（Aggregate Root）的设计
4. 仓储模式（Repository Pattern）的应用
5. 领域事件（Domain Event）的使用
6. CQRS 模式（命令查询职责分离）

---

## 技术栈

| 技术 | 说明 |
|-----|-----|
| Go 1.21+ | 编程语言 |
| Gin | Web 框架 |
| GORM | ORM 框架 |
| MySQL | 数据库 |
| JWT | 认证方案 |
| Viper | 配置管理 |
| bcrypt | 密码加密 |

---

## 项目结构

```
go-ddd/
├── cmd/
│   └── api/
│       └── main.go                 # 程序入口，依赖注入
├── config/
│   └── config.yaml                 # 配置文件
├── internal/
│   ├── domain/                     # 【领域层】核心业务逻辑
│   │   ├── entity/                 # 实体
│   │   │   └── user.go
│   │   ├── valueobject/            # 值对象
│   │   │   ├── email.go
│   │   │   └── password.go
│   │   ├── aggregate/              # 聚合
│   │   │   └── user_aggregate.go
│   │   ├── repository/             # 仓储接口
│   │   │   └── user_repository.go
│   │   ├── service/                # 领域服务
│   │   │   └── user_domain_service.go
│   │   └── event/                  # 领域事件
│   │       └── user_events.go
│   ├── application/                # 【应用层】用例编排
│   │   ├── dto/                    # 数据传输对象
│   │   │   └── user_dto.go
│   │   ├── command/                # 命令（写操作）
│   │   │   └── user_command.go
│   │   ├── query/                  # 查询（读操作）
│   │   │   └── user_query.go
│   │   └── service/                # 应用服务
│   │       └── user_service.go
│   ├── infrastructure/             # 【基础设施层】技术实现
│   │   ├── config/                 # 配置管理
│   │   │   └── config.go
│   │   └── persistence/            # 持久化
│   │       ├── model/              # 数据库模型
│   │       │   └── user_model.go
│   │       └── mysql/              # MySQL 实现
│   │           └── user_repository.go
│   └── interfaces/                 # 【接口层】对外暴露
│       └── api/
│           ├── handler/            # HTTP 处理器
│           │   └── user_handler.go
│           ├── middleware/         # 中间件
│           │   └── auth.go
│           └── router/             # 路由
│               └── router.go
├── pkg/                            # 公共包
│   └── errors/
│       └── errors.go
├── scripts/
│   └── sql/
│       └── init.sql                # 建表 SQL
├── go.mod
├── go.sum
└── README.md
```

---

## DDD 分层架构详解

### 1. 领域层（Domain Layer）

**位置**: `internal/domain/`

领域层是 DDD 的核心，包含所有业务逻辑，不依赖任何外部框架。

#### 1.1 实体（Entity）

```go
// internal/domain/entity/user.go

// 实体的特点：
// 1. 有唯一标识（ID）
// 2. 生命周期内标识不变
// 3. 可变的（属性可以改变）
// 4. 相等性由 ID 决定

type User struct {
    ID        uint64              // 唯一标识
    UUID      string              // 业务标识
    Username  string
    Email     valueobject.Email   // 值对象
    Password  valueobject.Password
    Status    UserStatus
    Role      UserRole
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 实体的行为方法
func (u *User) IsActive() bool {
    return u.Status == UserStatusActive
}

func (u *User) UpdateProfile(nickname, avatar string) {
    u.Nickname = nickname
    u.Avatar = avatar
    u.UpdatedAt = time.Now()
}
```

#### 1.2 值对象（Value Object）

```go
// internal/domain/valueobject/email.go

// 值对象的特点：
// 1. 没有唯一标识
// 2. 不可变（immutable）
// 3. 相等性由所有属性决定
// 4. 可以包含验证逻辑

type Email struct {
    value string  // 私有字段，确保不可变
}

// 工厂方法，包含验证逻辑
func NewEmail(email string) (Email, error) {
    if !emailRegex.MatchString(email) {
        return Email{}, ErrInvalidEmail
    }
    return Email{value: email}, nil
}

// 只读方法
func (e Email) String() string {
    return e.value
}

// 比较方法
func (e Email) Equals(other Email) bool {
    return e.value == other.value
}
```

#### 1.3 聚合（Aggregate）

```go
// internal/domain/aggregate/user_aggregate.go

// 聚合的特点：
// 1. 一组相关对象的集合
// 2. 有一个聚合根作为入口
// 3. 外部只能通过聚合根访问内部对象
// 4. 聚合根负责维护内部一致性
// 5. 可以发布领域事件

type UserAggregate struct {
    User   *entity.User    // 聚合根实体
    Events []event.Event   // 待发布的领域事件
}

// 通过聚合执行业务操作
func (a *UserAggregate) ChangePassword(newPassword valueobject.Password) {
    a.User.ChangePassword(newPassword)
    // 发布领域事件
    a.addEvent(event.NewUserPasswordChangedEvent(a.User.UUID))
}
```

#### 1.4 仓储接口（Repository Interface）

```go
// internal/domain/repository/user_repository.go

// 仓储模式的特点：
// 1. 领域层只定义接口
// 2. 基础设施层提供实现
// 3. 实现依赖倒置原则
// 4. 便于单元测试（可 mock）

type UserRepository interface {
    Save(ctx context.Context, user *entity.User) error
    FindByID(ctx context.Context, id uint64) (*entity.User, error)
    FindByUsername(ctx context.Context, username string) (*entity.User, error)
    Delete(ctx context.Context, id uint64) error
    List(ctx context.Context, offset, limit int) ([]*entity.User, int64, error)
}
```

#### 1.5 领域服务（Domain Service）

```go
// internal/domain/service/user_domain_service.go

// 领域服务的特点：
// 1. 处理不属于单个实体的业务逻辑
// 2. 涉及多个实体的操作
// 3. 需要访问仓储的验证逻辑

type UserDomainService struct {
    userRepo repository.UserRepository
}

// 验证用户名唯一性（需要查询数据库）
func (s *UserDomainService) ValidateUniqueUsername(ctx context.Context, username string) error {
    exists, err := s.userRepo.ExistsByUsername(ctx, username)
    if exists {
        return ErrUsernameAlreadyExists
    }
    return nil
}

// 验证登录凭证
func (s *UserDomainService) ValidateUserCredentials(ctx context.Context, username, password string) (*entity.User, error) {
    user, err := s.userRepo.FindByUsername(ctx, username)
    if err != nil {
        return nil, ErrInvalidCredentials
    }
    if err := user.Password.Verify(password); err != nil {
        return nil, ErrInvalidCredentials
    }
    return user, nil
}
```

#### 1.6 领域事件（Domain Event）

```go
// internal/domain/event/user_events.go

// 领域事件的特点：
// 1. 表示领域中发生的有意义的事情
// 2. 事件是不可变的
// 3. 用于解耦不同的领域/服务
// 4. 支持事件溯源

type UserRegisteredEvent struct {
    BaseEvent
    Username string `json:"username"`
    Email    string `json:"email"`
}

func NewUserRegisteredEvent(uuid, username, email string) *UserRegisteredEvent {
    return &UserRegisteredEvent{
        BaseEvent: BaseEvent{
            Name:        "user.registered",
            OccurredOn:  time.Now(),
            AggregateId: uuid,
        },
        Username: username,
        Email:    email,
    }
}
```

---

### 2. 应用层（Application Layer）

**位置**: `internal/application/`

应用层负责编排领域对象完成用例，不包含业务逻辑。

#### 2.1 DTO（Data Transfer Object）

```go
// internal/application/dto/user_dto.go

// DTO 用于层间数据传递，与领域实体分离

type RegisterRequest struct {
    Username string `json:"username" binding:"required,min=3,max=50"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type UserDTO struct {
    ID       uint64 `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    // 注意：不暴露密码
}

// 实体转 DTO
func ToUserDTO(user *entity.User) UserDTO {
    return UserDTO{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email.String(),
    }
}
```

#### 2.2 Command/Query（CQRS 模式）

```go
// Command - 写操作
type RegisterUserCommand struct {
    Username string
    Email    string
    Password string
}

// Query - 读操作
type GetUserByIDQuery struct {
    UserID uint64
}
```

#### 2.3 应用服务

```go
// internal/application/service/user_service.go

type UserApplicationService struct {
    userRepo          repository.UserRepository
    userDomainService *domainservice.UserDomainService
}

func (s *UserApplicationService) Register(ctx context.Context, cmd *command.RegisterUserCommand) (*dto.UserDTO, error) {
    // 1. 调用领域服务验证
    if err := s.userDomainService.ValidateUniqueUsername(ctx, cmd.Username); err != nil {
        return nil, err
    }

    // 2. 创建值对象
    email, err := valueobject.NewEmail(cmd.Email)
    password, err := valueobject.NewPassword(cmd.Password)

    // 3. 使用聚合创建用户
    userAggregate := aggregate.Register(uuid.New().String(), cmd.Username, email, password)

    // 4. 调用仓储保存
    if err := s.userRepo.Save(ctx, userAggregate.User); err != nil {
        return nil, err
    }

    // 5. 返回 DTO
    return dto.ToUserDTO(userAggregate.User), nil
}
```

---

### 3. 基础设施层（Infrastructure Layer）

**位置**: `internal/infrastructure/`

基础设施层提供技术实现，如数据库访问、外部服务调用等。

#### 3.1 仓储实现

```go
// internal/infrastructure/persistence/mysql/user_repository.go

// 实现领域层定义的仓储接口

type UserRepository struct {
    db *gorm.DB
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
    // 实体转数据库模型
    userModel := model.FromEntity(user)

    if user.ID == 0 {
        return r.db.WithContext(ctx).Create(userModel).Error
    }
    return r.db.WithContext(ctx).Save(userModel).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
    var userModel model.UserModel
    if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
        return nil, err
    }
    // 数据库模型转实体
    return userModel.ToEntity(), nil
}
```

#### 3.2 数据库模型

```go
// internal/infrastructure/persistence/model/user_model.go

// 数据库模型与领域实体分离

type UserModel struct {
    ID           uint64         `gorm:"primaryKey"`
    UUID         string         `gorm:"uniqueIndex"`
    Username     string         `gorm:"uniqueIndex"`
    Email        string         `gorm:"uniqueIndex"`
    PasswordHash string         // 存储哈希值
    DeletedAt    gorm.DeletedAt // 软删除
}

// 模型转实体
func (m *UserModel) ToEntity() *entity.User {
    email, _ := valueobject.NewEmail(m.Email)
    password := valueobject.NewPasswordFromHash(m.PasswordHash)
    return &entity.User{
        ID:       m.ID,
        Email:    email,
        Password: password,
    }
}

// 实体转模型
func FromEntity(user *entity.User) *UserModel {
    return &UserModel{
        ID:           user.ID,
        Email:        user.Email.String(),
        PasswordHash: user.Password.Hash(),
    }
}
```

---

### 4. 接口层（Interfaces Layer）

**位置**: `internal/interfaces/`

接口层负责与外部交互，如 HTTP、gRPC、消息队列等。

```go
// internal/interfaces/api/handler/user_handler.go

type UserHandler struct {
    userService *service.UserApplicationService
    jwtAuth     *middleware.JWTAuth
}

func (h *UserHandler) Register(c *gin.Context) {
    // 1. 绑定请求参数
    var req dto.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // 2. 构造命令
    cmd := command.NewRegisterUserCommand(req.Username, req.Email, req.Password)

    // 3. 调用应用服务
    user, err := h.userService.Register(c.Request.Context(), cmd)

    // 4. 返回响应
    c.JSON(201, gin.H{"data": user})
}
```

---

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 5.7+

### 2. 初始化数据库

```bash
# 登录 MySQL
mysql -u root -p

# 执行初始化脚本
source scripts/sql/init.sql
```

### 3. 修改配置

编辑 `config/config.yaml`：

```yaml
app:
  name: go-ddd
  port: 8080
  mode: debug

database:
  host: localhost
  port: 3306
  username: root
  password: your_password  # 修改为你的密码
  database: go_ddd

jwt:
  secret: your-secret-key  # 生产环境请修改
  expire_hour: 24
```

### 4. 运行项目

```bash
# 下载依赖
go mod tidy

# 运行
go run cmd/api/main.go

# 或指定配置文件
go run cmd/api/main.go -config config/config.yaml
```

### 5. 验证运行

```bash
# 健康检查
curl http://localhost:8080/health
# 返回: {"status":"ok"}
```

---

## API 接口

### 公开接口

#### 用户注册

```bash
POST /api/v1/users/register

# 请求
{
    "username": "testuser",
    "email": "test@example.com",
    "password": "Test1234"
}

# 响应
{
    "code": 0,
    "message": "success",
    "data": {
        "id": 1,
        "uuid": "xxx-xxx-xxx",
        "username": "testuser",
        "email": "test@example.com"
    }
}
```

#### 用户登录

```bash
POST /api/v1/users/login

# 请求
{
    "username": "testuser",
    "password": "Test1234"
}

# 响应
{
    "code": 0,
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIs...",
        "expires_at": 1703836800,
        "user": {
            "id": 1,
            "username": "testuser"
        }
    }
}
```

### 需要认证的接口

请求头需要添加：`Authorization: Bearer <token>`

#### 获取当前用户

```bash
GET /api/v1/users/me
```

#### 获取用户详情

```bash
GET /api/v1/users/:id
```

#### 更新用户资料

```bash
PUT /api/v1/users/:id

{
    "nickname": "新昵称",
    "avatar": "https://example.com/avatar.jpg"
}
```

#### 修改密码

```bash
POST /api/v1/users/:id/password

{
    "old_password": "Test1234",
    "new_password": "NewTest1234"
}
```

### 管理员接口

需要 admin 角色

#### 用户列表

```bash
GET /api/v1/users?page=1&page_size=10
```

#### 删除用户

```bash
DELETE /api/v1/users/:id
```

---

## 依赖注入流程

```
main.go
    │
    ├── 1. 加载配置
    │       config.Load()
    │
    ├── 2. 初始化数据库
    │       gorm.Open()
    │
    ├── 3. 初始化仓储层（基础设施层）
    │       userRepo := mysql.NewUserRepository(db)
    │
    ├── 4. 初始化领域服务（领域层）
    │       userDomainService := domainservice.NewUserDomainService(userRepo)
    │
    ├── 5. 初始化应用服务（应用层）
    │       userAppService := appservice.NewUserApplicationService(userRepo, userDomainService)
    │
    ├── 6. 初始化 HTTP 处理器（接口层）
    │       userHandler := handler.NewUserHandler(userAppService, jwtAuth)
    │
    └── 7. 启动服务
            router.Setup().Run()
```

---

## DDD 核心原则总结

| 原则 | 说明 |
|-----|-----|
| **领域优先** | 领域层是核心，不依赖任何外部框架 |
| **依赖倒置** | 高层不依赖低层，都依赖抽象（接口） |
| **充血模型** | 实体包含业务逻辑，而不是贫血的数据结构 |
| **聚合边界** | 通过聚合根维护业务一致性 |
| **值对象不可变** | 值对象创建后不可修改 |
| **仓储隔离** | 领域层只定义接口，基础设施层实现 |

---

## 测试账户

初始化脚本会创建一个管理员账户：

- 用户名：`admin`
- 密码：`Admin123`

> 注意：生产环境请修改或删除此账户

---

## License

MIT License
