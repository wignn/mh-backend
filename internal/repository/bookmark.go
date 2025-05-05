package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateBookmark(db *gorm.DB, bookmark *model.Bookmark) (*model.Bookmark, error) {
	err := db.Create(&bookmark).Error
	return bookmark, err
}

func GetBookmarksByUser(db *gorm.DB, userId string) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	err := db.Where("userID = ?", userId).Find(&bookmarks).Error
	return bookmarks, err
}

func GetBookmarkById(db *gorm.DB, id string) (*model.Bookmark, error) {
	var bookmark model.Bookmark
	err := db.First(&bookmark, id).Error
	return &bookmark, err
}

func IsBookmark(db *gorm.DB, userId int, bookId int) (*model.Bookmark, error) {
	var bookmark model.Bookmark
	err := db.Where("userID = ? AND bookID = ?", userId, bookId).First(&bookmark).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No bookmark found
		}
		return nil, err // Some other error occurred
	}
	return &bookmark, nil
}

func DeleteBookmark(db *gorm.DB, id int) error {
	return db.Delete(&model.Bookmark{}, id).Error
}