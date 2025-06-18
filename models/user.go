package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID    uint           `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"size:100;not null;unique"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	Photo     string         `json:"photo,omitempty"` // URL to user's photo
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete
}
