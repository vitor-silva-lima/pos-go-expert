package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "John Doe")
	assert.Equal(t, user.Email, "john.doe@mail.com")
	assert.NotEqual(t, user.Password, "123456")
}

func TestUserWhenNameIsEmpty(t *testing.T) {
	user, err := NewUser("", "john.doe@mail.com", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorUserNameIsRequired)
}

func TestUserWhenEmailIsEmpty(t *testing.T) {
	user, err := NewUser("John Doe", "", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorUserEmailIsRequired)
}

func TestUserWhenPasswordIsEmpty(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@mail.com", "")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.NotNil(t, ErrorUserPasswordIsRequired)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@mail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
}
