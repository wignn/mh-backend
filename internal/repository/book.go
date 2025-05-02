package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"github.com/wignn/mh-backend/pkg/utils"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	err := db.Create(&book).Error
	return book, err
}


func GetBookByQuery(db *gorm.DB, query string, page int, limit int) ([]*model.Book, error) {
	var book []*model.Book
	currentPage := 1
	currentLimit := 10

	offset, currentLimit := utils.Paginate(currentPage, currentLimit)

	tx := db.Model(&model.Book{})

	if query != "" {
		tx = tx.Where("title ILIKE ?", "%"+query+"%")
	}
	
	tx = tx.Offset(offset).Limit(currentLimit)

	if err := tx.Find(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}


func GetBookByID(db *gorm.DB, id *int) (*model.Book, error) {
	var book *model.Book
	err := db.Where(&id).First(&book).Error
	return book, err
}

func UpdateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	var existingBook model.Book
	if err := db.First(&existingBook, book.ID).Error; err != nil {
		return nil, err
	}
	
	if err := db.Model(&existingBook).Updates(book).Error; err != nil {
		return nil, err
	}
	
	return &existingBook, nil
}