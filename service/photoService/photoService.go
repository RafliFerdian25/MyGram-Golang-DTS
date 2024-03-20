package photoService

import (
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/photoRepository"
	"errors"

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
func (p *PhotoService) CreatePhoto(photoRequest model.PhotoRequest) (model.PhotoResponse, error) {
	// call repository to save photo
	createdPhoto, err := p.photoRepo.CreatePhoto(photoRequest)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	var photoResponse model.PhotoResponse
	err = copier.Copy(&photoResponse, &createdPhoto)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return photoResponse, nil
}

// get all photos
func (p *PhotoService) GetAllPhotos() ([]model.PhotoGetResponse, error) {
	// call repository to get all photos
	photos, err := p.photoRepo.GetAllPhotos()
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
func (p *PhotoService) GetPhotoByID(photoID uint) (model.PhotoGetResponse, error) {
	// call repository to get photo by id
	photo, err := p.photoRepo.GetPhotoByID(photoID)
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
func (p *PhotoService) UpdatePhoto(photoRequest model.PhotoRequest, photoID uint, userID uint) (model.PhotoResponse, error) {
	// check if photo belongs to user
	err := p.CheckPhotoOwner(photoID, userID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	// call repository to update photo
	updatedPhoto, err := p.photoRepo.UpdatePhoto(photoRequest, photoID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	var photoResponse model.PhotoResponse
	err = copier.Copy(&photoResponse, &updatedPhoto)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return photoResponse, nil
}

// delete photo
func (p *PhotoService) DeletePhoto(photoID uint, userID uint) error {
	// check if photo belongs to user
	err := p.CheckPhotoOwner(photoID, userID)
	if err != nil {
		return err
	}

	// call repository to delete photo
	err = p.photoRepo.DeletePhoto(photoID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PhotoService) CheckPhotoOwner(photoID uint, userID uint) error {
	photo, err := p.photoRepo.GetPhotoByID(photoID)
	if err != nil {
		return err
	}
	if photo.UserID != userID {
		return errors.New("photo not belongs to user")
	}
	return nil
}
