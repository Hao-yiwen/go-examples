package router

import (
	"example/simple-gin/internal/handler"
	"example/simple-gin/internal/middleware"
	"example/simple-gin/internal/service"
	"example/simple-gin/pkg/response"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(router *gin.Engine, userService service.UserService, productService service.ProductService) {
	// 应用中间件
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())

	// 健康检查 - 使用 pkg/response
	router.GET("/ping", func(c *gin.Context) {
		response.Success(c, gin.H{"message": "pong"})
	})

	// 创建处理器实例
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)          // 获取所有用户
			users.POST("", userHandler.CreateUser)       // 创建用户
			users.GET("/:id", userHandler.GetUser)       // 获取单个用户
			users.PUT("/:id", userHandler.UpdateUser)    // 更新用户
			users.DELETE("/:id", userHandler.DeleteUser) // 删除用户
		}

		// 产品相关路由
		products := v1.Group("/products")
		{
			products.GET("", productHandler.GetProducts)                   // 获取所有产品
			products.POST("", productHandler.CreateProduct)                // 创建产品
			products.GET("/:id", productHandler.GetProduct)                // 获取单个产品
			products.PUT("/:id", productHandler.UpdateProduct)             // 更新产品
			products.DELETE("/:id", productHandler.DeleteProduct)          // 删除产品
			products.POST("/:id/reduce-stock", productHandler.ReduceStock) // 减少库存
		}
	}
}
