package routes

import (
	userController "MyGram-Golang-DTS/controller/UserController"
	"MyGram-Golang-DTS/repo/userRepository"
	"MyGram-Golang-DTS/userService"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {

	// Repositories
	userRepository := userRepository.NewUserRepository(db)

	// Services
	userService := userService.NewUserService(userRepository)

	// Controllers
	userController := userController.NewUserController(userService)

	app := gin.Default()

	/*
		API Routes
	*/
	// config := middleware.JWTConfig{
	// 	Claims:     &helper.JWTCustomClaims{},
	// 	SigningKey: []byte(cfg.TOKEN_SECRET),
	// }

	// User Routes
	app.POST("/users/register", userController.CreateUser)
	// app.POST("/login", userController.LoginUser)

	return app
}
