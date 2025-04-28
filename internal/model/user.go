package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique;size:255"`
	Password string `gorm:"size:255"`
	Bookmark []Book `gorm:"many2many:user_books;"`
	isAdmin  bool   `gorm:"default:false"`
}
