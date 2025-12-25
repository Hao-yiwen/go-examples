# 配置文件使用指南

本项目使用 **Viper** 进行配置管理，这是Go企业级项目的标准做法。

## 📖 快速开始

### 1. 配置文件位置

项目会按以下顺序查找配置文件：
```
1. 当前目录           ./config.yaml
2. config 子目录      ./config/config.yaml
3. HOME 目录          ~/config.yaml
```

### 2. 配置文件格式

使用 **YAML** 格式（可扩展支持JSON、TOML等）

```yaml
server:
  port: 8080          # 服务器端口
  mode: debug         # debug 或 release
  timeout: 30s        # 请求超时

database:
  driver: postgres    # 数据库驱动
  host: localhost
  port: 5432
  user: postgres
  password: password123
  name: simple_gin_db
```

## 🔧 三种配置来源（优先级）

Viper 支持多种配置来源，优先级从高到低：

### 1️⃣ **环境变量** (最高优先级)
```bash
# 前缀为 SIMPLE_GIN_，使用下划线分隔嵌套字段
export SIMPLE_GIN_SERVER_PORT=9090
export SIMPLE_GIN_DATABASE_HOST=db.example.com
export SIMPLE_GIN_DATABASE_PORT=5432

./simple-gin
# 会使用环境变量中的配置，覆盖配置文件
```

### 2️⃣ **配置文件** (中等优先级)
```yaml
# config.yaml
server:
  port: 8080
  mode: debug
```

### 3️⃣ **默认值** (最低优先级)
如果配置文件和环境变量都未指定，使用代码中的默认值。

## 📝 完整配置示例

```yaml
# Simple-Gin 完整配置文件示例

# ========== 服务器配置 ==========
server:
  port: 8080          # 监听端口 (1-65535)
  mode: debug         # 运行模式: debug 或 release
  timeout: 30s        # 请求超时时间

# ========== 数据库配置 ==========
database:
  driver: postgres    # 数据库驱动
  host: localhost     # 主机地址
  port: 5432         # 端口号
  user: postgres     # 用户名
  password: password123  # 密码
  name: simple_gin_db    # 数据库名
  max_connections: 10    # 最大连接数
  idle_connections: 5    # 空闲连接数
  max_idle_time: 300     # 最大空闲时间（秒）

# ========== 日志配置 ==========
logger:
  level: info        # 日志级别: debug, info, warn, error
  format: json       # 格式: json 或 text
  output: stdout     # 输出: stdout 或 file
  file_path: ./logs/app.log  # 日志文件路径

# ========== 缓存配置 ==========
cache:
  type: memory       # 缓存类型: memory 或 redis
  ttl: 3600         # TTL (秒)

# ========== 中间件配置 ==========
middleware:
  cors:
    allowed_origins:
      - http://localhost:3000
      - http://localhost:8080
    allowed_methods:
      - GET
      - POST
      - PUT
      - DELETE
    allowed_headers:
      - Content-Type
      - Authorization

  request_timeout: 30s  # 请求超时
```

## 🌍 环境变量映射关系

| 环境变量 | 配置路径 | 说明 |
|---------|---------|------|
| `SIMPLE_GIN_SERVER_PORT` | `server.port` | 服务器端口 |
| `SIMPLE_GIN_SERVER_MODE` | `server.mode` | 运行模式 |
| `SIMPLE_GIN_DATABASE_HOST` | `database.host` | 数据库主机 |
| `SIMPLE_GIN_DATABASE_PORT` | `database.port` | 数据库端口 |
| `SIMPLE_GIN_DATABASE_USER` | `database.user` | 数据库用户 |
| `SIMPLE_GIN_DATABASE_PASSWORD` | `database.password` | 数据库密码 |
| `SIMPLE_GIN_DATABASE_NAME` | `database.name` | 数据库名 |
| `SIMPLE_GIN_LOGGER_LEVEL` | `logger.level` | 日志级别 |

## 💡 使用场景

### 场景1：本地开发
```bash
# 使用默认值或config.yaml即可
./simple-gin
```

### 场景2：测试环境
```bash
# 使用配置文件，但覆盖关键参数
export SIMPLE_GIN_DATABASE_HOST=test-db.local
export SIMPLE_GIN_DATABASE_PASSWORD=test_password
export SIMPLE_GIN_LOGGER_LEVEL=debug

./simple-gin
```

### 场景3：生产环境
```bash
# 通过环境变量完全控制配置（推荐）
export SIMPLE_GIN_SERVER_MODE=release
export SIMPLE_GIN_SERVER_PORT=80
export SIMPLE_GIN_DATABASE_HOST=prod-db.example.com
export SIMPLE_GIN_DATABASE_PASSWORD=$DB_PASSWORD  # 从密钥管理系统获取
export SIMPLE_GIN_LOGGER_LEVEL=warn

./simple-gin
```

## 🔐 安全最佳实践

### ❌ 不要做这些
```yaml
# ❌ 敏感信息不要写在配置文件中
database:
  password: secret123  # 明文密码很危险！
```

### ✅ 应该这样做
```yaml
# ✅ 使用环境变量注入敏感信息
database:
  password: ""  # 空值，由环境变量提供

# 或在部署时通过环境变量提供
# export SIMPLE_GIN_DATABASE_PASSWORD=$DB_PASSWORD
```

### 生产环境推荐
- 使用 **Kubernetes Secrets** 或 **Docker Secrets**
- 使用 **HashiCorp Vault** 等密钥管理系统
- 使用 **AWS Parameter Store** 或 **AWS Secrets Manager**
- 使用 **环境变量**（配合 CI/CD 密钥管理）

## 📖 配置加载流程

```
┌─────────────────────────────────┐
│  main.go: config.LoadConfig()   │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  创建 Viper 新实例              │ (避免全局状态污染)
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  1. 读取 config.yaml 文件       │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  2. 设置默认值（填充缺失项）    │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  3. 应用环境变量覆盖 (AutomaticEnv)  │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  4. Unmarshal 到 Config 结构体   │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  5. 验证配置有效性              │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────┐
│  返回 *Config，程序继续运行     │
└─────────────────────────────────┘
```

## 🛠️ 代码示例

### 加载配置
```go
package main

import (
    "example/simple-gin/config"
)

func main() {
    // 加载配置（自动查找config.yaml）
    cfg := config.LoadConfig()

    // 或从指定文件加载
    // cfg := config.LoadConfigFromFile("./config/prod.yaml")

    // 打印配置信息（不显示密码）
    cfg.Print()

    // 使用配置
    port := cfg.Server.Port
    dbHost := cfg.DB.Host
}
```

### 在容器中使用
```go
func (c *Container) init(cfg *config.Config) {
    // CORS 配置
    corsConfig := cfg.Middleware.CORS

    // 数据库连接
    dsn := cfg.DB.GetDSN()

    // 日志配置
    logLevel := cfg.Logger.Level
}
```

### 添加新配置项

1. 在 `config.yaml` 中添加
```yaml
redis:
  host: localhost
  port: 6379
  password: ""
```

2. 在 `config/config.go` 中添加结构体
```go
type Config struct {
    // ...
    Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
}
```

3. 在 `setDefaultsWithViper` 中添加默认值
```go
v.SetDefault("redis.host", "localhost")
v.SetDefault("redis.port", 6379)
v.SetDefault("redis.password", "")
```

## 🔍 调试配置

### 查看所有加载的配置
```go
cfg := config.LoadConfig()

// 打印所有配置
viper.AllKeys()  // 显示所有key

// 查看特定配置
viper.Get("database.host")
viper.GetInt("server.port")
viper.GetString("logger.level")
```

### 查看配置文件路径
```go
cfg := config.LoadConfig()
// viper 会打印: ✅ Config file loaded: /path/to/config.yaml
```

## 📚 相关文档

- [Viper 官方文档](https://github.com/spf13/viper)
- [YAML 格式规范](https://yaml.org/)
- [12-Factor App 配置原则](https://12factor.net/config)

## 总结

| 功能 | 状态 |
|------|------|
| 配置文件支持 | ✅ (YAML) |
| 环境变量覆盖 | ✅ |
| 默认值支持 | ✅ |
| 配置验证 | ✅ |
| 配置热重载 | ⚠️  (可扩展) |
| 多环境支持 | ✅ |

---

**推荐做法**：
- ✅ 开发环境：使用 `config.yaml`
- ✅ 测试环境：使用 `config.yaml` + 环境变量覆盖
- ✅ 生产环境：完全使用环境变量（无配置文件）
