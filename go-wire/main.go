package main

import (
	"log"
)

func main() {
	// InitializeApp 由 Wire 生成，会自动完成依赖注入
	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	log.Printf("Server starting on %s", app.Config.Server.Port)

	// 启动 HTTP 服务
	if err := app.Engine.Run(app.Config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
