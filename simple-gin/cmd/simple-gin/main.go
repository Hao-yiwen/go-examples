// Package main Simple Gin API
//
//	@title			Simple Gin API
//	@version		1.0
//	@description	一个基于 Gin 框架的示例 API 服务
//
//	@contact.name	API Support
//	@contact.email	support@example.com
//
//	@host			localhost:8080
//	@BasePath		/
//
//	@schemes		http
package main

import (
	"log/slog"
	"os"

	"example/simple-gin/internal/config"
	"example/simple-gin/internal/container"
	"example/simple-gin/internal/router"
	"example/simple-gin/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg := config.LoadConfig()

	// 2. 初始化日志
	logger.Setup(logger.Config{
		Level:  cfg.Logger.Level,
		Format: cfg.Logger.Format,
	})
	slog.Info("config loaded", "port", cfg.Server.Port, "mode", cfg.Server.Mode)

	// 3. 使用容器进行依赖注入
	c, err := container.NewContainer(cfg)
	if err != nil {
		slog.Error("failed to initialize container", "error", err)
		os.Exit(1)
	}

	// 4. 设置Gin引擎模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 5. 创建Gin引擎
	r := gin.New()

	// 6. 设置路由
	router.SetupRoutes(r, c.UserService, c.ProductService)

	// 7. 启动服务器
	slog.Info("starting server",
		"addr", cfg.Server.GetServerAddr(),
		"swagger", "http://localhost"+cfg.Server.GetServerAddr()+"/swagger/index.html",
	)
	if err := r.Run(cfg.Server.GetServerAddr()); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
