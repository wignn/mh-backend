package config

import (
	"log"
	"os"
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB = database
	database.AutoMigrate(&model.Book{}, &model.User{}, &model.Genre{}, &model.Bookmark{}, &model.BookGenre{}, &model.BookGenre{}, &model.Chapter{})
}