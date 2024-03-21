package middleware

import (
	"MyGram-Golang-DTS/service/commentService"
	"MyGram-Golang-DTS/service/photoService"
	"MyGram-Golang-DTS/service/socialMediaService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CommentAuthorization(commentService *commentService.CommentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get comment id from param
		commentID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid comment id",
				"error":   err.Error(),
			})
			return
		}

		// get user data from token
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		// call service to check comment owner
		err = commentService.CheckCommentOwner(uint(commentID), userID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Comment not found",
					"error":   err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		ctx.Next()
	}
}

func PhotoAuthorization(photoService *photoService.PhotoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get photo id from param
		photoID, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id social media",
				"error":   err.Error(),
			})
			return
		}

		// get user data from token
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		// call service to check photo owner
		err = photoService.CheckPhotoOwner(uint(photoID), userID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Photo not found",
					"error":   err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		ctx.Next()
	}
}

func SocialMediaAuthorization(socialMediaService *socialMediaService.SocialMediaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get socialMedia id from param
		socialMediaID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id social media",
				"error":   err.Error(),
			})
			return
		}

		// get user data from token
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		// call service to check socialMedia owner
		err = socialMediaService.CheckSocialMediaOwner(uint(socialMediaID), userID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Social media not found",
					"error":   err.Error(),
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		ctx.Next()
	}
}
