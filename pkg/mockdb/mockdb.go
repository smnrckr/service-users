package mockdb

import "staj-resftul/internal/models"

type MockDB struct {
	usersMap map[int]models.User
}

func NewMockDB() *MockDB {
	return &MockDB{
		usersMap: map[int]models.User{
			1: {Id: 1, Name: "Semanur", Lastname: "Çakır"},
			2: {Id: 2, Name: "Sedanur", Lastname: "Çakır"},
		},
	}
}

func (m *MockDB) GetUsers() []models.User {
	users := []models.User{}
	for _, user := range m.usersMap {
		users = append(users, user)
	}
	return users
}
