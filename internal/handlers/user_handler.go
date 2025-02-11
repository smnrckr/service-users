package handlers

import (
	"fmt"
	"staj-resftul/internal/models"
	"staj-resftul/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) handleGetUsers(c *fiber.Ctx) error {
	user, err := h.userService.GetUsers()
	if err != nil {
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func (h *UserHandler) handleGetUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	user, err := h.userService.GetUserById(userId)

	if err != nil {
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func (h *UserHandler) handleCreateUser(c *fiber.Ctx) error {
	userRequest := new(models.UserCreateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON verisi"})
	}

	file, _ := c.FormFile("profile_photo")

	user, err := h.userService.CreateUser(userRequest, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": user})
}

func (h *UserHandler) handleDeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}
	err = h.userService.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kullanıcı başarıyla silindi", "user_id": id})
}

func (h *UserHandler) handleUpdateUser(c *fiber.Ctx) error {
	userRequest := models.UserUpdateRequest{}
	err := c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON verisi"})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	updatedData := map[string]interface{}{}
	if userRequest.Name != "" {
		updatedData["name"] = userRequest.Name
	}
	if userRequest.Lastname != "" {
		updatedData["lastname"] = userRequest.Lastname
	}
	if userRequest.ProfilePhoto != "" {
		updatedData["profile_photo"] = userRequest.ProfilePhoto
	}

	err = h.userService.UpdateUser(id, updatedData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kullanıcı başarıyla güncellendi"})
}

func (h *UserHandler) SetRoutes(app *fiber.App) {
	app.Get("/users", h.handleGetUsers)
	app.Get("/users/:id<int>", h.handleGetUserById)
	app.Post("/users", h.handleCreateUser)
	app.Delete("/users/:id", h.handleDeleteUser)
	app.Put("/users/:id", h.handleUpdateUser)

}
