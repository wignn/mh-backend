package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Title         string  `json:"title" gorm:"not null;size:255"`
	ChapterNumber int     `json:"chapter_number" gorm:"not null"`
	BookID        uint    `json:"book_id" gorm:"not null"`
	Book          Book    `json:"book" gorm:"foreignKey:BookID"`
	Panels        []Panel `json:"panels" gorm:"foreignKey:ChapterID"`
	Pages         []string `json:"pages" gorm:"type:text[]"`
}
