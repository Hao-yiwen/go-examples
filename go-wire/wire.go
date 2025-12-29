//go:build wireinject
// +build wireinject

// wire.go 定义了依赖注入的规则
// 这个文件只在生成代码时使用，不会被编译到最终程序中

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"yiwen/go-wire/config"
	"yiwen/go-wire/handler"
	"yiwen/go-wire/repository"
	"yiwen/go-wire/service"
)

// ProviderSet 是所有 Provider 的集合
// wire.NewSet 用于将多个 Provider 组合在一起
var ProviderSet = wire.NewSet(
	config.NewConfig,
	repository.ProviderSet,
	service.ProviderSet,
	handler.ProviderSet,
)

// InitializeApp 是一个 injector 函数
// Wire 会根据这个函数生成实际的依赖注入代码
// 函数体中的 wire.Build() 告诉 Wire 使用哪些 Provider
func InitializeApp() (*App, error) {
	wire.Build(
		ProviderSet,
		NewApp,
	)
	// 下面这行代码会被 Wire 生成的代码替换
	// 返回值类型需要匹配函数签名
	return nil, nil
}

// App 是应用程序的主结构体
type App struct {
	Config         *config.Config
	Engine         *gin.Engine
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
}

// NewApp 创建应用程序（这也是一个 Provider）
func NewApp(cfg *config.Config, uh *handler.UserHandler, ph *handler.ProductHandler) *App {
	engine := gin.Default()
	uh.RegisterRoutes(engine)
	ph.RegisterRoutes(engine)

	return &App{
		Config:         cfg,
		Engine:         engine,
		UserHandler:    uh,
		ProductHandler: ph,
	}
}
