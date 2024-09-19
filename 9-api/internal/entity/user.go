package entity

import (
	"first-api/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   entity.ID `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func NewUser(name, email, password string) *User {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &User{
		UserID:   entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
