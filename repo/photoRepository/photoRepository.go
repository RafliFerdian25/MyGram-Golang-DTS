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
func (u *PhotoRepository) CreatePhoto(photo model.PhotoCreateRequest) (model.Photo, error) {
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

// GetPhotos implements PhotoRepository
func (u *PhotoRepository) GetPhotos() ([]model.PhotoGetModel, error) {
	var photos []model.PhotoGetModel
	err := u.db.Model(&model.Photo{}).Preload("User").Find(&photos).Error
	if err != nil {
		return []model.PhotoGetModel{}, err
	}
	return photos, nil
}

// UpdatePhoto implements PhotoRepository
// func (u *PhotoRepository) UpdatePhoto(photoRequest model.PhotoUpdateRequest, photoID uint) (model.Photo, error) {
// 	var photo model.Photo

// 	// Mencari pengguna dengan photoID yang diberikan
// 	err := u.db.First(&photo, photoID).Error
// 	if err != nil {
// 		return model.Photo{}, err
// 	}

// 	photo.Photoname = photoRequest.Photoname
// 	photo.Email = photoRequest.Email
// 	photo.Age = photoRequest.Age
// 	photo.ProfileImageUrl = photoRequest.ProfileImageUrl

// 	err = u.db.Save(&photo).Error
// 	if err != nil {
// 		return model.Photo{}, err
// 	}
// 	return photo, nil
// }

// DeletePhoto implements PhotoRepository
// func (u *PhotoRepository) DeletePhoto(photoID uint) error {
// 	err := u.db.Unscoped().Delete(&model.Photo{}, photoID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
