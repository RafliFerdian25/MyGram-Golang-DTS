package photoController

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/service/photoService"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PhotoController struct {
	PhotoService *photoService.PhotoService
}

func NewPhotoController(photoService *photoService.PhotoService) *PhotoController {
	return &PhotoController{
		PhotoService: photoService,
	}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	var photoRequest model.PhotoCreateRequest
	err := ctx.Bind(&photoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data photo
	validator := helper.NewValidator()

	err = validator.Validate(photoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// get photo data from token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoRequest.UserID = userID

	// call service to create photo
	photoResponse, err := p.PhotoService.CreatePhoto(photoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail create photo",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, photoResponse)
}

// get all photo
func (p *PhotoController) GetAllPhotos(ctx *gin.Context) {
	photos, err := p.PhotoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get photos",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

// get photo by id
func (p *PhotoController) GetPhotoByID(ctx *gin.Context) {
	paramPhotoID := ctx.Param("id")
	photoID, err := strconv.Atoi(paramPhotoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid photo id",
			"error":   err.Error(),
		})
		return
	}

	photo, err := p.PhotoService.GetPhotoByID(uint(photoID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail get photo",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
// 	// bind request data
// 	var photoRequest model.PhotoUpdateRequest
// 	err := ctx.Bind(&photoRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "fail bind data",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	// validate data photo
// 	validator := helper.NewValidator()
// 	err = validator.Validate(photoRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request format",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	// get photo data from token
// 	photoData := ctx.MustGet("photoData").(jwt.MapClaims)
// 	photoID := uint(photoData["id"].(float64))

// 	// call service to update photo
// 	photoResponse, err := p.PhotoService.UpdatePhoto(photoRequest, photoID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "fail update photo",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, photoResponse)
// }

// func (u *PhotoController) DeletePhoto(ctx *gin.Context) {
// 	photoData := ctx.MustGet("photoData").(jwt.MapClaims)
// 	photoID := uint(photoData["id"].(float64))

// 	err := p.PhotoService.DeletePhoto(photoID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "fail delete photo",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "success delete photo",
// 	})
// }
