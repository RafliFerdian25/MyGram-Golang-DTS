package routes

import (
	"MyGram-Golang-DTS/controller/commentController"
	"MyGram-Golang-DTS/controller/photoController"
	"MyGram-Golang-DTS/controller/socialMediaController"
	"MyGram-Golang-DTS/controller/userController"
	"MyGram-Golang-DTS/middleware"
	"MyGram-Golang-DTS/repo/commentRepository"
	"MyGram-Golang-DTS/repo/photoRepository"
	"MyGram-Golang-DTS/repo/socialMediaRepository"
	"MyGram-Golang-DTS/repo/userRepository"
	"MyGram-Golang-DTS/service/commentService"
	"MyGram-Golang-DTS/service/photoService"
	"MyGram-Golang-DTS/service/socialMediaService"
	"MyGram-Golang-DTS/service/userService"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {

	// Repositories
	userRepository := userRepository.NewUserRepository(db)
	photoRepository := photoRepository.NewPhotoRepository(db)
	commentRepository := commentRepository.NewCommentRepository(db)
	socialMediaRepository := socialMediaRepository.NewSocialMediaRepository(db)

	// Services
	userService := userService.NewUserService(userRepository)
	photoService := photoService.NewPhotoService(photoRepository)
	commentService := commentService.NewCommentService(commentRepository, photoRepository)
	socialMediaService := socialMediaService.NewSocialMediaService(socialMediaRepository)

	// Controllers
	userController := userController.NewUserController(userService)
	photoController := photoController.NewPhotoController(photoService)
	commentController := commentController.NewCommentController(commentService)
	socialMediasController := socialMediaController.NewSocialMediaController(socialMediaService)

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

	// Photo Routes
	photos := app.Group("/photos")
	photos.Use(middleware.Authentication())
	{
		photos.POST("/", photoController.CreatePhoto)
		photos.GET("/", photoController.GetAllPhotos)
		photos.GET("/:id", photoController.GetPhotoByID)
		photos.PUT("/:id", photoController.UpdatePhoto)
		photos.DELETE("/:id", photoController.DeletePhoto)
	}

	// Comment Routes
	comments := app.Group("/comments")
	comments.Use(middleware.Authentication())
	{
		comments.POST("/", commentController.CreateComment)
		comments.GET("/", commentController.GetAllComments)
		comments.GET("/:id", commentController.GetCommentByID)
		comments.PUT("/:id", middleware.CommentAuthorization(commentService), commentController.UpdateComment)
		comments.DELETE("/:id", middleware.CommentAuthorization(commentService), commentController.DeleteComment)
	}

	// Social Media Routes
	socialMedias := app.Group("/socialmedias")
	socialMedias.Use(middleware.Authentication())
	{
		socialMedias.POST("/", socialMediasController.CreateSocialMedia)
		socialMedias.GET("/", socialMediasController.GetAllSocialMedias)
		socialMedias.GET("/:id", middleware.SocialMediaAuthorization(socialMediaService), socialMediasController.GetSocialMediaByID)
	}

	return app
}
