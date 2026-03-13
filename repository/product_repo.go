package repository

import (
	"log"
	"vodpay/database"
)

func GetProductListByProjectID(projectID int64) ([]Product, error) {
	var products []Product
	err := database.DB.Where("project_id = ?", projectID).Find(&products).Error
	if err != nil {
		log.Printf("[GetProductListByProjectID] projectID = %d: %v", projectID, err)
		return nil, err
	}
	return products, nil
}

func UpdateProductSupplierStatus(id, status int) error {
	return database.DB.Model(&ProductSupplier{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": status,
		}).Error
}

func GetProductSupplierByProductID(productID, supplierProductID int64) (*ProductSupplier, error) {
	var relation ProductSupplier
	if err := database.DB.
		Where("product_id = ? AND supplier_product_id = ?", productID, supplierProductID).
		First(&relation).Error; err != nil {
		return nil, err
	}
	return &relation, nil
}

func GetProductSupplierByID(id int) (*ProductSupplier, error) {
	var relation ProductSupplier
	err := database.DB.Where("id = ?", id).First(&relation).Error
	if err != nil {
		log.Printf("[GetProductSupplierByID] id = %d: %v", id, err)
		return nil, err
	}
	return &relation, nil
}
