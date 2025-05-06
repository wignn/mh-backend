package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/internal/repository"
	"gorm.io/gorm"
)

func CreateChapter(db *gorm.DB, chapter *model.Chapter) (*model.Chapter, error) {
	return repository.CreateChapter(db, chapter)
}

func GetChapterById(db *gorm.DB, id int) (*model.Chapter, error) {
	return repository.GetChapterById(db, id)
}

func GetChapterByBook(db *gorm.DB, bookID int) ([]model.Chapter, error) {
	return repository.GetChapterByBook(db, bookID)
}

func UpdateChapter(db *gorm.DB, chapter *model.Chapter) (*model.Chapter, error) {
	return repository.UpdateChapter(db, chapter)
}