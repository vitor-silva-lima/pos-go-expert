package service

import (
	"first-api/internal/dto"
	"first-api/internal/infra/database/repository"
	"first-api/internal/root"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	productRepository, err := repository.NewProductRepository(conn)
	if err != nil {
		panic(err)
	}
	productService := NewProductService(productRepository)
	input := &dto.CreateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 1",
			Price: 1000.0,
		},
	}
	err = productService.CreateProduct(input)
	assert.Nil(t, err)
}

func TestProductServiceImpl_GetProducts(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	productRepository, err := repository.NewProductRepository(conn)
	if err != nil {
		panic(err)
	}
	productService := NewProductService(productRepository)

	var inputs []*dto.CreateProductDtoInput
	inputs = append(inputs, &dto.CreateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 1",
			Price: 1000.0,
		},
	})
	inputs = append(inputs, &dto.CreateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 2",
			Price: 2000.0,
		},
	})
	for _, input := range inputs {
		err = productService.CreateProduct(input)
		assert.Nil(t, err)
	}

	products, err := productService.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, inputs[0].Name, products[0].Name)
	assert.Equal(t, inputs[0].Price, products[0].Price)
	assert.Equal(t, inputs[1].Name, products[1].Name)
	assert.Equal(t, inputs[1].Price, products[1].Price)
}

func TestProductServiceImpl_UpdateProduct(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	productRepository, err := repository.NewProductRepository(conn)
	if err != nil {
		panic(err)
	}
	productService := NewProductService(productRepository)

	input := &dto.CreateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 1",
			Price: 1000.0,
		},
	}
	err = productService.CreateProduct(input)
	assert.Nil(t, err)

	products, err := productService.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products))

	updateInput := &dto.UpdateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 2",
			Price: 2000.0,
		},
		ProductID: string(products[0].ProductID.String()),
	}
	err = productService.UpdateProduct(updateInput)
	assert.Nil(t, err)

	products, err = productService.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, updateInput.Name, products[0].Name)
	assert.Equal(t, updateInput.Price, products[0].Price)
}

func TestProductServiceImpl_DeleteProduct(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	productRepository, err := repository.NewProductRepository(conn)
	if err != nil {
		panic(err)
	}
	productService := NewProductService(productRepository)

	input := &dto.CreateProductDtoInput{
		ProductDtoInputBase: dto.ProductDtoInputBase{
			Name:  "Product 1",
			Price: 1000.0,
		},
	}
	err = productService.CreateProduct(input)
	assert.Nil(t, err)

	products, err := productService.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products))

	err = productService.DeleteProduct(string(products[0].ProductID.String()))
	assert.Nil(t, err)

	products, err = productService.GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 0, len(products))
}
