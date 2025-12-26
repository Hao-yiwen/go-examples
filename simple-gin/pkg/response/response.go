// Package response 提供统一的 HTTP 响应格式
// 这个包可以被其他项目导入使用
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Created 创建成功响应
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Code:    0,
		Message: "created successfully",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, httpCode int, code int, message string) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
	})
}

// BadRequest 400 错误
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, 400, message)
}

// NotFound 404 错误
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, 404, message)
}

// InternalError 500 错误
func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, 500, message)
}

// Unauthorized 401 错误
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, 401, message)
}

// Forbidden 403 错误
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, 403, message)
}
