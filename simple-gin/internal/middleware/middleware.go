package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		return "[" + param.TimeStamp.Format("2006-01-02 15:04:05") + "] " +
			statusColor + fmt.Sprint(param.StatusCode) + resetColor + " " +
			methodColor + param.Method + resetColor + " " +
			param.Path + " " +
			param.Latency.String() + "\n"
	})
}

// RecoveryMiddleware 错误恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		log.Printf("panic error: %v", err)
		c.AbortWithStatusJSON(500, gin.H{
			"code": 500,
			"msg":  "internal server error",
		})
	})
}

// CORSMiddleware CORS跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// RequestIDMiddleware 请求ID中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = c.ClientIP() + "-" + time.Now().Format("20060102150405")
		}
		c.Set("request_id", requestID)
		c.Next()
	}
}
