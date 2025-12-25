package main

import (
	"example/simple-gin/config"
	"example/simple-gin/container"
	"example/simple-gin/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg := config.LoadConfig()
	log.Printf("Config loaded. Server will run on port %d", cfg.Server.Port)

	// 2. 使用容器进行依赖注入
	// 容器负责初始化所有依赖关系：数据库、服务、处理器等
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
	router := gin.New()

	// 5. 设置路由（从容器中获取服务）
	// 现在路由层只需关心服务接口，不需要关心初始化细节
	routes.SetupRoutes(router, c.UserService, c.ProductService)

	// 6. 启动服务器
	log.Printf("Starting server on %s", cfg.Server.GetServerAddr())
	if err := router.Run(cfg.Server.GetServerAddr()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// 7. 优雅关闭（可选）
	// if err := c.Close(); err != nil {
	//     log.Fatalf("Failed to close container: %v", err)
	// }
}
