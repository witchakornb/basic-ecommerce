package repository


import "github.com/witchakornb/basic-ecommerce/domain/entity"

type OrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrderByID(id string) (entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	DeleteOrder(id string) error
}