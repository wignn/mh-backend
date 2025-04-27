package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(user).Error
	return user, err
}