package repository

import (
	"first-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryMemory_Create(t *testing.T) {
	repo, err := NewUserRepositoryMemory()
	assert.NoError(t, err)
	user, err := entity.NewUser("John Doe", "j@d.com", "123456")
	assert.NoError(t, err)
	err = repo.Create(user)
	assert.NoError(t, err)
	userFind := &entity.User{}
	err = repo.db.Where("user_id = ?", user.UserID.String()).First(userFind).Error
	assert.NoError(t, err)
	assert.Equal(t, user.UserID, userFind.UserID)
	assert.Equal(t, user.Name, userFind.Name)
	assert.Equal(t, user.Email, userFind.Email)
	assert.Equal(t, user.Password, userFind.Password)

}

func TestUserRepositoryMemory_GetByEmail(t *testing.T) {
	repo, err := NewUserRepositoryMemory()
	assert.NoError(t, err)
	user, err := entity.NewUser("John Doe", "j@d.com", "123456")
	assert.NoError(t, err)
	err = repo.Create(user)
	assert.NoError(t, err)
	userFind, err := repo.GetByEmail("j@d.com")
	assert.NoError(t, err)
	assert.Equal(t, user.UserID, userFind.UserID)
	assert.Equal(t, user.Name, userFind.Name)
	assert.Equal(t, user.Email, userFind.Email)
	assert.Equal(t, user.Password, userFind.Password)
}
