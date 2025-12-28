package model

// User 用户模型
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=0,lte=150"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
	Age   int    `json:"age" binding:"gte=0,lte=150"`
}
