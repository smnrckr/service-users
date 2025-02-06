package main

import (
	"fmt"
	"os"
	"staj-resftul/internal/handlers"
	"staj-resftul/internal/repositories"
	"staj-resftul/internal/services"
	"staj-resftul/pkg/postgresql"
	"staj-resftul/pkg/redis"
	"staj-resftul/pkg/s3storage"
	"staj-resftul/utils"
	"strconv"

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

	db := postgresql.NewDB(postgresql.DbConfig{Host: host, Dbuser: dbuser, Dbname: dbname, Dbpassword: dbpassword, Port: port})

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

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	s3, err := s3storage.NewS3Service(&s3storage.S3Config{AccessKey: accessKey, SecretAccessKey: secretKey, Region: region})
	if err != nil {
		fmt.Println(err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, rdb, s3)
	userHandler := handlers.NewUserHandler(userService)
	app := fiber.New()

	userHandler.SetRoutes(app)

	app.Listen(":8080")
}
