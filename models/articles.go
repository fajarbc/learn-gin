package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" binding:"min=3,max=255"`
	Content   string `json:"content" binding:"max=1000"`
	Slug      string `json:"slug"`
	Status    uint8  `json:"status" binding:"required"`
	AuthorID  uint
	Author    Author    `gorm:"foreognKey:AuthorID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ArticleInsert struct {
	Title   string `json:"title" binding:"min=3,max=255"`
	Content string `json:"content" binding:"max=1000"`
	Slug    string `json:"slug"`
	Status  uint8  `json:"status" binding:"required"`
}
