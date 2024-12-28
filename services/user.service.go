package services

import (
	"CRUD_Gin/database"
	"CRUD_Gin/models"
	"log"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService() *UserService {
	db := database.DBConn()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}
	return &UserService{DB: db}
}

func (us *UserService) FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := us.DB.Find(&users)
	return users, result.Error
}

func (us *UserService) FindUserById(cccd string) (*models.User, error) {
	var user models.User
	result := us.DB.First(&user, "cccd = ?", cccd)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (us *UserService) Create(user *models.User) error {
	result := us.DB.Create(user)
	return result.Error
}
