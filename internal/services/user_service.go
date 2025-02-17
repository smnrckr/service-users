package services

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"staj-resftul/internal/models"
	"time"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUserById(userId int) (*models.User, error)
	CreateUser(newUser *models.User) error
	DeleteUserByID(id int) error
	UpdateUserById(userId int, updatedData models.User) (models.User, error)
}

type RedisDB interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type S3Service interface {
	UploadFile(BucketName string, Key string, Body []byte) (string, error)
}

type UserService struct {
	userRepository UserRepository
	redisDB        RedisDB
	s3Service      S3Service
}

func NewUserService(repository UserRepository, redisdb RedisDB, s3Service S3Service) *UserService {
	return &UserService{
		userRepository: repository,
		redisDB:        redisdb,
		s3Service:      s3Service,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.getUsersFromCache()
	if err == nil {
		return users, nil
	}
	fmt.Println("cacheden veri gelmedi")
	users, err = s.userRepository.GetUsers()
	fmt.Println("users dbden getirildi")
	if err != nil {
		return nil, err
	}

	err = s.saveUsersToCache(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserById(userId int) (*models.User, error) {

	redisResult, err := s.redisDB.Get("users")
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

func (s *UserService) UpdateUserById(userId int, updatedData models.User) (models.User, error) {

	newData, err := s.userRepository.UpdateUserById(userId, updatedData)
	if err != nil {
		return models.User{}, err
	}

	return newData, nil

}

func (s *UserService) getUsersFromCache() ([]models.User, error) {
	redisResult, err := s.redisDB.Get("users")
	if err != nil || redisResult == "" {
		return nil, err
	}

	users := []models.User{}
	if err := json.Unmarshal([]byte(redisResult), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) saveUsersToCache(users []models.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return s.redisDB.Set("users", data, time.Minute*5)
}
