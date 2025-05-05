package services

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
	"github.com/wignn/mh-backend/internal/repository"
)

func CreateGenre(db *gorm.DB, genre string) (*model.Genre, error) {
	return repository.CreateGenre(db, genre)
}

func GetGenreById(db *gorm.DB, id int) (*model.Genre, error) {
	return repository.GetGenreById(db, id)
}

func GetGenres(db *gorm.DB) ([]*model.Genre, error) {
	return repository.GetGenres(db)
}