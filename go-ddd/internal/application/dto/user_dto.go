package dto

import (
	"time"

	"yiwen/go-ddd/internal/domain/entity"
)

// DTO (Data Transfer Object) 数据传输对象
// DTO 用于层间数据传递，与领域实体分离
// 好处：
// 1. 隐藏领域模型的内部结构
// 2. 可以根据不同场景返回不同的数据结构
// 3. 便于API版本演进

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Nickname string `json:"nickname" binding:"max=50"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expires_at"`
	User      UserDTO  `json:"user"`
}

// UpdateProfileRequest 更新资料请求
type UpdateProfileRequest struct {
	Nickname string `json:"nickname" binding:"max=50"`
	Avatar   string `json:"avatar" binding:"max=255"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// UserDTO 用户响应DTO
type UserDTO struct {
	ID        uint64    `json:"id"`
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Status    int       `json:"status"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// UserListDTO 用户列表响应DTO
type UserListDTO struct {
	Total int64     `json:"total"`
	Items []UserDTO `json:"items"`
}

// ToUserDTO 将实体转换为DTO
func ToUserDTO(user *entity.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		UUID:      user.UUID,
		Username:  user.Username,
		Email:     user.Email.String(),
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Status:    int(user.Status),
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
	}
}

// ToUserDTOList 将实体列表转换为DTO列表
func ToUserDTOList(users []*entity.User) []UserDTO {
	dtos := make([]UserDTO, len(users))
	for i, user := range users {
		dtos[i] = ToUserDTO(user)
	}
	return dtos
}

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}

// GetOffset 计算偏移量
func (p *PaginationRequest) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return (p.Page - 1) * p.PageSize
}

// GetLimit 获取限制数
func (p *PaginationRequest) GetLimit() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	return p.PageSize
}
