package handler

import (
	"example/simple-gin/internal/model"
	"example/simple-gin/internal/service"
	"example/simple-gin/pkg/response"
	"log"
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
	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		log.Printf("Handler: error getting users: %v", err)
		response.InternalError(c, "failed to get users: "+err.Error())
		return
	}

	response.Success(c, users)
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Handler: error getting user %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.CreateUser(ctx, &req)
	if err != nil {
		log.Printf("Handler: error creating user: %v", err)
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, user)
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	user, err := h.userService.UpdateUser(ctx, id, &req)
	if err != nil {
		log.Printf("Handler: error updating user %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.userService.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("Handler: error deleting user %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, nil)
}
