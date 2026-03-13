package repository

import (
	"errors"
	"log"
	"vodpay/common"
	"vodpay/database"

	"gorm.io/gorm"
)

func CreateChannel(channel *Channel) error {
	err := database.DB.Create(channel).Error
	if err != nil {
		log.Printf("[CreateChannel]: %v", err)
		return common.ErrDBQuery
	}
	return nil
}

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

func GetChannelList() ([]Channel, error) {
	var channels []Channel
	err := database.DB.Order("created_at DESC").Find(&channels).Error
	if err != nil {
		log.Printf("[GetChannelList]: %v", err)
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
		log.Printf("[CreateProject]: %v", err)
		return common.ErrDBInsert
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

func GetProjectByID(id int64) (*Project, error) {
	var project Project
	err := database.DB.Where("id = ?", id).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrProjectNotFound
		}
		log.Printf("[GetProjectByID] id = %d: %v", id, err)
		return nil, common.ErrDBQuery
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

func CreateProduct(product *Product) (int64, error) {
	err := database.DB.Create(product).Error
	if err != nil {
		log.Printf("[CreateProduct]: %v", err)
		return 0, common.ErrDBInsert
	}
	return product.ID, nil
}

func GetProductCount(projectID, brandSpecSKUID int) (int, error) {
	var count int64
	err := database.DB.Model(&Product{}).
		Where("project_id = ? AND brand_spec_sku_id = ?", projectID, brandSpecSKUID).
		Count(&count).Error
	if err != nil {
		log.Printf("[GetProductCount]: %v", err)
		return 0, err
	}
	return int(count), nil
}

func CreateProductSupplier(relations []ProductSupplier) error {
	return database.DB.Create(&relations).Error
}

func ChangeProductSupplier(id int, supplier *Supplier, product *SupplierProduct) error {
	log.Printf("[ChangeProductSupplier] id = %d supplier = %v product = %v", id, supplier, product)
	return database.DB.Model(&Product{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"supplier_id":           supplier.ID,
			"supplier_name":         supplier.Name,
			"supplier_product_code": product.Code,
		}).Error
}

func CreateProductSupplierList(relations []ProductSupplier) error {
	return database.DB.Create(&relations).Error
}

func GetProductByID(id int64) (*Product, error) {
	var product Product
	err := database.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		log.Printf("[GetProductByID] id = %d: %v", id, err)
		return nil, err
	}
	return &product, nil
}

func GetChannelNameByID(id int64) (string, error) {
	var channel Channel
	err := database.DB.Where("id = ?", id).First(&channel).Error
	if err != nil {
		return "", err
	}
	return channel.Name, nil
}

func GetProjectNameByID(id int64) (string, error) {
	var project Project
	err := database.DB.Where("id = ?", id).First(&project).Error
	if err != nil {
		return "", err
	}
	return project.Name, nil
}

func UpdateProduct(product *Product) error {
	return database.DB.Updates(map[string]interface{}{
		"status":     product.Status,
		"face_price": product.FacePrice,
		"price":      product.Price,
	}).Error
}

func UpdateProduct2(product *Product) error {
	return database.DB.Model(&Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"status":                product.Status,
			"face_price":            product.FacePrice,
			"price":                 product.Price,
			"model":                 product.Model,
			"supplier_id":           product.SupplierID,
			"supplier_name":         product.SupplierName,
			"supplier_product_code": product.SupplierProductCode,
			"supplier_product_id":   product.SupplierProductID,
		}).Error
}

type ProjectProductQuery struct {
	ChannelID *int
	ProjectID *int
	SKUID     *int
	BrandID   *int
	SpecID    *int
}

func GetProductList(q *ProjectProductQuery) ([]Product, error) {
	var products []Product
	if q == nil {
		q = &ProjectProductQuery{}
	}

	query := database.DB.Model(&Product{})

	if q.ChannelID != nil {
		query = query.Where("channel_id = ?", *q.ChannelID)
	}
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
		log.Printf("[GetProductList] projectID = %d skuID = %d brandID = %d specID = %d: %v", *q.ProjectID, *q.SKUID, *q.BrandID, *q.SpecID, err)
		return nil, common.ErrDBQuery
	}
	return products, nil
}

func GetProductSupplierList(productID int64) ([]ProductSupplier, error) {
	var relations []ProductSupplier
	err := database.DB.Where("product_id = ?", productID).Find(&relations).Error
	if err != nil {
		log.Printf("[GetProductSupplierList] productID = %d: %v", productID, err)
		return nil, common.ErrDBQuery
	}
	return relations, nil
}

// func GetProductSupplierList(productID int) ([]ProductRelation, error) {
// 	var relations []ProductRelation
// 	err := database.DB.Where("channel_product_id = ?", productID).Find(&relations).Error
// 	if err != nil {
// 		log.Printf("[GetProductSupplierList] productID = %d: %v", productID, err)
// 		return nil, err
// 	}
// 	return relations, nil
// }
