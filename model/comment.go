package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null;size:200"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
}
