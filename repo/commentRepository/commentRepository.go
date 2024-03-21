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
func (c *CommentRepository) CreateComment(comment model.CommentRequest) (model.Comment, error) {
	commentModel := model.Comment{
		Message: comment.Message,
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
	}
	err := c.db.Create(&commentModel).Error
	if err != nil {
		return model.Comment{}, err
	}
	return commentModel, nil
}

// GetAllComments implements CommentRepository
func (c *CommentRepository) GetAllComments() ([]model.CommentGetModel, error) {
	var comments []model.CommentGetModel
	err := c.db.Model(&model.Comment{}).Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []model.CommentGetModel{}, err
	}
	return comments, nil
}

// GetCommentByID implements CommentRepository
func (c *CommentRepository) GetCommentByID(commentID uint) (model.CommentGetModel, error) {
	var comment model.CommentGetModel
	err := c.db.Model(&model.Comment{}).Preload("User").Preload("Photo").First(&comment, commentID).Error
	if err != nil {
		return model.CommentGetModel{}, err
	}
	return comment, nil
}

// UpdateComment implements CommentRepository
func (c *CommentRepository) UpdateComment(commentRequest model.CommentUpdateRequest, commentID uint) (model.Comment, error) {
	var commentModel model.Comment
	err := c.db.First(&commentModel, commentID).Error
	if err != nil {
		return model.Comment{}, err
	}

	commentModel.Message = commentRequest.Message

	err = c.db.Save(&commentModel).Error
	if err != nil {
		return model.Comment{}, err
	}
	return commentModel, nil
}

// DeleteComment implements CommentRepository
func (u *CommentRepository) DeleteComment(commentID uint) error {
	var comment model.Comment
	err := u.db.Unscoped().Delete(&comment, commentID).Error
	if err != nil {
		return err
	}
	return nil
}
