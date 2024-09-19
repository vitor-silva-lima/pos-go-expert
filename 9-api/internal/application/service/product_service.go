package service

import (
	"first-api/internal/application/repository"
	"first-api/internal/dto"
	"first-api/internal/entity"
)

type ProductService interface {
	GetProducts() ([]*entity.Product, error)
	CreateProduct(input *dto.CreateProductDtoInput) error
	UpdateProduct(input *dto.UpdateProductDtoInput) error
	DeleteProduct(id string) error
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (ps *ProductServiceImpl) GetProducts() ([]*entity.Product, error) {
	return ps.ProductRepository.GetAll()
}

func (ps *ProductServiceImpl) CreateProduct(input *dto.CreateProductDtoInput) error {
	product, err := entity.NewProduct(input.Name, input.Price)
	if err != nil {
		return err
	}
	return ps.ProductRepository.Create(product)
}

func (ps *ProductServiceImpl) UpdateProduct(input *dto.UpdateProductDtoInput) error {
	product, err := ps.ProductRepository.GetByID(string(input.ProductID))
	if err != nil {
		return err
	}
	product.Name = input.Name
	product.Price = input.Price
	return ps.ProductRepository.Update(product)
}

func (ps *ProductServiceImpl) DeleteProduct(id string) error {
	return ps.ProductRepository.Delete(id)
}
