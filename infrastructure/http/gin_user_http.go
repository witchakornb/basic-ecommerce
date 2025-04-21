package infrastructure

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/witchakornb/basic-ecommerce/usecase"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUser handles the creation of a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userUseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}


// GetUserByID handles retrieving a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles updating a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.userUseCase.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles deleting a user by ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.userUseCase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
