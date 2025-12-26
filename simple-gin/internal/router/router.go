package router

import (
	"example/simple-gin/internal/handler"
	"example/simple-gin/internal/middleware"
	"example/simple-gin/internal/service"
	"example/simple-gin/pkg/response"

	_ "example/simple-gin/docs" // swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes 设置所有路由
func SetupRoutes(router *gin.Engine, userService service.UserService, productService service.ProductService) {
	// 应用中间件
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())

	// Swagger 文档路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查
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
			users.GET("", userHandler.GetUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// 产品相关路由
		products := v1.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.POST("", productHandler.CreateProduct)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.POST("/:id/reduce-stock", productHandler.ReduceStock)
		}
	}
}
