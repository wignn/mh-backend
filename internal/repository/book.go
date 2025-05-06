package repository

import (
	"fmt"

	"github.com/wignn/mh-backend/internal/model"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}
	return book, nil
}

func GetBookByQuery(db *gorm.DB, query string, limit int, offset int) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64

	tx := db.Model(&model.Book{}).Where("title ILIKE ?", "%"+query+"%")

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count books: %w", err)
	}

	if err := tx.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get books: %w", err)
	}

	return books, total, nil
}

func GetBookByID(db *gorm.DB, id *int) (*model.Book, error) {
	var book model.Book
	if err := db.Preload("Chapters").Preload("Genres").First(&book, *id).Error; err != nil {
		return nil, fmt.Errorf("failed to get book by id %d: %w", *id, err)
	}
	return &book, nil
}

func UpdateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	var existing model.Book
	if err := db.First(&existing, book.ID).Error; err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	if err := db.Model(&existing).Updates(book).Error; err != nil {
		return nil, fmt.Errorf("failed to update book id %d: %w", book.ID, err)
	}

	return &existing, nil
}

func DeleteBook(db *gorm.DB, id int) error {
	if err := db.Delete(&model.Book{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete book id %d: %w", id, err)
	}
	return nil
}

func CreateBookGenre(db *gorm.DB, bookGenre *model.BookGenre) (*model.BookGenre, error) {
	if err := db.Create(bookGenre).Error; err != nil {
		return nil, fmt.Errorf("failed to create book-genre relation: %w", err)
	}
	return bookGenre, nil
}
