package handlers

import (
	"example/simple-gin/models"
	"example/simple-gin/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ProductHandler 产品处理器
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler 创建产品处理器实例
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProducts 获取所有产品
func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	products, err := h.productService.GetProducts(ctx)
	if err != nil {
		log.Printf("Handler: error getting products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "failed to get products: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": products,
	})
}

// GetProduct 获取单个产品
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid product id",
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		log.Printf("Handler: error getting product %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": product,
	})
}

// CreateProduct 创建产品
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid request body: " + err.Error(),
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.CreateProduct(ctx, &req)
	if err != nil {
		log.Printf("Handler: error creating product: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "product created successfully",
		"data": product,
	})
}

// UpdateProduct 更新产品
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid product id",
		})
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid request body: " + err.Error(),
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.UpdateProduct(ctx, id, &req)
	if err != nil {
		log.Printf("Handler: error updating product %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "product updated successfully",
		"data": product,
	})
}

// DeleteProduct 删除产品
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid product id",
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.productService.DeleteProduct(ctx, id)
	if err != nil {
		log.Printf("Handler: error deleting product %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "product deleted successfully",
	})
}

// ReduceStock 减少产品库存（额外操作示例）
func (h *ProductHandler) ReduceStock(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid product id",
		})
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid request body: " + err.Error(),
		})
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.productService.ReduceStock(ctx, id, req.Quantity)
	if err != nil {
		log.Printf("Handler: error reducing stock for product %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "stock reduced successfully",
	})
}
