package usecase

import (
	"errors"

	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
)

type OrderUseCase interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrderByID(id int) (entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	DeleteOrder(id int) error
}

type OrderUseCaseImpl struct {
	OrderRepo   repository.OrderRepository
	UserRepo    repository.UserRepository
	ProductRepo repository.ProductRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository, userRepo repository.UserRepository, productRepo repository.ProductRepository) OrderUseCase {
	return &OrderUseCaseImpl{
		OrderRepo:   orderRepo,
		UserRepo:    userRepo,
		ProductRepo: productRepo,
	}
}

func (o *OrderUseCaseImpl) CreateOrder(order entity.Order) (entity.Order, error) {
	// Check if user exists
	user, err := o.UserRepo.GetUserByID(order.CustomerID)
	if err != nil || user.ID == 0 {
		return entity.Order{}, errors.New("user not found")
	}

	// Check if product exists
	product, err := o.ProductRepo.GetProductByID(order.ProductID)
	if err != nil {
		return entity.Order{}, errors.New("product not found")
	}

	// Check if product is in stock
	if product.Stock < order.Quantity {
		return entity.Order{}, errors.New("not enough stock")
	}

	// Update product stock
	product.Stock -= order.Quantity
	_, err = o.ProductRepo.UpdateProduct(product)
	if err != nil {
		return entity.Order{}, errors.New("failed to update product stock")
	}

	// Create order
	order, err = o.OrderRepo.CreateOrder(order)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (o *OrderUseCaseImpl) GetOrderByID(id int) (entity.Order, error) {
	order, err := o.OrderRepo.GetOrderByID(id)
	if err != nil {
		return entity.Order{}, errors.New("order not found")
	}

	return order, nil
}

func (o *OrderUseCaseImpl) GetAllOrders() ([]entity.Order, error) {
	orders, err := o.OrderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderUseCaseImpl) DeleteOrder(id int) error {
	err := o.OrderRepo.DeleteOrder(id)
	if err != nil {
		return errors.New("failed to delete order")
	}

	return nil
}
