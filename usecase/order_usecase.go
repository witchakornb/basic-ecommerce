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

// OrderUseCaseImpl is the implementation of OrderUseCase
type OrderUseCaseImpl struct {
	uow repository.UnitOfWork // เปลี่ยนจาก repo แต่ละตัวมาเป็น UoW
}

// NewOrderUseCase creates a new OrderUseCase
func NewOrderUseCase(uow repository.UnitOfWork) OrderUseCase {
	return &OrderUseCaseImpl{
		uow: uow,
	}
}

func (o *OrderUseCaseImpl) CreateOrder(order entity.Order) (createdOrder entity.Order, err error) { // Modified return to named
	err = o.uow.Execute(func(store repository.UnitOfWorkStore) error {
		// 1. Get repositories from the store
		userRepo := store.Users()
		productRepo := store.Products()
		orderRepo := store.Orders()

		// 2. Check if user exists
		user, err := userRepo.GetUserByID(order.CustomerID)
		if err != nil || user.ID == 0 {
			return errors.New("user not found")
		}

		// 3. Check if product exists
		product, err := productRepo.GetProductByID(order.ProductID)
		if err != nil {
			return errors.New("product not found")
		}

		// 4. Check if product is in stock
		if product.Stock < order.Quantity {
			return errors.New("not enough stock")
		}

		// 5. Update product stock (within transaction)
		product.Stock -= order.Quantity
		_, err = productRepo.UpdateProduct(product)
		if err != nil {
			return errors.New("failed to update product stock")
		}

		// 6. Create order (within transaction)
		createdOrder, err = orderRepo.CreateOrder(order)
		if err != nil {
			return err
		}

		return nil // No error, transaction will be committed
	})

	return createdOrder, err
}

// ----- (Optional but recommended) Update other methods to use UoW as well -----

func (o *OrderUseCaseImpl) GetOrderByID(id int) (order entity.Order, err error) { // Modified return to named
	err = o.uow.Execute(func(store repository.UnitOfWorkStore) error {
		var err error
		order, err = store.Orders().GetOrderByID(id)
		if err != nil {
			return errors.New("order not found")
		}
		return nil
	})
	return order, err
}

func (o *OrderUseCaseImpl) GetAllOrders() (orders []entity.Order, err error) { // Modified return to named
	err = o.uow.Execute(func(store repository.UnitOfWorkStore) error {
		var err error
		orders, err = store.Orders().GetAllOrders()
		return err
	})
	return orders, err
}

func (o *OrderUseCaseImpl) DeleteOrder(id int) error {
	return o.uow.Execute(func(store repository.UnitOfWorkStore) error {
		err := store.Orders().DeleteOrder(id)
		if err != nil {
			return errors.New("failed to delete order")
		}
		return nil
	})
}
