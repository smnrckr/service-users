package main

import (
	"fmt"
	"os"
	"staj-resftul/internal/handlers"
	"staj-resftul/internal/repositories"
	"staj-resftul/internal/services"
	"staj-resftul/pkg/postgresql"
	"staj-resftul/pkg/redis"
	"staj-resftul/utils"
	"strconv"

	//"github.com/go-redis/redis/v8"

	"github.com/gofiber/fiber/v2"
)

func init() {
	utils.LoadEnviromentVariables()
}

func main() {

	host := os.Getenv("HOST")
	dbuser := os.Getenv("USER_NAME")
	dbname := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDBstr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBstr)
	if err != nil {
		fmt.Println(err)
	}
	rdb := redis.NewClient(redis.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		Db:       redisDB,
	})

	db := postgresql.NewDB(postgresql.DbConfig{Host: host, Dbuser: dbuser, Dbname: dbname, Dbpassword: dbpassword, Port: port})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, rdb)
	userHandler := handlers.NewUserHandler(userService)
	app := fiber.New()

	userHandler.SetRoutes(app)

	app.Listen(":8080")
}
