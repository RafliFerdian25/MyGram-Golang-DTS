package userService

import (
	"MyGram-Golang-DTS/helper"
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/userRepository"
	"errors"
	"fmt"

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
func (u *UserService) LoginUser(userLogin model.UserLoginRequest) (string, error) {
	// call repository to get user
	user, err := u.userRepo.LoginUser(userLogin)
	if err != nil {
		return "", err
	}
	fmt.Println(user)
	fmt.Println(userLogin)

	match := helper.CheckPasswordHash(userLogin.Password, user.Password)
	if !match {
		return "", errors.New("email or password is incorrect")
	}

	token, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// update user
func (u *UserService) UpdateUser(userRequest model.UserUpdateRequest, userID uint) (model.UserResponse, error) {
	// call repository to update user
	updatedUser, err := u.userRepo.UpdateUser(userRequest, userID)
	if err != nil {
		return model.UserResponse{}, err
	}

	var userResponse model.UserResponse
	err = copier.Copy(&userResponse, &updatedUser)
	if err != nil {
		return model.UserResponse{}, err
	}

	return userResponse, nil
}
