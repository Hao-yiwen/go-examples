package model

import "time"

// User 用户模型
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest 创建用户请求体
// 注意：email/phone 格式验证由 service 层使用 pkg/validator 进行
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

// UpdateUserRequest 更新用户请求体
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
