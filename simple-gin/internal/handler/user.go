package handler

import (
	"log/slog"
	"strconv"
	"time"

	"example/simple-gin/internal/model"
	"example/simple-gin/internal/service"
	"example/simple-gin/pkg/response"

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

// GetUsers godoc
//
//	@Summary		获取所有用户
//	@Description	获取用户列表
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]model.User}
//	@Failure		500	{object}	response.Response
//	@Router			/api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		slog.Error("error getting users", "error", err)
		response.InternalError(c, "failed to get users: "+err.Error())
		return
	}

	response.Success(c, users)
}

// GetUser godoc
//
//	@Summary		获取单个用户
//	@Description	根据ID获取用户详情
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"用户ID"
//	@Success		200	{object}	response.Response{data=model.User}
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Router			/api/v1/users/{id} [get]
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
		slog.Error("error getting user", "id", id, "error", err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

// CreateUser godoc
//
//	@Summary		创建用户
//	@Description	创建一个新用户
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.CreateUserRequest	true	"用户信息"
//	@Success		201		{object}	response.Response{data=model.User}
//	@Failure		400		{object}	response.Response
//	@Router			/api/v1/users [post]
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
		slog.Error("error creating user", "error", err)
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, user)
}

// UpdateUser godoc
//
//	@Summary		更新用户
//	@Description	根据ID更新用户信息
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"用户ID"
//	@Param			user	body		model.UpdateUserRequest	true	"更新信息"
//	@Success		200		{object}	response.Response{data=model.User}
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Router			/api/v1/users/{id} [put]
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
		slog.Error("error updating user", "id", id, "error", err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

// DeleteUser godoc
//
//	@Summary		删除用户
//	@Description	根据ID删除用户
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"用户ID"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Router			/api/v1/users/{id} [delete]
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
		slog.Error("error deleting user", "id", id, "error", err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, nil)
}
