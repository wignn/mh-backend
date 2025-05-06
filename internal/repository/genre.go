package repository

import (
	"fmt"

	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateGenre(db *gorm.DB, title string) (*model.Genre, error) {
	genre := &model.Genre{Title: title}
	if err := db.Create(genre).Error; err != nil {
		return nil, fmt.Errorf("failed to create genre '%s': %w", title, err)
	}
	return genre, nil
}

func GetGenreById(db *gorm.DB, id int) (*model.Genre, error) {
	var genre model.Genre
	if err := db.First(&genre, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get genre by id %d: %w", id, err)
	}
	return &genre, nil
}

func GetGenres(db *gorm.DB) ([]*model.Genre, error) {
	var genres []*model.Genre
	if err := db.Find(&genres).Error; err != nil {
		return nil, fmt.Errorf("failed to get genres: %w", err)
	}
	return genres, nil
}
