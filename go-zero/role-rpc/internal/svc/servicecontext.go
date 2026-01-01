package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/internal/config"
	"github.com/haoyiwen/go-examples/go-zero/role-rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel model.RoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		RoleModel: model.NewRoleModel(conn),
	}
}
