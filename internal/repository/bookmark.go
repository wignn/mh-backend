package repository

import (
	"fmt"
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateBookmark(db *gorm.DB, bookmark *model.Bookmark) (*model.Bookmark, error) {
	if err := db.Create(bookmark).Error; err != nil {
		return nil, fmt.Errorf("failed to create bookmark: %w", err)
	}
	return bookmark, nil
}

func GetBookmarksByUser(db *gorm.DB, userId string) ([]*model.Bookmark, error) {
	var bookmarks []*model.Bookmark
	if err := db.Where("user_id = ?", userId).Find(&bookmarks).Error; err != nil {
		return nil, fmt.Errorf("failed to get bookmarks for user %s: %w", userId, err)
	}
	return bookmarks, nil
}

func GetBookmarkById(db *gorm.DB, id string) (*model.Bookmark, error) {
	var bookmark model.Bookmark
	if err := db.First(&bookmark, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get bookmark by id %s: %w", id, err)
	}
	return &bookmark, nil
}

func IsBookmark(db *gorm.DB, userId int, bookId int) (*model.Bookmark, error) {
	var bookmark model.Bookmark
	if err := db.Where("user_id = ? AND book_id = ?", userId, bookId).First(&bookmark).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to check bookmark: %w", err)
	}
	return &bookmark, nil
}

func DeleteBookmark(db *gorm.DB, id int) error {
	if err := db.Delete(&model.Bookmark{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete bookmark id %d: %w", id, err)
	}
	return nil
}
