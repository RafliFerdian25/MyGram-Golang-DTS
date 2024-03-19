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
// func (u *UserRepository) LoginUser(user model.User) (model.User, error) {
// 	var userLogin model.User
// 	err := u.db.Model(&model.User{}).First(&userLogin, "email = ?", user.Email).Error
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	match := helper.CheckPasswordHash(user.Password, userLogin.Password)
// 	if !match {
// 		return model.User{}, errors.New(constantError.ErrorEmailOrPasswordNotMatch)
// 	}
// 	return userLogin, nil
// }
