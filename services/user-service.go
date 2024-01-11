package services

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser() ([]models.User, error)
	GetUserById(ID int) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ru *userRepository) GetUser() ([]models.User, error) {
	var users []models.User
	err := ru.db.Find(&users).Error
	return users, err
}

func (ru *userRepository) GetUserById(ID int) (models.User, error) {
	var user models.User
	err := ru.db.First(&user, ID).Error
	return user, err
}
