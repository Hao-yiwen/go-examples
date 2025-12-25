package service

import "example/simple-gin/models"

// Database 数据库接口定义
// Service层通过这个接口与Database层交互，实现依赖倒置
type Database interface {
	// User operations
	GetUser(id int) *models.User
	GetAllUsers() []*models.User
	CreateUser(req *models.CreateUserRequest) *models.User
	UpdateUser(id int, req *models.UpdateUserRequest) *models.User
	DeleteUser(id int) bool

	// Product operations
	GetProduct(id int) *models.Product
	GetAllProducts() []*models.Product
	CreateProduct(req *models.CreateProductRequest) *models.Product
	UpdateProduct(id int, req *models.UpdateProductRequest) *models.Product
	DeleteProduct(id int) bool
}
