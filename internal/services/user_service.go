package services

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"staj-resftul/internal/models"
	"staj-resftul/internal/repositories"
	"staj-resftul/pkg/redis"
	"staj-resftul/pkg/s3storage"
	"time"
)

type UserService struct {
	userRepository *repositories.UserRepository
	redisDB        *redis.RedisDB
	s3Service      *s3storage.S3Service
}

func NewUserService(repository *repositories.UserRepository, redisdb *redis.RedisDB, s3Service *s3storage.S3Service) *UserService {
	return &UserService{
		userRepository: repository,
		redisDB:        redisdb,
		s3Service:      s3Service,
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

func (s *UserService) GetUserById(userId int) (*models.User, error) {

	redisResult, err := s.redisDB.RedisClient.Get("users").Result()
	users := []models.User{}
	if err == nil && redisResult != "" {
		if err := json.Unmarshal([]byte(redisResult), &users); err == nil {
			for _, user := range users {
				if user.Id == userId {
					return &user, nil
				}
			}
		}
	}

	user, err := s.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) CreateUser(req *models.UserCreateRequest, file *multipart.FileHeader) (*models.User, error) {
	var fileURL string

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		fileBytes := make([]byte, file.Size)
		_, err = src.Read(fileBytes)
		if err != nil {
			return nil, err
		}

		fileKey := fmt.Sprintf("profile_photos/%d_%s", time.Now().Unix(), file.Filename)
		fileURL, err = s.s3Service.UploadFile("cimristaj", fileKey, fileBytes)
		if err != nil {
			return nil, err
		}
	}

	user := &models.User{
		Name:         req.Name,
		Lastname:     req.Lastname,
		ProfilePhoto: fileURL,
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUserByID(id)
}

func (s *UserService) UpdateUser(id int, updatedData map[string]interface{}) error {
	return s.userRepository.UpdateUser(id, updatedData)
}
