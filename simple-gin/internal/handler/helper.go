package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// createContextWithTimeout 从Gin context创建具有超时的context
// 这样可以在请求层面控制业务逻辑的超时时间
func createContextWithTimeout(c *gin.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	// 获取Gin的request context
	ctx := c.Request.Context()

	// 创建具有超时的context，如果Gin context更早被取消也会传播
	return context.WithTimeout(ctx, timeout)
}
