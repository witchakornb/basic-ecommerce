package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/usecase"
)

// OrderHandler handles HTTP requests related to orders
type OrderHandler struct {
	orderUseCase    usecase.OrderUseCase
	productUserCase usecase.ProductUseCase
	userUsercase    usecase.UserUseCase
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(orderUseCase usecase.OrderUseCase, productUserCase usecase.ProductUseCase, userUsercase usecase.UserUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase:    orderUseCase,
		productUserCase: productUserCase,
		userUsercase:    userUsercase,
	}
}

// CreateOrder handles the creation of a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate user ID and product ID
	if _, err := h.userUsercase.GetUserByID(order.CustomerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	product, err := h.productUserCase.GetProductByID(order.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Check if product is in stock
	if product.Stock < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
		return
	}

	// Update product stock
	product.Stock -= order.Quantity
	_, err = h.productUserCase.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	// Create order

	createdOrder, err := h.orderUseCase.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

// GetOrderByID handles retrieving an order by ID
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	order, err := h.orderUseCase.GetOrderByID(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAllOrders handles retrieving all orders
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderUseCase.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// DeleteOrder handles deleting an order by ID
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.orderUseCase.DeleteOrder(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
