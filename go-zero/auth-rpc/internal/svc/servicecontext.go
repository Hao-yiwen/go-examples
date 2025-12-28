package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-user/auth-rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
	})

	return &ServiceContext{
		Config: c,
		Redis:  rds,
	}
}
