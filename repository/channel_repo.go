package repository

import (
	"errors"
	"vodpay/database"

	"gorm.io/gorm"
)

func CreateChannel(channel *Channel) error {
	err := database.DB.Create(channel).Error
	if err != nil {
		return err
	}
	return nil
}

func GetChannelByID(id int) (*Channel, error) {
	var channel Channel
	err := database.DB.Where("id = ?", id).First(&channel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrChannelNotFound
		}
		return nil, err
	}
	return &channel, nil
}

func GetChannelByAppID(appID string) (*Channel, error) {
	var channel Channel
	err := database.DB.Where("app_id = ?", appID).First(&channel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrChannelNotFound
		}
		return nil, err
	}
	return &channel, nil
}

func GetChannelList() ([]Channel, error) {
	var channels []Channel
	err := database.DB.Order("created_at DESC").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func UpdateChannel(channel *Channel) error {
	return database.DB.Updates(channel).Error
}

func CreateProject(project *Project) error {
	err := database.DB.Create(project).Error
	if err != nil {
		return err
	}
	return nil
}

type ProjectQuery struct {
	ChannelID *int
}

func GetProjectList(q *ProjectQuery) ([]Project, error) {
	var projects []Project

	if q == nil {
		q = &ProjectQuery{}
	}

	query := database.DB.Model(&Project{})

	if q.ChannelID != nil {
		query = query.Where("channel_id = ?", *q.ChannelID)
	}
	err := query.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectByID(id int) (*Project, error) {
	var project Project
	err := database.DB.Where("id = ?", id).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}
	return &project, nil
}

func UpdateProject(project *Project) error {
	return database.DB.Model(&Project{}).
		Where("id = ?", project.ID).
		Updates(map[string]interface{}{
			"status": project.Status,
		}).Error
}

func CreateProjectProduct(product *ProjectProduct) (int, error) {
	err := database.DB.Create(product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func CreateProductRelation(relations []ProductRelation) error {
	return database.DB.Create(&relations).Error
}

func GetProjectProductByID(id int) (*ProjectProduct, error) {
	var product ProjectProduct
	err := database.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProjectProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func UpdateProjectProduct(product *ProjectProduct) error {
	return database.DB.Updates(map[string]interface{}{
		"status":     product.Status,
		"face_price": product.FacePrice,
		"price":      product.Price,
	}).Error
}

type ProjectProductQuery struct {
	ProjectID *int
	SKUID     *int
	BrandID   *int
	SpecID    *int
}

func GetProjectProductList(q *ProjectProductQuery) ([]ProjectProduct, error) {
	var products []ProjectProduct
	if q == nil {
		q = &ProjectProductQuery{}
	}

	query := database.DB.Model(&ProjectProduct{})

	if q.ProjectID != nil {
		query = query.Where("project_id = ?", *q.ProjectID)
	}
	if q.SKUID != nil {
		query = query.Where("sku_id = ?", *q.SKUID)
	}
	if q.BrandID != nil {
		query = query.Where("brand_id = ?", *q.BrandID)
	}
	if q.SpecID != nil {
		query = query.Where("spec_id = ?", *q.SpecID)
	}
	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductRelationList(productID int) ([]ProductRelation, error) {
	var relations []ProductRelation
	err := database.DB.Where("channel_product_id = ?", productID).Find(&relations).Error
	if err != nil {
		return nil, err
	}
	return relations, nil
}
