package model

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	BookID     uint   `gorm:"not null"`
	UserID     string `gorm:"not null"`
}
