package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser("John Doe", "john.doe@mail.com", "123456")
	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "John Doe")
	assert.Equal(t, user.Email, "john.doe@mail.com")
	assert.NotEqual(t, user.Password, "123456")
}

func TestUser_ValidatePassword(t *testing.T) {
	user := NewUser("John Doe", "john.doe@mail.com", "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
}
