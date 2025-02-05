package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentVariables() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}
}
