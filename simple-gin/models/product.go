package models

// Product 产品模型
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

// CreateProductRequest 创建产品请求体
type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Stock    int     `json:"stock" binding:"required,gte=0"`
	Category string  `json:"category" binding:"required"`
}

// UpdateProductRequest 更新产品请求体
type UpdateProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price" binding:"gt=0"`
	Stock    int     `json:"stock" binding:"gte=0"`
	Category string  `json:"category"`
}
