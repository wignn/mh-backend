package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, user *model.User) (*model.User, error) {
	return repository.CreateUser(db, user)
}