package utils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis bağlantı hatası: %v", err)
	} else {
		log.Println("Redis başarıyla bağlandı.")
	}
}

func GetRedisClient() *redis.Client {
	return rdb
}
