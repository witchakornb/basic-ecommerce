package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	infradb "github.com/witchakornb/basic-ecommerce/infrastructure/db"
	infrahttp "github.com/witchakornb/basic-ecommerce/infrastructure/http"
	"github.com/witchakornb/basic-ecommerce/usecase"

	"github.com/glebarez/sqlite" // Pure Go SQLite driver
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the database schema
	err = db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Initialize the repositories
	userRepo := infradb.NewGormUserRepository(db)
	productRepo := infradb.NewGormProductRepository(db)
	orderRepo := infradb.NewGormOrderRepository(db)

	// Initialize the use cases
	userUseCase := usecase.NewUserUseCase(userRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, userRepo, productRepo)

	// Initialize the Handlers
	userHandler := infrahttp.NewUserHandler(userUseCase)
	productHandler := infrahttp.NewProductHandler(productUseCase)
	orderHandler := infrahttp.NewOrderHandler(orderUseCase, productUseCase, userUseCase)

	// Define / health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Define the routes
	api := router.Group("/api")
	{
		// User routes
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		// Product routes
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.GET("/products", productHandler.GetAllProducts)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)

		// Order routes
		api.POST("/orders", orderHandler.CreateOrder)
		api.GET("/orders/:id", orderHandler.GetOrderByID)
		api.GET("/orders", orderHandler.GetAllOrders)
		api.DELETE("/orders/:id", orderHandler.DeleteOrder)
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	log.Println("Server started on :8080")
}
