package model

import "time"

// User 用户模型
type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"张三"`
	Email     string    `json:"email" example:"zhangsan@example.com"`
	Phone     string    `json:"phone" example:"13800138000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest 创建用户请求体
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"张三"`
	Email string `json:"email" binding:"required" example:"zhangsan@example.com"`
	Phone string `json:"phone" binding:"required" example:"13800138000"`
}

// UpdateUserRequest 更新用户请求体
type UpdateUserRequest struct {
	Name  string `json:"name" example:"李四"`
	Email string `json:"email" example:"lisi@example.com"`
	Phone string `json:"phone" example:"13900139000"`
}
