package socialMediaController

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/service/socialMediaService"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SocialMediaController struct {
	SocialMediaService *socialMediaService.SocialMediaService
}

func NewSocialMediaController(socialMediaService *socialMediaService.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: socialMediaService,
	}
}

func (c *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var socialMediaRequest model.SocialMediaRequest
	err := ctx.Bind(&socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data socialMedia
	validator := helper.NewValidator()
	err = validator.Validate(socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// get socialMedia data from token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMediaRequest.UserID = userID

	// call service to create socialMedia
	socialMediaResponse, err := c.SocialMediaService.CreateSocialMedia(socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail create socialMedia",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, socialMediaResponse)
}

// get all socialMedia
func (p *SocialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	// get socialMedia data from token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	socialMedias, err := p.SocialMediaService.GetAllSocialMedias(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get socialMedias",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

// get socialMedia by id
func (c *SocialMediaController) GetSocialMediaByID(ctx *gin.Context) {
	paramSocialMediaID := ctx.Param("id")
	socialMediaID, err := strconv.Atoi(paramSocialMediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid socialMedia id",
			"error":   err.Error(),
		})
		return
	}

	socialMedia, err := c.SocialMediaService.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get social media",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

// update socialMedia
func (p *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	paramSocialMediaID := ctx.Param("id")
	socialMediaID, err := strconv.Atoi(paramSocialMediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid socialMedia id",
			"error":   err.Error(),
		})
		return
	}

	// bind data socialMedia
	var socialMediaRequest model.SocialMediaRequest
	err = ctx.Bind(&socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data socialMedia
	validator := helper.NewValidator()
	err = validator.Validate(socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// call service to update socialMedia
	socialMediaResponse, err := p.SocialMediaService.UpdateSocialMedia(socialMediaRequest, uint(socialMediaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail update socialMedia",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMediaResponse)
}

// delete socialMedia
func (p *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	// get socialMedia id from param
	paramSocialMediaID := ctx.Param("id")
	socialMediaID, err := strconv.Atoi(paramSocialMediaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid social media id",
			"error":   err.Error(),
		})
		return
	}

	err = p.SocialMediaService.DeleteSocialMedia(uint(socialMediaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail delete social media",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success delete social media",
	})
}
