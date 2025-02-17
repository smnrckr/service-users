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
	result := r.storage.GetConnection().Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}

func (r *UserRepository) GetUserById(userId int) (*models.User, error) {
	user := models.User{}
	result := r.storage.GetConnection().First(&user, "id = ? ", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(newUser *models.User) error {

	result := r.storage.GetConnection().Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) DeleteUserByID(id int) error {
	result := r.storage.GetConnection().Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) UpdateUserById(id int, updatedData models.User) (models.User, error) {

	result := r.storage.GetConnection().Where("id = ?", id).Updates(&updatedData)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, models.ErrorNoRowsAffected
	}
	return updatedData, nil

}

//docker kullan TAMAM
//gerçek dbye bağlan TAMAM
//db özelliklerini elle yazmak yerin env. dosyasından al TAMAM, unit test ekle TAMAM, swagger ekle TAMAM
//2.gün -> user her profil sayfasına gittiğinde dbye gitmek yerine redis'te cashlensin TAMAM, (rediste varsa yoksa dbden okuyup redise yazsın)
//dosya s3'e nasıl yüklenir bak TAMAM
//3.gün -> amazonda makine oluştur TAMAM, redis kur TAMAM, postgrl sql olutur TAMAM, amazonda çalışır hale getir
//4.gün -> CI/CD github action kullanarak oto deploy yap
