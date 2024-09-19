package repository

import (
	"first-api/internal/entity"
	database "first-api/internal/infra/database/connection"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(connection database.ConnectionAdapter) (*UserRepositoryImpl, error) {
	db, err := connection.Connect()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entity.User{})
	return &UserRepositoryImpl{db: db}, nil
}

func (r *UserRepositoryImpl) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
