package service

import (
	"context"
	"errors"
	"example/simple-gin/internal/model"
	"example/simple-gin/pkg/validator"
	"log"
)

// ProductService 产品服务接口定义
type ProductService interface {
	// GetProducts 获取所有产品
	GetProducts(ctx context.Context) ([]*model.Product, error)
	// GetProductByID 根据ID获取产品
	GetProductByID(ctx context.Context, id int) (*model.Product, error)
	// CreateProduct 创建产品
	CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.Product, error)
	// UpdateProduct 更新产品
	UpdateProduct(ctx context.Context, id int, req *model.UpdateProductRequest) (*model.Product, error)
	// DeleteProduct 删除产品
	DeleteProduct(ctx context.Context, id int) error
	// ReduceStock 减少产品库存
	ReduceStock(ctx context.Context, id, quantity int) error
}

// productService 产品服务实现
type productService struct {
	db Database
}

// NewProductService 创建产品服务实例
func NewProductService(db Database) ProductService {
	return &productService{
		db: db,
	}
}

// GetProducts 实现获取所有产品
func (s *productService) GetProducts(ctx context.Context) ([]*model.Product, error) {
	select {
	case <-ctx.Done():
		log.Printf("GetProducts request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	log.Println("Service: fetching all products")
	products := s.db.GetAllProducts()

	if products == nil {
		products = make([]*model.Product, 0)
	}

	return products, nil
}

// GetProductByID 实现根据ID获取产品
func (s *productService) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	select {
	case <-ctx.Done():
		log.Printf("GetProductByID request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if id <= 0 {
		return nil, errors.New("invalid product id")
	}

	log.Printf("Service: fetching product by id: %d", id)
	product := s.db.GetProduct(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// CreateProduct 实现创建产品
func (s *productService) CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.Product, error) {
	select {
	case <-ctx.Done():
		log.Printf("CreateProduct request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if req == nil {
		return nil, errors.New("invalid request")
	}

	// 使用 pkg/validator 进行验证
	if !validator.IsNotEmpty(req.Name) {
		return nil, errors.New("product name is required")
	}

	if !validator.IsPositive(req.Price) {
		return nil, errors.New("price must be greater than 0")
	}

	if !validator.IsNotEmpty(req.Category) {
		return nil, errors.New("category is required")
	}

	if !validator.IsNonNegative(float64(req.Stock)) {
		return nil, errors.New("stock cannot be negative")
	}

	log.Printf("Service: creating product with name: %s", req.Name)
	product := s.db.CreateProduct(req)

	return product, nil
}

// UpdateProduct 实现更新产品
func (s *productService) UpdateProduct(ctx context.Context, id int, req *model.UpdateProductRequest) (*model.Product, error) {
	select {
	case <-ctx.Done():
		log.Printf("UpdateProduct request cancelled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
	}

	if id <= 0 {
		return nil, errors.New("invalid product id")
	}

	// 检查产品是否存在
	existingProduct := s.db.GetProduct(id)
	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	// 使用 pkg/validator 验证更新字段
	if req.Price != 0 && !validator.IsPositive(req.Price) {
		return nil, errors.New("price must be greater than 0")
	}

	log.Printf("Service: updating product with id: %d", id)
	product := s.db.UpdateProduct(id, req)

	return product, nil
}

// DeleteProduct 实现删除产品
func (s *productService) DeleteProduct(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		log.Printf("DeleteProduct request cancelled: %v", ctx.Err())
		return ctx.Err()
	default:
	}

	if id <= 0 {
		return errors.New("invalid product id")
	}

	log.Printf("Service: deleting product with id: %d", id)
	if !s.db.DeleteProduct(id) {
		return errors.New("product not found")
	}

	return nil
}

// ReduceStock 实现减少产品库存
func (s *productService) ReduceStock(ctx context.Context, id, quantity int) error {
	select {
	case <-ctx.Done():
		log.Printf("ReduceStock request cancelled: %v", ctx.Err())
		return ctx.Err()
	default:
	}

	if id <= 0 {
		return errors.New("invalid product id")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	product := s.db.GetProduct(id)
	if product == nil {
		return errors.New("product not found")
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	log.Printf("Service: reducing stock for product id: %d, quantity: %d", id, quantity)
	product.Stock -= quantity

	return nil
}
