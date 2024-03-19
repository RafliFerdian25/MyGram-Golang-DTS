package model

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null;size:50"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"not null"`
}
