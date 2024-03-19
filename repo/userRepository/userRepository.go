package userRepository

import (
	"MyGram-Golang-DTS/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser implements UserRepository
func (u *UserRepository) CreateUser(user model.UserRequest) (model.User, error) {
	userModel := model.User{
		Username:        user.Username,
		Email:           user.Email,
		Password:        user.Password,
		Age:             user.Age,
		ProfileImageUrl: &user.ProfileImageUrl,
	}
	err := u.db.Create(&userModel).Error
	if err != nil {
		return model.User{}, err
	}
	return userModel, nil
}

// LoginUser implements UserRepository
func (u *UserRepository) LoginUser(userLogin model.UserLoginRequest) (model.User, error) {
	user := model.User{}
	err := u.db.Where("email = ?", userLogin.Email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
