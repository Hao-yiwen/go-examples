package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/internal/config"
	"github.com/Hao-yiwen/go-examples/go-zero/user-rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
