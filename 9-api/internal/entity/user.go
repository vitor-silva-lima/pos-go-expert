package entity

import (
	"errors"
	"first-api/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorUserNameIsRequired     = errors.New("name is required")
	ErrorUserEmailIsRequired    = errors.New("email is required")
	ErrorUserPasswordIsRequired = errors.New("password is required")
)

type User struct {
	UserID   entity.ID `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := &User{
		UserID:   entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	err = user.validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return ErrorUserNameIsRequired
	}
	if u.Email == "" {
		return ErrorUserEmailIsRequired
	}
	if u.ValidatePassword("") {
		return ErrorUserPasswordIsRequired
	}
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
