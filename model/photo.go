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

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
	UserID   uint   `json:"user_id"`
}

type PhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type PhotoGetModel struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
}

type PhotoGetResponse struct {
	ID       uint                    `json:"id"`
	Title    string                  `json:"title"`
	Caption  string                  `json:"caption"`
	PhotoUrl string                  `json:"photo_url"`
	UserID   uint                    `json:"user_id"`
	User     UserResponseAssociation `json:"user"`
}

type PhotoResponseAssociation struct {
	ID       uint   `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
