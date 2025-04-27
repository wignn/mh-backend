package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)



func CreateUser(db *gorm.DB, user *model.User) (*model.UserResponse, error) {
	err := db.Create(user).Error
	userResponse := model.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
	return &userResponse, err
}

func GetUserByID(db *gorm.DB, id *int) (*model.UserResponse, error) {
	var user model.UserResponse
	err := db.First(&user, id).Error
	userResponse := model.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	return &userResponse, err
}

func UpdateUser(db *gorm.DB, id *int, user *model.User) (*model.User, error) {
	err := db.Save(user).Error
	return user, err
}

func DeleteUser(db *gorm.DB, id *int) error {
	err := db.Delete(&model.User{}, id).Error
	return err
}
