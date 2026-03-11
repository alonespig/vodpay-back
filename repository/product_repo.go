package repository

import (
	"errors"
	"log"
	"vodpay/common"
	"vodpay/database"

	"gorm.io/gorm"
)

func GetProductListByProjectID(projectID int) ([]Product, error) {
	var products []Product
	err := database.DB.Where("project_id = ?", projectID).Find(&products).Error
	if err != nil {
		log.Printf("[GetProductListByProjectID] projectID = %d: %v", projectID, err)
		return nil, err
	}
	return products, nil
}
func UpdateProductRelationStatus(id, status int) error {
	return database.DB.Model(&ProductRelation{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": status,
		}).Error
}

func GetProductRelationByProductID(productID, supplierProductID int) (*ProductRelation, error) {
	var relation ProductRelation
	if err := database.DB.
		Where("channel_product_id = ? AND supplier_product_id = ?", productID, supplierProductID).
		First(&relation).Error; err != nil {
		return nil, err
	}
	return &relation, nil
}

func GetProductRelationByID(id int) (*ProductRelation, error) {
	var relation ProductRelation
	err := database.DB.Where("id = ?", id).First(&relation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrProductRelationNotFound
		}
		log.Printf("[GetProductRelationByID] id = %d: %v", id, err)
		return nil, common.ErrDBQuery
	}
	return &relation, nil
}
