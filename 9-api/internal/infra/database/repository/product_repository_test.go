package repository

import (
	"first-api/internal/entity"
	"first-api/internal/root"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository_Create(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewProductRepository(conn)
	assert.Nil(t, err)

	product, err := entity.NewProduct("product", 1000)
	assert.Nil(t, err)

	err = repo.Create(product)
	assert.Nil(t, err)
}

func TestProductRepository_GetByID(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewProductRepository(conn)
	assert.Nil(t, err)

	product, err := entity.NewProduct("product", 1000)
	assert.Nil(t, err)

	err = repo.Create(product)
	assert.Nil(t, err)

	product, err = repo.GetByID(product.ProductID.String())
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "product", product.Name)
	assert.Equal(t, 1000.0, product.Price)
}

func TestProductRepository_GetAll(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewProductRepository(conn)
	assert.Nil(t, err)

	products := []*entity.Product{
		{
			Name:  "product1",
			Price: 1000,
		},
		{
			Name:  "product2",
			Price: 2000,
		},
	}

	for _, product := range products {
		err = repo.Create(product)
		assert.Nil(t, err)
	}

	products, err = repo.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products))
}

func TestProductRepository_Update(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewProductRepository(conn)
	assert.Nil(t, err)

	product, err := entity.NewProduct("product", 1000)
	assert.Nil(t, err)

	err = repo.Create(product)
	assert.Nil(t, err)

	product.Name = "product updated"
	product.Price = 2000

	err = repo.Update(product)
	assert.Nil(t, err)

	product, err = repo.GetByID(product.ProductID.String())
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "product updated", product.Name)
	assert.Equal(t, 2000.0, product.Price)
}

func TestProductRepository_Delete(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewProductRepository(conn)
	assert.Nil(t, err)

	product, err := entity.NewProduct("product", 1000)
	assert.Nil(t, err)

	err = repo.Create(product)
	assert.Nil(t, err)

	err = repo.Delete(product.ProductID.String())
	assert.Nil(t, err)

	product, err = repo.GetByID(product.ProductID.String())
	assert.NotNil(t, err)
	assert.Nil(t, product)
}
