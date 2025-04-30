package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	return repository.CreateBook(db, book)
}

func GetBookByQuery(db *gorm.DB, book []*model.Book, query *string, page *int, limit *int) ([]*model.Book, error) {
	return repository.GetBookByQuery(db, book, query, page, limit)
}
