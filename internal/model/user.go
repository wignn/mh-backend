package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255"`
	Email    string `gorm:"unique;size:255"`
	Password string `gorm:"size:255"`
	Bookmark []Book `gorm:"many2many:user_books;"`
	IsAdmin  bool   `gorm:"default:false"`
	bookmark []Bookmark `gorm:"foreignKey:UserID"`
}


type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}
