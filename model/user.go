package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string        `json:"username" gorm:"unique;not null;size:50"`
	Email           string        `json:"email" gorm:"unique;not null;size:150"`
	Password        string        `json:"password" gorm:"not null"`
	Age             int           `json:"age" gorm:"not null;type:int"`
	ProfileImageUrl *string       `json:"profile_image_url"`
	SocialMedias    []SocialMedia `json:"social_medias"`
	Photos          []Photo       `json:"photos"`
	Comments        []Comment     `json:"comments"`
}
