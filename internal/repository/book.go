package repository

import (
	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	err := db.Create(&book).Error
	return book, err
}

func GetBookByQuery(db *gorm.DB, query string, limit int, offset int) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64

	tx := db.Model(&model.Book{}).
		Where("title ILIKE ?", "%"+query+"%")
	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = tx.Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, total, err
}

func GetBookByID(db *gorm.DB, id *int) (*model.Book, error) {
	var book *model.Book
	err := db.Preload("Chapters").Preload("Genres").Where("id = ?", *id).First(&book).Error
	return book, err
}

func UpdateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	var existingBook model.Book
	err := db.First(&existingBook, book.ID).Error
	err = db.Model(&existingBook).Updates(book).Error
	return &existingBook, err
}

func DeleteBook(db *gorm.DB, id int) error {
	return db.Delete(&model.Book{}, id).Error
}

func CreateBookGenre(db *gorm.DB, book *model.BookGenre) (*model.BookGenre, error) {
	err := db.Create(&book).Error
	return book, err
}
