package socialMediaRepository

import (
	"MyGram-Golang-DTS/model"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

// CreateSocialMedia implements SocialMediaRepository
func (c *SocialMediaRepository) CreateSocialMedia(socialMedia model.SocialMediaRequest) (model.SocialMedia, error) {
	socialMediaModel := model.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
	}
	err := c.db.Create(&socialMediaModel).Error
	if err != nil {
		return model.SocialMedia{}, err
	}
	return socialMediaModel, nil
}

// GetAllSocialMedias implements SocialMediaRepository
func (c *SocialMediaRepository) GetAllSocialMedias(userID uint) ([]model.SocialMediaGetModel, error) {
	var socialMedias []model.SocialMediaGetModel
	err := c.db.Model(&model.SocialMedia{}).Preload("User").Where("user_id = ?", userID).Find(&socialMedias).Error
	if err != nil {
		return []model.SocialMediaGetModel{}, err
	}
	return socialMedias, nil
}

// GetSocialMediaByID implements SocialMediaRepository
// func (c *SocialMediaRepository) GetSocialMediaByID(socialMediaID uint) (model.SocialMediaGetModel, error) {
// 	var socialMedia model.SocialMediaGetModel
// 	err := c.db.Model(&model.SocialMedia{}).Preload("User").Preload("Photo").First(&socialMedia, socialMediaID).Error
// 	if err != nil {
// 		return model.SocialMediaGetModel{}, err
// 	}
// 	return socialMedia, nil
// }

// UpdateSocialMedia implements SocialMediaRepository
// func (c *SocialMediaRepository) UpdateSocialMedia(socialMediaRequest model.SocialMediaUpdateRequest, socialMediaID uint) (model.SocialMedia, error) {
// 	var socialMediaModel model.SocialMedia
// 	err := c.db.First(&socialMediaModel, socialMediaID).Error
// 	if err != nil {
// 		return model.SocialMedia{}, err
// 	}

// 	socialMediaModel.Message = socialMediaRequest.Message

// 	err = c.db.Save(&socialMediaModel).Error
// 	if err != nil {
// 		return model.SocialMedia{}, err
// 	}
// 	return socialMediaModel, nil
// }

// DeleteSocialMedia implements SocialMediaRepository
// func (u *SocialMediaRepository) DeleteSocialMedia(socialMediaID uint) error {
// 	var socialMedia model.SocialMedia
// 	err := u.db.Unscoped().Delete(&socialMedia, socialMediaID).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
