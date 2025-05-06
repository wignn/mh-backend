package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255"`
	Email    string `gorm:"unique;size:255"`
	Password string `gorm:"size:255"`
	IsAdmin  bool   `gorm:"default:false"`
	Bookmarks []Bookmark 
}


type LoginRequest struct {
    Username    string `json:"username" binding:"required,username"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}
