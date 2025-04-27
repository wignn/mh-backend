package model

import "gorm.io/gorm"

type Bookmark struct {
    gorm.Model
    PageNumber int   `gorm:"not null"`
    BookID     uint 
    Book       Book  `gorm:"foreignKey:BookID"` 
}