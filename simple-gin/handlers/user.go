package handlers

import (
	"example/simple-gin/models"
	"example/simple-gin/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers 获取所有用户
func (h *UserHandler) GetUsers(c *gin.Context) {
	// 从request context创建一个具有5秒超时的context
	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		log.Printf("Handler: error getting users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "failed to get users: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": users,
	})
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid user id",
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Handler: error getting user %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid request body: " + err.Error(),
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.CreateUser(ctx, &req)
	if err != nil {
		log.Printf("Handler: error creating user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "user created successfully",
		"data": user,
	})
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid user id",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid request body: " + err.Error(),
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.UpdateUser(ctx, id, &req)
	if err != nil {
		log.Printf("Handler: error updating user %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "user updated successfully",
		"data": user,
	})
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid user id",
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.userService.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("Handler: error deleting user %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "user deleted successfully",
	})
}
