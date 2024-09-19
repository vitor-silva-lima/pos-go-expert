package repository

import "first-api/internal/entity"

type ProductRepository interface {
	Create(user *entity.Product) error
	GetByID(id string) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Update(user *entity.Product) error
	Delete(id string) error
}
