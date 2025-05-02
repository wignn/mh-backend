package model

import "gorm.io/gorm"

type Panel struct {
	gorm.Model
	PanelID   uint   `json:"panel_id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null;size:255"`
	Image     string `json:"image" gorm:"not null;size:255"`
	ChapterID uint   `json:"chapter_id" gorm:"not null"`
}
