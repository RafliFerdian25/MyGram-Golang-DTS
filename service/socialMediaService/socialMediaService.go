package socialMediaService

import (
	"MyGram-Golang-DTS/model"
	"MyGram-Golang-DTS/repo/socialMediaRepository"
	"errors"

	"github.com/jinzhu/copier"
)

type SocialMediaService struct {
	socialMediaRepo *socialMediaRepository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository *socialMediaRepository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: socialMediaRepository,
	}
}

// CreateSocialMedia implements SocialMediaService
func (c *SocialMediaService) CreateSocialMedia(socialMediaRequest model.SocialMediaRequest) (model.SocialMediaResponse, error) {
	// call repository to save socialMedia
	createdSocialMedia, err := c.socialMediaRepo.CreateSocialMedia(socialMediaRequest)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	// copy data from createdSocialMedia to socialMediaResponse
	var socialMediaResponse model.SocialMediaResponse
	err = copier.Copy(&socialMediaResponse, &createdSocialMedia)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	return socialMediaResponse, nil
}

// get all socialMedias
func (c *SocialMediaService) GetAllSocialMedias(userID uint) ([]model.SocialMediaGetResponse, error) {
	// call repository to get all socialMedias
	socialMedias, err := c.socialMediaRepo.GetAllSocialMedias(userID)
	if err != nil {
		return []model.SocialMediaGetResponse{}, err
	}

	var socialMediaResponses []model.SocialMediaGetResponse
	err = copier.Copy(&socialMediaResponses, &socialMedias)
	if err != nil {
		return []model.SocialMediaGetResponse{}, err
	}

	return socialMediaResponses, nil
}

// get socialMedia by id
func (c *SocialMediaService) GetSocialMediaByID(socialMediaID uint) (model.SocialMediaGetResponse, error) {
	// call repository to get socialMedia by id
	socialMedia, err := c.socialMediaRepo.GetSocialMediaByID(socialMediaID)
	if err != nil {
		return model.SocialMediaGetResponse{}, err
	}

	var socialMediaResponse model.SocialMediaGetResponse
	err = copier.Copy(&socialMediaResponse, &socialMedia)
	if err != nil {
		return model.SocialMediaGetResponse{}, err
	}

	return socialMediaResponse, nil
}

// update socialMedia
func (c *SocialMediaService) UpdateSocialMedia(socialMediaRequest model.SocialMediaRequest, socialMediaID uint) (model.SocialMediaResponse, error) {
	// call repository to update socialMedia
	updatedSocialMedia, err := c.socialMediaRepo.UpdateSocialMedia(socialMediaRequest, socialMediaID)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	var socialMediaResponse model.SocialMediaResponse
	err = copier.Copy(&socialMediaResponse, &updatedSocialMedia)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	return socialMediaResponse, nil
}

// delete socialMedia
// func (c *SocialMediaService) DeleteSocialMedia(socialMediaID uint, userID uint) error {
// 	// check if socialMedia belongs to user
// 	err := c.CheckSocialMediaOwner(socialMediaID, userID)
// 	if err != nil {
// 		return err
// 	}

// 	// call repository to delete socialMedia
// 	err = c.socialMediaRepo.DeleteSocialMedia(socialMediaID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (c *SocialMediaService) CheckSocialMediaOwner(socialMediaID uint, userID uint) error {
	socialMedia, err := c.socialMediaRepo.GetSocialMediaByID(socialMediaID)
	if err != nil {
		return err
	}
	if socialMedia.UserID != userID {
		return errors.New("social media not belongs to user")
	}
	return nil
}
