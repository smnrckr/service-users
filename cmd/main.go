package main

import (
	"os"
	"staj-resftul/internal/handlers"
	"staj-resftul/internal/repositories"
	"staj-resftul/internal/services"
	"staj-resftul/pkg/postgresql"
	"staj-resftul/utils"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
	Version string `json:"verison"`
}

func init() {
	utils.LoadEnviromentVariables()
}

func main() {

	host := os.Getenv("HOST")
	dbuser := os.Getenv("USER_NAME")
	dbname := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	db := postgresql.NewDB(postgresql.DbConfig{Host: host, Dbuser: dbuser, Dbname: dbname, Dbpassword: dbpassword, Port: port})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	app := fiber.New()

	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Hello World",
	// 		"verison": "1.0.0",
	// 	})
	// })

	userHandler.SetRoutes(app)

	app.Listen(":8080")
}
