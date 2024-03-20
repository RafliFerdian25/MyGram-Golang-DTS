package photoRepository

import (
	"MyGram-Golang-DTS/model"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

// CreatePhoto implements PhotoRepository
func (u *PhotoRepository) CreatePhoto(photo model.PhotoRequest) (model.Photo, error) {
	photoModel := model.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   photo.UserID,
	}
	err := u.db.Create(&photoModel).Error
	if err != nil {
		return model.Photo{}, err
	}
	return photoModel, nil
}

// GetAllPhotos implements PhotoRepository
func (u *PhotoRepository) GetAllPhotos() ([]model.PhotoGetModel, error) {
	var photos []model.PhotoGetModel
	err := u.db.Model(&model.Photo{}).Preload("User").Find(&photos).Error
	if err != nil {
		return []model.PhotoGetModel{}, err
	}
	return photos, nil
}

// GetPhotoByID implements PhotoRepository
func (u *PhotoRepository) GetPhotoByID(photoID uint) (model.PhotoGetModel, error) {
	var photo model.PhotoGetModel
	err := u.db.Model(&model.Photo{}).Preload("User").First(&photo, photoID).Error
	if err != nil {
		return model.PhotoGetModel{}, err
	}
	return photo, nil
}

// UpdatePhoto implements PhotoRepository
func (u *PhotoRepository) UpdatePhoto(photo model.PhotoRequest, photoID uint) (model.Photo, error) {
	var photoModel model.Photo
	err := u.db.First(&photoModel, photoID).Error
	if err != nil {
		return model.Photo{}, err
	}

	photoModel.Title = photo.Title
	photoModel.Caption = photo.Caption
	photoModel.PhotoUrl = photo.PhotoUrl

	err = u.db.Save(&photoModel).Error
	if err != nil {
		return model.Photo{}, err
	}
	return photoModel, nil
}

// DeletePhoto implements PhotoRepository
func (u *PhotoRepository) DeletePhoto(photoID uint) error {
	var photo model.Photo
	err := u.db.Unscoped().Delete(&photo, photoID).Error
	if err != nil {
		return err
	}
	return nil
}
