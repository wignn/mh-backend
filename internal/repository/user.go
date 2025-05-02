package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func GetUserByID(db *gorm.DB, id *int) (*model.User, error) {
	var user model.User
	err := db.First(&user, *id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Save(&user).Error
	return user, err
}

func DeleteUser(db *gorm.DB, id *int) error {
	err := db.Delete(&model.User{}, id).Error
	return err
}

func GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
