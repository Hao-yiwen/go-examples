package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config 应用程序全局配置结构
type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	DB         DatabaseConfig   `mapstructure:"database"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	Cache      CacheConfig      `mapstructure:"cache"`
	Middleware MiddlewareConfig `mapstructure:"middleware"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Mode    string `mapstructure:"mode"` // debug, release
	Timeout string `mapstructure:"timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `mapstructure:"driver"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxConnections  int    `mapstructure:"max_connections"`
	IdleConnections int    `mapstructure:"idle_connections"`
	MaxIdleTime     int    `mapstructure:"max_idle_time"` // 秒数
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level    string `mapstructure:"level"`  // debug, info, warn, error
	Format   string `mapstructure:"format"` // json, text
	Output   string `mapstructure:"output"` // stdout, file
	FilePath string `mapstructure:"file_path"`
}

// CacheConfig 缓存配置
type CacheConfig struct {
	Type string `mapstructure:"type"` // memory, redis
	TTL  int    `mapstructure:"ttl"`
}

// MiddlewareConfig 中间件配置
type MiddlewareConfig struct {
	CORS           CORSConfig `mapstructure:"cors"`
	RequestTimeout string     `mapstructure:"request_timeout"`
}

// CORSConfig CORS 配置
type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

// LoadConfig 从配置文件加载配置（使用 Viper）
// 优先级：环境变量 > 配置文件 > 默认值
func LoadConfig() *Config {
	// 创建新的Viper实例，避免全局状态污染
	v := viper.New()

	// 配置文件设置
	v.SetConfigName("config")          // 配置文件名（不包括扩展名）
	v.SetConfigType("yaml")            // 配置文件格式
	v.AddConfigPath(".")               // 在当前目录查找
	v.AddConfigPath("./configs")       // 在 configs 目录查找
	v.AddConfigPath(os.Getenv("HOME")) // 在 HOME 目录查找

	// 环境变量支持
	v.SetEnvPrefix("SIMPLE_GIN")
	v.AutomaticEnv()

	// 读取配置文件（先读，后设置默认值）
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults and environment variables")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	} else {
		log.Printf("Config file loaded: %s", v.ConfigFileUsed())
	}

	// 设置所有默认值（覆盖缺失的配置项）
	setDefaultsWithViper(v)

	// 反序列化为结构体
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil
	}

	// 验证配置
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Config validation failed: %v", err)
	}

	log.Println("Config loaded successfully")
	return cfg
}

// setDefaultsWithViper 设置Viper实例的配置默认值
func setDefaultsWithViper(v *viper.Viper) {
	// Server
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.mode", "debug")
	v.SetDefault("server.timeout", "30s")

	// Database
	v.SetDefault("database.driver", "postgres")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "password123")
	v.SetDefault("database.name", "simple_gin_db")
	v.SetDefault("database.max_connections", 10)
	v.SetDefault("database.idle_connections", 5)
	v.SetDefault("database.max_idle_time", 300)

	// Logger
	v.SetDefault("logger.level", "info")
	v.SetDefault("logger.format", "json")
	v.SetDefault("logger.output", "stdout")
	v.SetDefault("logger.file_path", "./logs/app.log")

	// Cache
	v.SetDefault("cache.type", "memory")
	v.SetDefault("cache.ttl", 3600)

	// Middleware
	v.SetDefault("middleware.request_timeout", "30s")
	v.SetDefault("middleware.cors.allowed_origins", []string{"*"})
	v.SetDefault("middleware.cors.allowed_methods", []string{"GET", "POST", "PUT", "DELETE"})
	v.SetDefault("middleware.cors.allowed_headers", []string{"*"})
}

// Validate 验证配置的合法性
func (c *Config) Validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Server.Mode != "debug" && c.Server.Mode != "release" {
		return fmt.Errorf("invalid server mode: %s (must be 'debug' or 'release')", c.Server.Mode)
	}

	if c.DB.Port < 1 || c.DB.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", c.DB.Port)
	}

	if c.DB.Host == "" {
		return fmt.Errorf("database host is required")
	}

	return nil
}

// GetServerAddr 获取服务器地址
func (c *ServerConfig) GetServerAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}
