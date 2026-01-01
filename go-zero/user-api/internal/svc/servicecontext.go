package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/haoyiwen/go-examples/go-zero/auth-rpc/auth"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/role"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/config"
	"github.com/haoyiwen/go-examples/go-zero/user-rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	AuthRpc  auth.Auth
	RoleRpc  role.Role
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthRpc:  auth.NewAuth(zrpc.MustNewClient(c.AuthRpc)),
		RoleRpc:  role.NewRole(zrpc.MustNewClient(c.RoleRpc)),
	}
}
