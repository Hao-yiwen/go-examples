package config

// Config 应用配置
type Config struct {
	Server ServerConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
}

// NewConfig 创建配置（这是一个 Provider）
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: ":8080",
		},
	}
}
