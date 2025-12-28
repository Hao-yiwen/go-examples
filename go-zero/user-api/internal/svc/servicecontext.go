package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-user/auth-rpc/auth"
	"go-zero-user/role-rpc/role"
	"go-zero-user/user-api/internal/config"
	"go-zero-user/user-rpc/user"
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
