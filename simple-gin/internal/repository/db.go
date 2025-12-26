package repository

import (
	"example/simple-gin/internal/config"
	"example/simple-gin/internal/model"
	"example/simple-gin/internal/service"
	"sync"
	"time"
)

// 编译时验证 DB 实现了 service.Database 接口
var _ service.Database = (*DB)(nil)

// DB 模拟数据库结构
type DB struct {
	users     map[int]*model.User
	products  map[int]*model.Product
	userID    int
	productID int
	mu        sync.RWMutex
}

var db *DB

// Init 初始化数据库连接（模拟）
func Init(cfg *config.Config) (*DB, error) {
	// 这里模拟数据库连接
	// 实际项目中会连接真实数据库（PostgreSQL, MySQL等）
	db = &DB{
		users:     make(map[int]*model.User),
		products:  make(map[int]*model.Product),
		userID:    1,
		productID: 1,
	}

	// 初始化一些模拟数据
	db.seedData()

	return db, nil
}

// seedData 初始化模拟数据
func (d *DB) seedData() {
	now := time.Now()

	// 初始化用户数据
	d.users[1] = &model.User{
		ID:        1,
		Name:      "张三",
		Email:     "zhangsan@example.com",
		Phone:     "13800138000",
		CreatedAt: now,
		UpdatedAt: now,
	}

	d.users[2] = &model.User{
		ID:        2,
		Name:      "李四",
		Email:     "lisi@example.com",
		Phone:     "13800138001",
		CreatedAt: now,
		UpdatedAt: now,
	}

	d.userID = 3

	// 初始化产品数据
	d.products[1] = &model.Product{
		ID:       1,
		Name:     "iPhone 15",
		Price:    5999,
		Stock:    50,
		Category: "Electronics",
	}

	d.products[2] = &model.Product{
		ID:       2,
		Name:     "MacBook Pro",
		Price:    12999,
		Stock:    30,
		Category: "Electronics",
	}

	d.productID = 3
}

// GetDB 获取数据库实例
func GetDB() *DB {
	return db
}

// ======== User Operations ========

// GetUser 获取单个用户
func (d *DB) GetUser(id int) *model.User {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.users[id]
}

// GetAllUsers 获取所有用户
func (d *DB) GetAllUsers() []*model.User {
	d.mu.RLock()
	defer d.mu.RUnlock()

	var users []*model.User
	for _, user := range d.users {
		users = append(users, user)
	}
	return users
}

// CreateUser 创建用户
func (d *DB) CreateUser(req *model.CreateUserRequest) *model.User {
	d.mu.Lock()
	defer d.mu.Unlock()

	user := &model.User{
		ID:        d.userID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	d.users[d.userID] = user
	d.userID++
	return user
}

// UpdateUser 更新用户
func (d *DB) UpdateUser(id int, req *model.UpdateUserRequest) *model.User {
	d.mu.Lock()
	defer d.mu.Unlock()

	user, exists := d.users[id]
	if !exists {
		return nil
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	user.UpdatedAt = time.Now()

	return user
}

// DeleteUser 删除用户
func (d *DB) DeleteUser(id int) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.users[id]; exists {
		delete(d.users, id)
		return true
	}
	return false
}

// ======== Product Operations ========

// GetProduct 获取单个产品
func (d *DB) GetProduct(id int) *model.Product {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.products[id]
}

// GetAllProducts 获取所有产品
func (d *DB) GetAllProducts() []*model.Product {
	d.mu.RLock()
	defer d.mu.RUnlock()

	var products []*model.Product
	for _, product := range d.products {
		products = append(products, product)
	}
	return products
}

// CreateProduct 创建产品
func (d *DB) CreateProduct(req *model.CreateProductRequest) *model.Product {
	d.mu.Lock()
	defer d.mu.Unlock()

	product := &model.Product{
		ID:       d.productID,
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	d.products[d.productID] = product
	d.productID++
	return product
}

// UpdateProduct 更新产品
func (d *DB) UpdateProduct(id int, req *model.UpdateProductRequest) *model.Product {
	d.mu.Lock()
	defer d.mu.Unlock()

	product, exists := d.products[id]
	if !exists {
		return nil
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}

	return product
}

// DeleteProduct 删除产品
func (d *DB) DeleteProduct(id int) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.products[id]; exists {
		delete(d.products, id)
		return true
	}
	return false
}
