package repository

import (
	"first-api/internal/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepositoryMemory struct {
	db *gorm.DB
}

func NewUserRepositoryMemory() (*UserRepositoryMemory, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entity.User{})
	return &UserRepositoryMemory{db: db}, nil
}

func (r *UserRepositoryMemory) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryMemory) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
