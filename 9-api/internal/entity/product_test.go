package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Laptop", 1000.0)
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Laptop")
	assert.Equal(t, product.Price, 1000.0)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("", 1000.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorProductNameIsRequired)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	product, err := NewProduct("Laptop", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorProductPriceIsRequired)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	product, err := NewProduct("Laptop", -1000.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorInvalidProductPrice)
}
