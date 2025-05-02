package model

import "gorm.io/gorm"

type BookStatus string

const (
	BookStatusOngoing   BookStatus = "Ongoing"
	BookStatusCompleted BookStatus = "Completed"
	BookStatusDropped   BookStatus = "Dropped"
)

type Book struct {
	gorm.Model
	Title       string     `gorm:"size:255"`
	Author      string     `gorm:"size:255"`
	Description string     `gorm:"size:255"`
	Cover       string     `gorm:"size:255"`
	Genres      []Genre    `gorm:"many2many:book_genres;"`
	Bookmarks   []Bookmark `gorm:"foreignKey:BookID"`
	Users       []User     `gorm:"many2many:user_books;"`
	RealesedAt  string     `gorm:"size:255"`
	Status      BookStatus `gorm:"type:varchar(20);default:'Ongoing'"`
	Chapters    []Chapter  `gorm:"foreignKey:BookID"`
}
