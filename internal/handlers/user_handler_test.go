package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"staj-resftul/internal/handlers"
	"staj-resftul/internal/models"
	"staj-resftul/internal/repositories"
	"staj-resftul/internal/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
	mockRedis := handlers.NewMockRedis()
	mockS3 := new(handlers.MockS3Storage)
	userRepo := repositories.NewUserRepository(testDb)
	userService := services.NewUserService(userRepo, mockRedis, mockS3)
	handler := handlers.NewUserHandler(userService)

	app := fiber.New()

	handler.SetRoutes(app)

	t.Run("Create User", func(t *testing.T) {
		request := models.UserCreateRequest{
			Name:     "sema",
			Lastname: "çakır",
		}
		requestJSON, err := json.Marshal(request)
		if err != nil {
			t.Fatalf("Error marshaling request: %v", err)
		}

		req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(requestJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		jsonDataFromHttp, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)

		user := models.User{}
		err = json.Unmarshal(jsonDataFromHttp, &user)
		assert.NoError(t, err)
		fmt.Println(user)
		assert.Equal(t, 201, resp.StatusCode)
		assert.Equal(t, "sema", user.Name)

	})

	t.Run("Get All Users", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		users := []models.User{}
		err = json.Unmarshal(jsonDataFromHttp, &users)
		assert.NoError(t, err)
		fmt.Println(users)
		assert.NotEmpty(t, users)

	})

	t.Run("Update Favorite Lists", func(t *testing.T) {
		newData := models.UserUpdateRequest{
			Name:         "seda",
			Lastname:     "çakır",
			ProfilePhoto: "newphoto.jpg",
		}
		requestJSON, err := json.Marshal(newData)
		if err != nil {
			t.Fatalf("Error marshaling request: %v", err)
		}

		req := httptest.NewRequest("PUT", "/users/1", bytes.NewBuffer(requestJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		userUpdated := models.User{}
		err = json.Unmarshal(jsonDataFromHttp, &userUpdated)
		assert.NoError(t, err)

		assert.Equal(t, "seda", userUpdated.Name)

	})

	t.Run("Get User By Id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		user := models.User{}
		err = json.Unmarshal(jsonDataFromHttp, &user)
		assert.NoError(t, err)

		assert.NotEmpty(t, user)

	})

	t.Run("Delete User By Id", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

	})

}
