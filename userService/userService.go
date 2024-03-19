package userService

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/userRepository"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo *userRepository.UserRepository
}

func NewUserService(userRepository *userRepository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepository,
	}
}

// CreateUser implements UserService
func (u *UserService) CreateUser(userRequest model.UserRequest) (model.UserResponse, error) {
	// hash password
	hashedPassword, err := helper.HashPassword(userRequest.Password)
	userRequest.Password = hashedPassword
	if err != nil {
		return model.UserResponse{}, err
	}

	// call repository to save user
	createdUser, err := u.userRepo.CreateUser(userRequest)
	if err != nil {
		return model.UserResponse{}, err
	}

	var userResponse model.UserResponse
	err = copier.Copy(&userResponse, &createdUser)
	if err != nil {
		return model.UserResponse{}, err
	}

	return userResponse, nil
}

// LoginUser implements UserService
// func (u *UserService) LoginUser(user model.User) (model.User, error) {
// 	// call repository to get user
// 	userLogin, err := u.userRepo.LoginUser(user)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return userLogin, nil
// }
