package repository

import (
	"fmt"

	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id *int) (*model.User, error) {
	var user model.User
	if err := db.First(&user, *id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by id %d: %w", *id, err)
	}
	return &user, nil
}

func UpdateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user id %d: %w", user.ID, err)
	}
	return user, nil
}

func DeleteUser(db *gorm.DB, id *int) error {
	if err := db.Delete(&model.User{}, *id).Error; err != nil {
		return fmt.Errorf("failed to delete user id %d: %w", *id, err)
	}
	return nil
}

func GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by username '%s': %w", username, err)
	}
	return &user, nil
}
