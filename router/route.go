package routes

import (
	userController "MyGram-Golang-DTS/controller/UserController"
	"MyGram-Golang-DTS/middleware"
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

	// User Routes
	users := app.Group("/users")
	users.POST("/register", userController.CreateUser)
	users.POST("/login", userController.LoginUser)

	users.Use(middleware.Authentication())
	{
		users.Use(middleware.Authentication())
		users.PUT("/", userController.UpdateUser)
		users.DELETE("/", userController.DeleteUser)
	}

	return app
}
