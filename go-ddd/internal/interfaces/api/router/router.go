package router

import (
	"github.com/gin-gonic/gin"

	"yiwen/go-ddd/internal/interfaces/api/handler"
	"yiwen/go-ddd/internal/interfaces/api/middleware"
)

// Router 路由管理
type Router struct {
	engine      *gin.Engine
	userHandler *handler.UserHandler
	jwtAuth     *middleware.JWTAuth
}

// NewRouter 创建路由
func NewRouter(userHandler *handler.UserHandler, jwtAuth *middleware.JWTAuth) *Router {
	return &Router{
		engine:      gin.New(),
		userHandler: userHandler,
		jwtAuth:     jwtAuth,
	}
}

// Setup 设置路由
func (r *Router) Setup() *gin.Engine {
	// 全局中间件
	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())
	r.engine.Use(CORSMiddleware())

	// 健康检查
	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1
	v1 := r.engine.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			// 公开接口（无需认证）
			users.POST("/register", r.userHandler.Register)
			users.POST("/login", r.userHandler.Login)

			// 需要认证的接口
			authUsers := users.Group("")
			authUsers.Use(r.jwtAuth.AuthMiddleware())
			{
				authUsers.GET("/me", r.userHandler.GetCurrentUser)
				authUsers.GET("/:id", r.userHandler.GetUser)
				authUsers.PUT("/:id", r.userHandler.UpdateProfile)
				authUsers.POST("/:id/password", r.userHandler.ChangePassword)
			}

			// 管理员接口
			adminUsers := users.Group("")
			adminUsers.Use(r.jwtAuth.AuthMiddleware())
			adminUsers.Use(r.jwtAuth.AdminMiddleware())
			{
				adminUsers.GET("", r.userHandler.ListUsers)
				adminUsers.DELETE("/:id", r.userHandler.DeleteUser)
			}
		}
	}

	return r.engine
}

// CORSMiddleware CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
