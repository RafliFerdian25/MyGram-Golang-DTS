package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string        `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email           string        `json:"email" gorm:"uniqueIndex;not null;size:150"`
	Password        string        `json:"password" gorm:"not null"`
	Age             int           `json:"age" gorm:"not null;type:int"`
	ProfileImageUrl string        `json:"profile_image_url"`
	SocialMedias    []SocialMedia `json:"social_medias" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Photos          []Photo       `json:"photos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments        []Comment     `json:"comments" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserRequest struct {
	Username        string `json:"username" form:"username" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=6"`
	Age             int    `json:"age" form:"age" validate:"required,numeric,gt=8"`
	ProfileImageUrl string `json:"profile_image_url" form:"profile_image_url" validate:"omitempty,url"`
}

type UserResponse struct {
	ID              uint    `json:"id"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	Age             int     `json:"age"`
	ProfileImageUrl *string `json:"profile_image_url"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Username        string `json:"username" form:"username" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Age             int    `json:"age" form:"age" validate:"required,numeric,gt=8"`
	ProfileImageUrl string `json:"profile_image_url" form:"profile_image_url" validate:"omitempty,url"`
}
