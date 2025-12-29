package repository

import (
	"sync"

	"yiwen/go-wire/model"
)

// ProductRepository 商品数据访问接口
type ProductRepository interface {
	FindAll() ([]*model.Product, error)
	Create(product *model.Product) error
}

// productRepository 商品数据访问实现（内存存储）
type productRepository struct {
	mu       sync.RWMutex
	products map[int64]*model.Product
	nextID   int64
}

// NewProductRepository 创建商品仓库（这是一个 Provider）
func NewProductRepository() ProductRepository {
	return &productRepository{
		products: make(map[int64]*model.Product),
		nextID:   1,
	}
}

func (r *productRepository) FindAll() ([]*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]*model.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepository) Create(product *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.nextID++
	r.products[product.ID] = product
	return nil
}
