package usecase

import (
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
)

type ProductUseCase interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}

type ProductUseCaseImpl struct {
	ProductRepo repository.ProductRepository
}

func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		ProductRepo: productRepo,
	}
}

func (p *ProductUseCaseImpl) CreateProduct(product entity.Product) (entity.Product, error) {
	product, err := p.ProductRepo.CreateProduct(product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *ProductUseCaseImpl) GetProductByID(id string) (entity.Product, error) {
	product, err := p.ProductRepo.GetProductByID(id)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *ProductUseCaseImpl) GetAllProducts() ([]entity.Product, error) {
	products, err := p.ProductRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUseCaseImpl) UpdateProduct(product entity.Product) (entity.Product, error) {
	product, err := p.ProductRepo.UpdateProduct(product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *ProductUseCaseImpl) DeleteProduct(id string) error {
	err := p.ProductRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
