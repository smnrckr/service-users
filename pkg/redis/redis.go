package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"staj-resftul/internal/models"
	"time"

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

func (r *RedisDB) GetUsersFromCache() ([]models.User, error) {
	redisResult, err := r.RedisClient.Get("users").Result()
	if err != nil || redisResult == "" {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal([]byte(redisResult), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *RedisDB) SaveUsersToCache(users []models.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return r.RedisClient.Set("users", data, time.Minute*5).Err()
}
