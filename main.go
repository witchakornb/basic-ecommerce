package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	infradb "github.com/witchakornb/basic-ecommerce/infrastructure/db" // Alias for infrastructure/db
	infrahttp "github.com/witchakornb/basic-ecommerce/infrastructure/http"
	"github.com/witchakornb/basic-ecommerce/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Database connection and migration
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Initialize Unit of Work
	uow := infradb.NewGormUnitOfWork(db) // Corrected package alias

	// Initialize repositories (still needed for handlers/usecases that do simple reads)
	userRepo := infradb.NewGormUserRepository(db) // Corrected package alias
	productRepo := infradb.NewGormProductRepository(db) // Corrected package alias

	// Initialize the use cases
	userUseCase := usecase.NewUserUseCase(userRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	// Pass the Unit of Work to the OrderUseCase
	orderUseCase := usecase.NewOrderUseCase(uow)

	// Initialize the Handlers
	userHandler := infrahttp.NewUserHandler(userUseCase)
	productHandler := infrahttp.NewProductHandler(productUseCase)
	orderHandler := infrahttp.NewOrderHandler(orderUseCase, productUseCase, userUseCase)

	// Routes and server startup
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		// User routes
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.GET("/:id", userHandler.GetUser)
			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.PUT("/:id", userHandler.UpdateUser)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}

		// Product routes
		productRoutes := api.Group("/products")
		{
			productRoutes.POST("/", productHandler.CreateProduct)
			productRoutes.GET("/:id", productHandler.GetProduct)
			productRoutes.GET("/", productHandler.GetAllProducts)
			productRoutes.PUT("/:id", productHandler.UpdateProduct)
			productRoutes.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Order routes
		orderRoutes := api.Group("/orders")
		{
			orderRoutes.POST("/", orderHandler.CreateOrder)
			orderRoutes.GET("/:id", orderHandler.GetOrder)
			orderRoutes.GET("/", orderHandler.GetAllOrders)
			orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
		}
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	log.Println("Server started on :8080")
}
