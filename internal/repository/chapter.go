package repository

import (
	"fmt"

	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateChapter(db *gorm.DB, chapter *model.Chapter) (*model.Chapter, error) {
	if err := db.Create(chapter).Error; err != nil {
		return nil, fmt.Errorf("failed to create chapter: %w", err)
	}
	return chapter, nil
}

func GetChapterById(db *gorm.DB, id int) (*model.Chapter, error) {
	var chapter model.Chapter
	if err := db.First(&chapter, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get chapter by id %d: %w", id, err)
	}
	return &chapter, nil
}

func GetChapterByBook(db *gorm.DB, bookID int) ([]model.Chapter, error) {
	var chapters []model.Chapter
	if err := db.Where("book_id = ?", bookID).Find(&chapters).Error; err != nil {
		return nil, fmt.Errorf("failed to get chapters for book id %d: %w", bookID, err)
	}
	return chapters, nil
}

func UpdateChapter(db *gorm.DB, chapter *model.Chapter) (*model.Chapter, error) {
	if err := db.Save(chapter).Error; err != nil {
		return nil, fmt.Errorf("failed to update chapter id %d: %w", chapter.ID, err)
	}
	return chapter, nil
}
