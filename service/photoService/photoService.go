package photoService

import (
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/photoRepository"

	"github.com/jinzhu/copier"
)

type PhotoService struct {
	photoRepo *photoRepository.PhotoRepository
}

func NewPhotoService(photoRepository *photoRepository.PhotoRepository) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepository,
	}
}

// CreatePhoto implements PhotoService
func (u *PhotoService) CreatePhoto(photoRequest model.PhotoCreateRequest) (model.PhotoCreateResponse, error) {
	// call repository to save photo
	createdPhoto, err := u.photoRepo.CreatePhoto(photoRequest)
	if err != nil {
		return model.PhotoCreateResponse{}, err
	}

	var photoResponse model.PhotoCreateResponse
	err = copier.Copy(&photoResponse, &createdPhoto)
	if err != nil {
		return model.PhotoCreateResponse{}, err
	}

	return photoResponse, nil
}

// get all photos
func (u *PhotoService) GetAllPhotos() ([]model.PhotoGetResponse, error) {
	// call repository to get all photos
	photos, err := u.photoRepo.GetAllPhotos()
	if err != nil {
		return []model.PhotoGetResponse{}, err
	}

	var photoResponses []model.PhotoGetResponse
	err = copier.Copy(&photoResponses, &photos)
	if err != nil {
		return []model.PhotoGetResponse{}, err
	}

	return photoResponses, nil
}

// get photo by id
func (u *PhotoService) GetPhotoByID(photoID uint) (model.PhotoGetResponse, error) {
	// call repository to get photo by id
	photo, err := u.photoRepo.GetPhotoByID(photoID)
	if err != nil {
		return model.PhotoGetResponse{}, err
	}

	var photoResponse model.PhotoGetResponse
	err = copier.Copy(&photoResponse, &photo)
	if err != nil {
		return model.PhotoGetResponse{}, err
	}

	return photoResponse, nil
}

// update photo
// func (u *PhotoService) UpdatePhoto(photoRequest model.PhotoUpdateRequest, photoID uint) (model.PhotoResponse, error) {
// 	// call repository to update photo
// 	updatedPhoto, err := u.photoRepo.UpdatePhoto(photoRequest, photoID)
// 	if err != nil {
// 		return model.PhotoResponse{}, err
// 	}

// 	var photoResponse model.PhotoResponse
// 	err = copier.Copy(&photoResponse, &updatedPhoto)
// 	if err != nil {
// 		return model.PhotoResponse{}, err
// 	}

// 	return photoResponse, nil
// }

// delete photo
// func (u *PhotoService) DeletePhoto(photoID uint) error {
// 	// call repository to delete photo
// 	err := u.photoRepo.DeletePhoto(photoID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
