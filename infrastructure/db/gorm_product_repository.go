package infrastructure

import (
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
	"gorm.io/gorm"
)

// GormProductRepository is a GORM implementation of the ProductRepository interface.
type GormProductRepository struct {
	db *gorm.DB
}

// NewGormProductRepository creates a new GormProductRepository instance.
func NewGormProductRepository(db *gorm.DB) repository.ProductRepository {
	return &GormProductRepository{db: db}
}

// CreateProduct creates a new product in the database.
func (r *GormProductRepository) CreateProduct(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// GetProductByID retrieves a product by ID from the database.
func (r *GormProductRepository) GetProductByID(id int) (entity.Product, error) {
	var product entity.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// GetAllProducts retrieves all products from the database.
func (r *GormProductRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct updates an existing product in the database.
func (r *GormProductRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// DeleteProduct deletes a product by ID from the database.
func (r *GormProductRepository) DeleteProduct(id int) error {
	var product entity.Product
	err := r.db.Delete(&product, id).Error
	if err != nil {
		return err
	}
	return nil
}
