package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `gorm:"size:255"`
	Author    string `gorm:"size:255"`
	GenreID   uint
	Genre     []Genre    `gorm:"many2many:book_genres;"`
	Bookmarks []Bookmark `gorm:"foreignKey:BookID;"`
	Users     []User     `gorm:"many2many:user_books;"`
}
