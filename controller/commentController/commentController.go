package commentController

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/service/commentService"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CommentController struct {
	CommentService *commentService.CommentService
}

func NewCommentController(commentService *commentService.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var commentRequest model.CommentRequest
	err := ctx.Bind(&commentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data comment
	validator := helper.NewValidator()

	err = validator.Validate(commentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// get comment data from token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentRequest.UserID = userID

	// call service to create comment
	commentResponse, err := c.CommentService.CreateComment(commentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail create comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, commentResponse)
}

// get all comment
func (p *CommentController) GetAllComments(ctx *gin.Context) {
	comments, err := p.CommentService.GetAllComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get comments",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// get comment by id
func (c *CommentController) GetCommentByID(ctx *gin.Context) {
	paramCommentID := ctx.Param("id")
	commentID, err := strconv.Atoi(paramCommentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid comment id",
			"error":   err.Error(),
		})
		return
	}

	comment, err := c.CommentService.GetCommentByID(uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// update comment
func (p *CommentController) UpdateComment(ctx *gin.Context) {
	paramCommentID := ctx.Param("id")
	commentID, err := strconv.Atoi(paramCommentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid comment id",
			"error":   err.Error(),
		})
		return
	}

	// bind data comment
	var commentRequest model.CommentUpdateRequest
	err = ctx.Bind(&commentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data comment
	validator := helper.NewValidator()
	err = validator.Validate(commentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// get comment data from token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// call service to update comment
	commentResponse, err := p.CommentService.UpdateComment(commentRequest, uint(commentID), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail update comment",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, commentResponse)
}

// delete comment
// func (p *CommentController) DeleteComment(ctx *gin.Context) {
// 	// get comment id from param
// 	paramCommentID := ctx.Param("id")
// 	commentID, err := strconv.Atoi(paramCommentID)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid comment id",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	// get user data from token
// 	userData := ctx.MustGet("userData").(jwt.MapClaims)
// 	userID := uint(userData["id"].(float64))

// 	err = p.CommentService.DeleteComment(uint(commentID), userID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "fail delete comment",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "success delete comment",
// 	})
// }
