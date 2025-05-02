package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateGenre(db *gorm.DB, genre string) (*model.Genre, error) {
	var newGenre model.Genre
	newGenre.Title = genre

	err := db.Create(&newGenre).Error
	if err != nil {
		return nil, err
	}

	return &newGenre, nil
}

func GetGenreById(db *gorm.DB, id int) (*model.Genre, error) {
	var genre model.Genre
	err := db.First(&genre, id).Error
	return &genre, err
}