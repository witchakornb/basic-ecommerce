package infrastructure

import (
	"github.com/witchakornb/basic-ecommerce/domain/repository"
	"gorm.io/gorm"
)

// gormUnitOfWork implements the UnitOfWork interface for GORM.
type gormUnitOfWork struct {
	db *gorm.DB
}

// gormUnitOfWorkStore implements the UnitOfWorkStore interface.
type gormUnitOfWorkStore struct {
	userRepo    repository.UserRepository
	productRepo repository.ProductRepository
	orderRepo   repository.OrderRepository
}

func (s *gormUnitOfWorkStore) Users() repository.UserRepository {
	return s.userRepo
}

func (s *gormUnitOfWorkStore) Products() repository.ProductRepository {
	return s.productRepo
}

func (s *gormUnitOfWorkStore) Orders() repository.OrderRepository {
	return s.orderRepo
}

// NewGormUnitOfWork creates a new GORM unit of work.
func NewGormUnitOfWork(db *gorm.DB) repository.UnitOfWork {
	return &gormUnitOfWork{db: db}
}

// Execute runs a function within a GORM transaction.
func (uow *gormUnitOfWork) Execute(fn func(store repository.UnitOfWorkStore) error) error {
	return uow.db.Transaction(func(tx *gorm.DB) error {
		store := &gormUnitOfWorkStore{
			userRepo:    NewGormUserRepository(tx),
			productRepo: NewGormProductRepository(tx),
			orderRepo:   NewGormOrderRepository(tx),
		}
		return fn(store)
	})
}
