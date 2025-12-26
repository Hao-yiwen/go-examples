package model

// Product 产品模型
type Product struct {
	ID       int     `json:"id" example:"1"`
	Name     string  `json:"name" example:"iPhone 15"`
	Price    float64 `json:"price" example:"5999.00"`
	Stock    int     `json:"stock" example:"100"`
	Category string  `json:"category" example:"Electronics"`
}

// CreateProductRequest 创建产品请求体
type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required" example:"iPhone 15"`
	Price    float64 `json:"price" binding:"required,gt=0" example:"5999.00"`
	Stock    int     `json:"stock" binding:"required,gte=0" example:"100"`
	Category string  `json:"category" binding:"required" example:"Electronics"`
}

// UpdateProductRequest 更新产品请求体
type UpdateProductRequest struct {
	Name     string  `json:"name" example:"iPhone 15 Pro"`
	Price    float64 `json:"price" binding:"gt=0" example:"7999.00"`
	Stock    int     `json:"stock" binding:"gte=0" example:"50"`
	Category string  `json:"category" example:"Electronics"`
}
