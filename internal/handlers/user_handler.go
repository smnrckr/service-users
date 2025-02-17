package handlers

import (
	"fmt"
	"mime/multipart"
	"staj-resftul/internal/models"

	"github.com/go-swagno/swagno/components/endpoint"
	"github.com/go-swagno/swagno/components/http/response"
	"github.com/go-swagno/swagno/components/mime"
	"github.com/go-swagno/swagno/components/parameter"
	"github.com/gofiber/fiber/v2"
)

type UserServiceInterface interface {
	GetUsers() ([]models.User, error)
	GetUserById(userId int) (*models.User, error)
	CreateUser(req *models.UserCreateRequest, file *multipart.FileHeader) (*models.User, error)
	DeleteUser(id int) error
	UpdateUserById(userId int, updatedData models.User) (models.User, error)
}

type UserHandler struct {
	userService UserServiceInterface
}

func NewUserHandler(service UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) handleGetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) handleGetUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) handleCreateUser(c *fiber.Ctx) error {
	userRequest := new(models.UserCreateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: err.Error()})
	}
	if err := userRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	file, _ := c.FormFile("profile_photo")

	user, err := h.userService.CreateUser(userRequest, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) handleDeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Geçersiz kullanıcı Id"})
	}
	err = h.userService.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{Message: "user deleted successfuly"})
}

func (h *UserHandler) handleUpdateUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Geçersiz kullanıcı ID"})
	}

	userRequest := models.UserUpdateRequest{}
	err = c.BodyParser(&userRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: err.Error()})
	}

	if err := userRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedData := models.User{
		Name:         userRequest.Name,
		Lastname:     userRequest.Lastname,
		ProfilePhoto: userRequest.ProfilePhoto,
	}

	newData, err := h.userService.UpdateUserById(id, updatedData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(newData)
}

func (h *UserHandler) SetRoutes(app *fiber.App) {

	userGroup := app.Group("/users")
	userGroup.Get("/", h.handleGetUsers)
	userGroup.Get("/:id<int>", h.handleGetUserById)
	userGroup.Post("/", h.handleCreateUser)
	userGroup.Delete("/:id<int>", h.handleDeleteUser)
	userGroup.Put("/:id<int>", h.handleUpdateUserById)

}

var UserEndpoints = []*endpoint.EndPoint{
	endpoint.New(
		endpoint.GET,
		"/users",
		endpoint.WithTags("user"),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			[]models.User{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "500")}),
		endpoint.WithDescription("tüm kullanıcıları döner"),
	),
	endpoint.New(
		endpoint.GET,
		"/users/{id}",
		endpoint.WithTags("user"),
		endpoint.WithParams(parameter.IntParam("id", parameter.Path, parameter.WithRequired(), parameter.WithDescription("Kullanıcı id"))),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			models.User{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "500")}),
		endpoint.WithDescription("id'e göre kullanıcı döner"),
	),

	endpoint.New(
		endpoint.POST,
		"/users",
		endpoint.WithTags("user"),
		endpoint.WithBody(models.UserCreateRequest{}),
		endpoint.WithConsume([]mime.MIME{mime.MULTIFORM}),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			models.User{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "500")}),
		endpoint.WithDescription("Yeni kullanıcı oluşturur"),
	),
	endpoint.New(
		endpoint.DELETE,
		"/users/{id}",
		endpoint.WithTags("user"),
		endpoint.WithParams(parameter.IntParam("id", parameter.Path, parameter.WithRequired(), parameter.WithDescription("Kullanıcı id"))),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(models.SuccessResponse{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "500")}),
		endpoint.WithDescription("Kullanıcı silinir"),
	),
	endpoint.New(
		endpoint.PUT,
		"/users/{id}",
		endpoint.WithTags("user"),
		endpoint.WithParams(parameter.IntParam("id", parameter.Path, parameter.WithRequired())),
		endpoint.WithBody(models.UserUpdateRequest{}),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			models.User{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "500")}),
		endpoint.WithDescription("Kullanıcı bilgilerini günceller"),
	),
}
