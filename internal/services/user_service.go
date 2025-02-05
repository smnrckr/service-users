package services

import (
	"fmt"
	"staj-resftul/internal/models"
	"staj-resftul/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	result, err := s.userRepository.GetUsers()
	if err != nil {
		fmt.Println(err)
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
