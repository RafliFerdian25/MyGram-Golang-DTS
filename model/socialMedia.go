package model

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null;size:50"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"not null"`
}

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,url"`
	UserID         uint   `json:"user_id"`
}

type SocialMediaResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
}

type SocialMediaGetModel struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `json:"user"`
}

type SocialMediaGetResponse struct {
	ID             uint                    `json:"id"`
	Name           string                  `json:"name"`
	SocialMediaUrl string                  `json:"social_media_url"`
	UserID         uint                    `json:"user_id"`
	User           UserResponseAssociation `json:"user"`
}
