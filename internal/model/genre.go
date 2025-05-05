package model


import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Title string `json:"title" gorm:"not null;size:255"`
	Books []Book `gorm:"many2many:book_genres;"`
}


type BookGenre struct {
	gorm.Model
    BookID  uint
    GenreID uint
}


