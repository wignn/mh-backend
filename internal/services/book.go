package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	err := db.Create(&book).Error
	return book, err
}