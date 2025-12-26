package service

import "example/simple-gin/internal/model"

// Database 数据库接口定义
// Service层通过这个接口与Database层交互，实现依赖倒置
type Database interface {
	// User operations
	GetUser(id int) *model.User
	GetAllUsers() []*model.User
	CreateUser(req *model.CreateUserRequest) *model.User
	UpdateUser(id int, req *model.UpdateUserRequest) *model.User
	DeleteUser(id int) bool

	// Product operations
	GetProduct(id int) *model.Product
	GetAllProducts() []*model.Product
	CreateProduct(req *model.CreateProductRequest) *model.Product
	UpdateProduct(id int, req *model.UpdateProductRequest) *model.Product
	DeleteProduct(id int) bool
}
