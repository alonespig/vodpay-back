package repository

import (
	"vodpay/database"
	"vodpay/model"
)

type SupplierRepo struct {
}

func CreateSupplier(supplier *model.Supplier) error {
	return database.DB.Create(supplier).Error
}

func GetSupplierByID(id int) (*model.Supplier, error) {
	var supplier model.Supplier
	if err := database.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func GetSupplierList() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	if err := database.DB.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func GetSupplierByCode(code string) (*model.Supplier, error) {
	var supplier model.Supplier
	if err := database.DB.Where("code = ?", code).First(&supplier).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func UpdateSupplier(supplier *model.Supplier) error {
	return database.DB.Save(supplier).Error
}
