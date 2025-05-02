package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	return repository.CreateBook(db, book)
}

func GetBookByQuery(db *gorm.DB, query string, page int, limit int) ([]*model.Book, error) {
	return repository.GetBookByQuery(db, query, page, limit)
}

func GetBookByID(db *gorm.DB, id *int) (*model.Book, error) {
	return repository.GetBookByID(db, id)
}

func UpdateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	return repository.UpdateBook(db, book)
}

func DeleteBook(db *gorm.DB, id int) error {
	return repository.DeleteBook(db, id)
}