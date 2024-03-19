package userController

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/userService"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *userService.UserService
}

func NewUserController(userService *userService.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var userRequest model.UserRequest
	err := ctx.Bind(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail bind data",
			"error":   err.Error(),
		})
		return
	}

	// validate data user
	validator := helper.NewValidator()

	err = validator.Validate(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	userResponse, err := u.UserService.CreateUser(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail create user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, userResponse)
}

// func (u *UserController) LoginUser(c gin.Context) error {
// 	var user model.User
// 	err := c.Bind(&user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, gin.Map{
// 			"message": "fail bind data",
// 			"error":   err.Error(),
// 		})
// 	}

// 	user, err = u.UserService.LoginUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, gin.Map{
// 			"message": "fail login",
// 			"error":   err.Error(),
// 		})
// 	}

// 	token, errToken := helper.CreateToken(user.ID, user.Name)

// 	if errToken != nil {
// 		return c.JSON(http.StatusInternalServerError, gin.Map{
// 			"message": "fail create token",
// 			"error":   errToken,
// 		})
// 	}

// 	userResponse := dto.UserResponseDTO{
// 		Name:  user.Name,
// 		Email: user.Email,
// 		Token: token,
// 	}

// 	return c.JSON(200, gin.Map{
// 		"message": "success login",
// 		"user":    userResponse,
// 	})
// }
