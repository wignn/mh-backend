package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *gorm.DB, user *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return repository.CreateUser(db, user)
}

func GetUserByID(db *gorm.DB, id *int) (*model.User, error) {
	return repository.GetUserByID(db, id)
}


func UpdateUser(db *gorm.DB,user *model.User) (*model.User, error) {
	return repository.UpdateUser(db, user)
}

func DeleteUser(db *gorm.DB, id *int) error {
	return repository.DeleteUser(db, id)
}