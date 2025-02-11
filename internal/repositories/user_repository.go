package repositories

import (
	"staj-resftul/internal/models"

	"gorm.io/gorm"
)

type Storage interface {
	GetConnection() *gorm.DB
	Close()
}

type UserRepository struct {
	storage Storage
}

func NewUserRepository(s Storage) *UserRepository {
	return &UserRepository{
		storage: s,
	}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	users := []models.User{}
	result := r.storage.GetConnection().Table("users").Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}

func (r *UserRepository) GetUserById(userId int) (*models.User, error) {
	user := models.User{}
	result := r.storage.GetConnection().Table("users").First(&user, "id = ? ", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *UserRepository) CreateUser(newUser *models.User) error {

	result := c.storage.GetConnection().Table("users").Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *UserRepository) DeleteUserByID(id int) error {
	result := c.storage.GetConnection().Table("users").Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *UserRepository) UpdateUser(id int, updatedData map[string]interface{}) error {
	if len(updatedData) == 0 {
		return nil
	}

	result := c.storage.GetConnection().Table("users").Where("id = ?", id).Updates(updatedData)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

//docker kullan TAMAM
//gerçek dbye bağlan TAMAM
//db özelliklerini elle yazmak yerin env. dosyasından al TAMAM, unit test ekle, swagger ekle
//2.gün -> user her profil sayfasına gittiğinde dbye gitmek yerine redis'te cashlensin TAMAM, (rediste varsa yoksa dbden okuyup redise yazsın)
//dosya s3'e nasıl yüklenir bak
//3.gün -> amazonda makine oluştur, redis kur, postgrl sql olutur, amazonda çalışır hale getir
//4.gün -> CI/CD github action kullanarak oto deploy yap
