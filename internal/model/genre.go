package model


import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Title string `json:"title" gorm:"not null;size:255"`
	Books []Book `gorm:"foreignKey:GenreID"`
}


type BookGenre struct {
    BookID  uint
    GenreID uint
}