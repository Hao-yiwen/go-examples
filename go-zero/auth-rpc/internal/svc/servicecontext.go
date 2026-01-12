package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/Hao-yiwen/go-examples/go-zero/auth-rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.AuthRedis.Host,
		Type: c.AuthRedis.Type,
		Pass: c.AuthRedis.Pass,
	})

	return &ServiceContext{
		Config: c,
		Redis:  rds,
	}
}
