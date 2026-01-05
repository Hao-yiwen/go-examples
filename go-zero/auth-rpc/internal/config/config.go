package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Jwt struct {
		AccessSecret   string
		AccessExpire   int64
		RefreshSecret  string
		RefreshExpire  int64
	}
	AuthRedis struct {
		Host string
		Type string
		Pass string
	}
}
