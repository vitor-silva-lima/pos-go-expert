package repository

import (
	"first-api/internal/entity"
	database "first-api/internal/infra/database/connection"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connection database.ConnectionAdapter) (*UserRepository, error) {
	db, err := connection.Connect()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entity.User{})
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
