package commentRepository

import (
	"MyGram-Golang-DTS/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

// CreateComment implements CommentRepository
func (u *CommentRepository) CreateComment(comment model.CommentRequest) (model.Comment, error) {
	commentModel := model.Comment{
		Message: comment.Message,
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
	}
	err := u.db.Create(&commentModel).Error
	if err != nil {
		return model.Comment{}, err
	}
	return commentModel, nil
}

// GetAllComments implements CommentRepository
func (u *CommentRepository) GetAllComments() ([]model.CommentGetModel, error) {
	var comments []model.CommentGetModel
	err := u.db.Model(&model.Comment{}).Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []model.CommentGetModel{}, err
	}
	return comments, nil
}

// GetCommentByID implements CommentRepository
// func (u *CommentRepository) GetCommentByID(commentID uint) (model.CommentGetModel, error) {
// 	var comment model.CommentGetModel
// 	err := u.db.Model(&model.Comment{}).Preload("User").First(&comment, commentID).Error
// 	if err != nil {
// 		return model.CommentGetModel{}, err
// 	}
// 	return comment, nil
// }

// UpdateComment implements CommentRepository
// func (u *CommentRepository) UpdateComment(comment model.CommentRequest, commentID uint) (model.Comment, error) {
// 	var commentModel model.Comment
// 	err := u.db.First(&commentModel, commentID).Error
// 	if err != nil {
// 		return model.Comment{}, err
// 	}

// 	commentModel.Title = comment.Title
// 	commentModel.Caption = comment.Caption
// 	commentModel.CommentUrl = comment.CommentUrl

// 	err = u.db.Save(&commentModel).Error
// 	if err != nil {
// 		return model.Comment{}, err
// 	}
// 	return commentModel, nil
// }

// DeleteComment implements CommentRepository
// func (u *CommentRepository) DeleteComment(commentID uint) error {
// 	var comment model.Comment
// 	err := u.db.Unscoped().Delete(&comment, commentID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
