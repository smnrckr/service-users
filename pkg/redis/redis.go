package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

type RedisDB struct {
	RedisClient *redis.Client
}

func NewClient(config RedisConfig) *RedisDB {

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	}
	fmt.Println("connection successful")

	return &RedisDB{RedisClient: rdb}
}
