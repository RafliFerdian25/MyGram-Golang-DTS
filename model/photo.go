package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null;size:100"`
	Caption  string    `json:"caption" gorm:"size:200"`
	PhotoUrl string    `json:"photo_url" gorm:"not null"`
	UserID   uint      `json:"user_id" gorm:"not null"`
	Comments []Comment `json:"comments"`
}
