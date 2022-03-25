package models

import (
	"time"
)

type Author struct {
	ID        int
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" validate:"email"`
	Status    uint8     `json:"status" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
