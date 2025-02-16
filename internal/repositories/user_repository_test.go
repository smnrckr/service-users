package repositories_test

import (
	"staj-resftul/internal/models"
	"staj-resftul/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavoritesListsRepository(t *testing.T) {
	userRepo := repositories.NewUserRepository(testDb)
	t.Run("Create User", func(t *testing.T) {
		user := models.User{
			Name:     "Semanur",
			Lastname: "Çakır",
		}
		err := userRepo.CreateUser(&user)
		assert.NoError(t, err)
		assert.Equal(t, 1, user.Id)
	})

	t.Run("Get Users", func(t *testing.T) {
		users, err := userRepo.GetUsers()
		assert.NoError(t, err)
		assert.NotEmpty(t, users)
		assert.Len(t, users, 1)

	})

	t.Run("Get User By Id", func(t *testing.T) {
		user, err := userRepo.GetUserById(1)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, "Semanur", user.Name)

	})

	t.Run("UpdateFavoriteList", func(t *testing.T) {

		updatedData := models.User{
			Name: "Sema",
		}

		updatedUser, err := userRepo.UpdateUserById(1, updatedData)
		assert.NoError(t, err)

		user, err := userRepo.GetUserById(1)

		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, updatedUser.Name, user.Name)
	})

	t.Run("DeleteFavoriteList", func(t *testing.T) {

		err := userRepo.DeleteUserByID(1)
		assert.NoError(t, err)
		user, err := userRepo.GetUserById(1)
		assert.Error(t, err)
		assert.Empty(t, user)

	})
}
