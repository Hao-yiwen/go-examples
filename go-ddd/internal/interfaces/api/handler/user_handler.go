package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"yiwen/go-ddd/internal/application/command"
	"yiwen/go-ddd/internal/application/dto"
	"yiwen/go-ddd/internal/application/query"
	"yiwen/go-ddd/internal/application/service"
	"yiwen/go-ddd/internal/interfaces/api/middleware"
)

// UserHandler 用户HTTP处理器
// 接口层负责：
// 1. 接收HTTP请求
// 2. 参数校验和转换
// 3. 调用应用服务
// 4. 返回HTTP响应
type UserHandler struct {
	userService *service.UserApplicationService
	jwtAuth     *middleware.JWTAuth
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService *service.UserApplicationService, jwtAuth *middleware.JWTAuth) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtAuth:     jwtAuth,
	}
}

// Register 用户注册
// POST /api/v1/users/register
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	cmd := command.NewRegisterUserCommand(req.Username, req.Email, req.Password, req.Nickname)
	user, err := h.userService.Register(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}

// Login 用户登录
// POST /api/v1/users/login
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	q := query.NewLoginQuery(req.Username, req.Password)
	user, err := h.userService.Login(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	// 生成JWT Token
	token, expiresAt, err := h.jwtAuth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": dto.LoginResponse{
			Token:     token,
			ExpiresAt: expiresAt,
			User:      *user,
		},
	})
}

// GetUser 获取用户信息
// GET /api/v1/users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid user id",
		})
		return
	}

	q := query.NewGetUserByIDQuery(id)
	user, err := h.userService.GetUserByID(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}

// ListUsers 获取用户列表
// GET /api/v1/users
func (h *UserHandler) ListUsers(c *gin.Context) {
	var req dto.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	q := query.NewListUsersQuery(req.GetOffset(), req.GetLimit())
	result, err := h.userService.ListUsers(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

// UpdateProfile 更新用户资料
// PUT /api/v1/users/:id
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid user id",
		})
		return
	}

	// 验证是否是本人操作
	currentUserID, _ := middleware.GetUserIDFromContext(c)
	if currentUserID != id {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "permission denied",
		})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	cmd := command.NewUpdateProfileCommand(id, req.Nickname, req.Avatar)
	user, err := h.userService.UpdateProfile(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}

// ChangePassword 修改密码
// POST /api/v1/users/:id/password
func (h *UserHandler) ChangePassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid user id",
		})
		return
	}

	// 验证是否是本人操作
	currentUserID, _ := middleware.GetUserIDFromContext(c)
	if currentUserID != id {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "permission denied",
		})
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	cmd := command.NewChangePasswordCommand(id, req.OldPassword, req.NewPassword)
	if err := h.userService.ChangePassword(c.Request.Context(), cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "password changed successfully",
	})
}

// DeleteUser 删除用户
// DELETE /api/v1/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "invalid user id",
		})
		return
	}

	cmd := command.NewDeleteUserCommand(id)
	if err := h.userService.DeleteUser(c.Request.Context(), cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "user deleted successfully",
	})
}

// GetCurrentUser 获取当前登录用户信息
// GET /api/v1/users/me
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}

	q := query.NewGetUserByIDQuery(userID)
	user, err := h.userService.GetUserByID(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}
