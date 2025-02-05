package services

import (
	"encoding/json"
	"staj-resftul/internal/models"
	"staj-resftul/internal/repositories"
	"staj-resftul/pkg/redis"
	"time"
)

type UserService struct {
	userRepository *repositories.UserRepository
	redisDB        *redis.RedisDB
}

func NewUserService(repository *repositories.UserRepository, redisdb *redis.RedisDB) *UserService {
	return &UserService{
		userRepository: repository,
		redisDB:        redisdb,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	redisResult := s.redisDB.RedisClient.Get("users")
	users := []models.User{}
	json.Unmarshal([]byte(redisResult.String()), &users)

	if len(users) != 0 {
		return users, nil
	}

	result, err := s.userRepository.GetUsers()
	if err != nil {
		return []models.User{}, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return []models.User{}, err
	}

	err = s.redisDB.RedisClient.Set("users", data, time.Minute*5).Err()
	if err != nil {
		return []models.User{}, err
	}

	return result, nil
}

func (s *UserService) CreateUser(newUser *models.User) error {
	return s.userRepository.CreateUser(newUser)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUserByID(id)
}

func (s *UserService) UpdateUser(id int, updatedData map[string]interface{}) error {
	return s.userRepository.UpdateUser(id, updatedData)
}
