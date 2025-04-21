package infrastructure

import (
	"gorm.io/gorm"
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
)

// GormUserRepository is a struct that implements the UserRepository interface
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormOderRepository creates a new instance of GormUserRepository
func NewGormOderRepository(db *gorm.DB) repository.OrderRepository {
	return &GormUserRepository{db: db}
}

// CreateOrder creates a new order in the database
func (r *GormUserRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// GetOrderByID retrieves an order by ID from the database
func (r *GormUserRepository) GetOrderByID(id string) (entity.Order, error) {
	var order entity.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// GetAllOrders retrieves all orders from the database
func (r *GormUserRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}	

// DeleteOrder deletes an order by ID from the database
func (r *GormUserRepository) DeleteOrder(id string) error {
	var order entity.Order
	err := r.db.Delete(&order, id).Error
	if err != nil {
		return err
	}
	return nil
}