package repository

import (
	"first-api/internal/entity"
	"first-api/internal/root"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewUserRepository(conn)
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

func TestUserRepository_GetByEmail(t *testing.T) {
	conn := root.NewDatabaseConnectionAdapter()
	repo, err := NewUserRepository(conn)
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
