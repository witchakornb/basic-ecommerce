package infrastructure

import (
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
	"gorm.io/gorm"
)

// GormOrderRepository is a struct that implements the OrderRepository interface
type GormOrderRepository struct {
	db *gorm.DB
}

// NewGormOrderRepository creates a new instance of GormOrderRepository
func NewGormOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &GormOrderRepository{db: db}
}

// CreateOrder creates a new order in the database
func (r *GormOrderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// GetOrderByID retrieves an order by ID from the database
func (r *GormOrderRepository) GetOrderByID(id int) (entity.Order, error) {
	var order entity.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// GetAllOrders retrieves all orders from the database
func (r *GormOrderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// DeleteOrder deletes an order by ID from the database
func (r *GormOrderRepository) DeleteOrder(id int) error {
	var order entity.Order
	err := r.db.Delete(&order, id).Error
	if err != nil {
		return err
	}
	return nil
}
