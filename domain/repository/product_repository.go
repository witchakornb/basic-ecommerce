package repository

import "github.com/witchakornb/basic-ecommerce/domain/entity"

type ProductRepository interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProductByID(id int) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(id int) error
}
