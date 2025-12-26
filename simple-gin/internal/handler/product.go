package handler

import (
	"example/simple-gin/internal/model"
	"example/simple-gin/internal/service"
	"example/simple-gin/pkg/response"
	"log"
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

// GetProducts godoc
//
//	@Summary		获取所有产品
//	@Description	获取产品列表
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]model.Product}
//	@Failure		500	{object}	response.Response
//	@Router			/api/v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	products, err := h.productService.GetProducts(ctx)
	if err != nil {
		log.Printf("Handler: error getting products: %v", err)
		response.InternalError(c, "failed to get products: "+err.Error())
		return
	}

	response.Success(c, products)
}

// GetProduct godoc
//
//	@Summary		获取单个产品
//	@Description	根据ID获取产品详情
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"产品ID"
//	@Success		200	{object}	response.Response{data=model.Product}
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Router			/api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.GetProductByID(ctx, id)
	if err != nil {
		log.Printf("Handler: error getting product %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, product)
}

// CreateProduct godoc
//
//	@Summary		创建产品
//	@Description	创建一个新产品
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		model.CreateProductRequest	true	"产品信息"
//	@Success		201		{object}	response.Response{data=model.Product}
//	@Failure		400		{object}	response.Response
//	@Router			/api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.CreateProduct(ctx, &req)
	if err != nil {
		log.Printf("Handler: error creating product: %v", err)
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, product)
}

// UpdateProduct godoc
//
//	@Summary		更新产品
//	@Description	根据ID更新产品信息
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"产品ID"
//	@Param			product	body		model.UpdateProductRequest	true	"更新信息"
//	@Success		200		{object}	response.Response{data=model.Product}
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Router			/api/v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	product, err := h.productService.UpdateProduct(ctx, id, &req)
	if err != nil {
		log.Printf("Handler: error updating product %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, product)
}

// DeleteProduct godoc
//
//	@Summary		删除产品
//	@Description	根据ID删除产品
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"产品ID"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Router			/api/v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.productService.DeleteProduct(ctx, id)
	if err != nil {
		log.Printf("Handler: error deleting product %d: %v", id, err)
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// ReduceStock godoc
//
//	@Summary		减少库存
//	@Description	减少产品库存数量
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"产品ID"
//	@Param			request	body		ReduceStockRequest	true	"减少数量"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Router			/api/v1/products/{id}/reduce-stock [post]
func (h *ProductHandler) ReduceStock(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "invalid product id")
		return
	}

	var req ReduceStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return
	}

	ctx, cancel := createContextWithTimeout(c, 5*time.Second)
	defer cancel()

	err = h.productService.ReduceStock(ctx, id, req.Quantity)
	if err != nil {
		log.Printf("Handler: error reducing stock for product %d: %v", id, err)
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// ReduceStockRequest 减少库存请求
type ReduceStockRequest struct {
	Quantity int `json:"quantity" binding:"required,gt=0" example:"10"`
}
