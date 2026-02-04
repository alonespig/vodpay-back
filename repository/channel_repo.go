package repository

import "vodpay/database"

// type ChannelRepo struct {
// 	db *gorm.DB
// }

// func NewChannelRepo(db *gorm.DB) *ChannelRepo {
// 	return &ChannelRepo{db: db}
// }

func GetChannelByID(id int) (*Channel, error) {
	var channel Channel
	err := database.DB.Where("id = ?", id).First(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func GetChannelByAppID(appID string) (*Channel, error) {
	var channel Channel
	err := database.DB.Where("app_id = ?", appID).First(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func CreateChannel(channel *Channel) error {
	err := database.DB.Create(channel).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateProject(project *Project) error {
	err := database.DB.Create(project).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateProjectProduct(product *ProjectProduct) error {
	err := database.DB.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}
