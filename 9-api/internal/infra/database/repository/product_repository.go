package repository

import (
	"first-api/internal/entity"
	database "first-api/internal/infra/database/connection"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(conn database.ConnectionAdapter) (*ProductRepositoryImpl, error) {
	db, err := conn.Connect()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entity.Product{})
	return &ProductRepositoryImpl{db: db}, nil
}

func (r *ProductRepositoryImpl) Create(user *entity.Product) error {
	return r.db.Create(user).Error
}

func (r *ProductRepositoryImpl) GetByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Where("product_id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll() ([]*entity.Product, error) {
	var products []*entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) Update(user *entity.Product) error {
	_, err := r.GetByID(user.ProductID.String())
	if err != nil {
		return err
	}
	return r.db.Where("product_id = ?", user.ProductID.String()).Save(user).Error
}

func (r *ProductRepositoryImpl) Delete(id string) error {
	product, err := r.GetByID(id)
	if err != nil {
		return err
	}
	return r.db.Where("product_id = ?", product.ProductID.String()).Delete(product).Error
}
