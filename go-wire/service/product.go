package service

import (
	"yiwen/go-wire/model"
	"yiwen/go-wire/repository"
)

// ProductService 商品业务逻辑接口
type ProductService interface {
	GetAllProducts() ([]*model.Product, error)
	CreateProduct(req *model.CreateProductRequest) (*model.Product, error)
}

// productService 商品业务逻辑实现
type productService struct {
	repo repository.ProductRepository
}

// NewProductService 创建商品服务（这是一个 Provider）
// 注意：它依赖 ProductRepository，Wire 会自动注入
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetAllProducts() ([]*model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) CreateProduct(req *model.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	if err := s.repo.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}
