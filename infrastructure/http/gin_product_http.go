package infrastructure

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/usecase"
)

// ProductHandler handles HTTP requests related to products
type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productUseCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

// CreateProduct handles the creation of a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.productUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// GetProductByID handles retrieving a product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productUseCase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts handles retrieving all products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productUseCase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct handles updating a product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.productUseCase.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles deleting a product by ID
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.productUseCase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}