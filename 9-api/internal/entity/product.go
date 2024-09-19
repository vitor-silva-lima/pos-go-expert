package entity

import (
	"errors"
	"first-api/pkg/entity"
	"time"
)

var (
	ErrorProductIDIsRequired    = errors.New("product_id is required")
	ErrorInvalidProductID       = errors.New("invalid product_id")
	ErrorProductNameIsRequired  = errors.New("name is required")
	ErrorProductPriceIsRequired = errors.New("price is required")
	ErrorInvalidProductPrice    = errors.New("invalid price")
)

type Product struct {
	ProductID entity.ID `json:"product_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ProductID: entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) validate() error {
	if p.ProductID.String() == "" {
		return ErrorProductIDIsRequired
	}
	if _, err := entity.StringToID(p.ProductID.String()); err != nil {
		return ErrorInvalidProductID
	}
	if p.Name == "" {
		return ErrorProductNameIsRequired
	}
	if p.Price == 0 {
		return ErrorProductPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorInvalidProductPrice
	}
	return nil
}
