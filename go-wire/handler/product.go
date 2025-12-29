package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yiwen/go-wire/model"
	"yiwen/go-wire/service"
)

// ProductHandler 商品 HTTP 处理器
type ProductHandler struct {
	svc service.ProductService
}

// NewProductHandler 创建商品处理器（这是一个 Provider）
// 注意：它依赖 ProductService，Wire 会自动注入
func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

// RegisterRoutes 注册路由
func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/products", h.GetProducts)
		api.POST("/products", h.CreateProduct)
	}
}

// GetProducts 获取所有商品
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.svc.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct 创建商品
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.svc.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": product})
}
