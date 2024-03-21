package commentService

import (
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/commentRepository"
	"MyGram-Golang-DTS/repo/photoRepository"
	"errors"

	"github.com/jinzhu/copier"
)

type CommentService struct {
	commentRepo *commentRepository.CommentRepository
	photoRepo   *photoRepository.PhotoRepository
}

func NewCommentService(commentRepository *commentRepository.CommentRepository, photoRepository *photoRepository.PhotoRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepository,
		photoRepo:   photoRepository,
	}
}

// CreateComment implements CommentService
func (c *CommentService) CreateComment(commentRequest model.CommentRequest) (model.CommentResponse, error) {
	// check if photo exists
	_, err := c.photoRepo.GetPhotoByID(commentRequest.PhotoID)
	if err != nil {
		return model.CommentResponse{}, errors.New("photo not found")
	}

	// call repository to save comment
	createdComment, err := c.commentRepo.CreateComment(commentRequest)
	if err != nil {
		return model.CommentResponse{}, err
	}

	// copy data from createdComment to commentResponse
	var commentResponse model.CommentResponse
	err = copier.Copy(&commentResponse, &createdComment)
	if err != nil {
		return model.CommentResponse{}, err
	}

	return commentResponse, nil
}

// get all comments
func (c *CommentService) GetAllComments() ([]model.CommentGetResponse, error) {
	// call repository to get all comments
	comments, err := c.commentRepo.GetAllComments()
	if err != nil {
		return []model.CommentGetResponse{}, err
	}

	var commentResponses []model.CommentGetResponse
	err = copier.Copy(&commentResponses, &comments)
	if err != nil {
		return []model.CommentGetResponse{}, err
	}

	return commentResponses, nil
}

// get comment by id
func (c *CommentService) GetCommentByID(commentID uint) (model.CommentGetResponse, error) {
	// call repository to get comment by id
	comment, err := c.commentRepo.GetCommentByID(commentID)
	if err != nil {
		return model.CommentGetResponse{}, err
	}

	var commentResponse model.CommentGetResponse
	err = copier.Copy(&commentResponse, &comment)
	if err != nil {
		return model.CommentGetResponse{}, err
	}

	return commentResponse, nil
}

// update comment
func (c *CommentService) UpdateComment(commentRequest model.CommentUpdateRequest, commentID uint) (model.CommentResponse, error) {
	// call repository to update comment
	updatedComment, err := c.commentRepo.UpdateComment(commentRequest, commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	var commentResponse model.CommentResponse
	err = copier.Copy(&commentResponse, &updatedComment)
	if err != nil {
		return model.CommentResponse{}, err
	}

	return commentResponse, nil
}

// delete comment
func (c *CommentService) DeleteComment(commentID uint) error {
	// call repository to delete comment
	err := c.commentRepo.DeleteComment(commentID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommentService) CheckCommentOwner(commentID uint, userID uint) error {
	comment, err := c.commentRepo.GetCommentByID(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return errors.New("comment not belongs to user")
	}
	return nil
}
