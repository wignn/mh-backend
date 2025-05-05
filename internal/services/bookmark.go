package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
)

func CreateBookmark(db *gorm.DB, bookmark *model.Bookmark) (*model.Bookmark, error) {
	return repository.CreateBookmark(db, bookmark)
}

func GetBookmarksByUser(db *gorm.DB, userId string) ([]*model.Bookmark, error) {
	return repository.GetBookmarksByUser(db, userId)
}

func GetBookmarkById(db *gorm.DB, id string) (*model.Bookmark, error) {
	return repository.GetBookmarkById(db, id)
}

func IsBookmarked(db *gorm.DB, userId int, bookId int) (*model.Bookmark, error) {
	return repository.IsBookmark(db, userId, bookId)
}

func DeleteBookmark(db *gorm.DB, id int) error {
	return repository.DeleteBookmark(db, id)
}