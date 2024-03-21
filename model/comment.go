package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null;size:200"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
}

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id" validate:"required"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

type CommentGetModel struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PhotoID uint   `json:"photo_id"`
	Photo   Photo  `json:"photo"`
}

type CommentGetResponse struct {
	ID      uint                     `json:"id"`
	Message string                   `json:"message"`
	UserID  uint                     `json:"user_id"`
	User    UserResponseAssociation  `json:"user"`
	PhotoID uint                     `json:"photo_id"`
	Photo   PhotoResponseAssociation `json:"photo"`
}
