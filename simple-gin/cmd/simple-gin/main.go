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
	"example/simple-gin/internal/config"
	"example/simple-gin/internal/container"
	"example/simple-gin/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg := config.LoadConfig()
	log.Printf("Config loaded. Server will run on port %d", cfg.Server.Port)

	// 2. 使用容器进行依赖注入
	c, err := container.NewContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// 3. 设置Gin引擎模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 4. 创建Gin引擎
	r := gin.New()

	// 5. 设置路由
	router.SetupRoutes(r, c.UserService, c.ProductService)

	// 6. 启动服务器
	log.Printf("Starting server on %s", cfg.Server.GetServerAddr())
	log.Printf("Swagger UI: http://localhost%s/swagger/index.html", cfg.Server.GetServerAddr())
	if err := r.Run(cfg.Server.GetServerAddr()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
