package repository

import (
	"time"
	"vodpay/database"
)

type BrandSpecSKU struct {
	ID        int       `gorm:"primaryKey;AutoIncrement"`
	Name      string    `gorm:"column:name"`
	BrandID   int       `gorm:"column:brand_id"`
	SpecID    int       `gorm:"column:spec_id"`
	SKUID     int       `gorm:"column:sku_id"`
	CreatedAt time.Time `gorm:"column:created_at;AutoCreateTime"`
}

func (b *BrandSpecSKU) TableName() string {
	return "brand_spec_skus"
}

func CreateBrandSpecSKU(brandSpecSKU *BrandSpecSKU) (*BrandSpecSKU, error) {
	err := database.DB.Create(brandSpecSKU).Error
	if err != nil {
		return nil, err
	}
	return brandSpecSKU, nil
}

func GetBrandSpecSKUByIDInfo(brandID, specID, skuID int) (*BrandSpecSKU, error) {
	brandSpecSKU := &BrandSpecSKU{}
	err := database.DB.Model(&BrandSpecSKU{}).
		Where("brand_id = ? AND spec_id = ? AND sku_id = ?", brandID, specID, skuID).
		First(brandSpecSKU).Error
	if err != nil {
		return nil, err
	}
	return brandSpecSKU, nil
}

func GetBrandSpecSKUList() ([]*BrandSpecSKU, error) {
	brandSpecSKUs := []*BrandSpecSKU{}
	err := database.DB.Model(&BrandSpecSKU{}).
		Find(&brandSpecSKUs).Error
	if err != nil {
		return nil, err
	}
	return brandSpecSKUs, nil
}

func GetBrandSpecSKUByID(id int) (*BrandSpecSKU, error) {
	brandSpecSKU := &BrandSpecSKU{}
	err := database.DB.Model(&BrandSpecSKU{}).
		Where("id = ?", id).
		First(brandSpecSKU).Error
	if err != nil {
		return nil, err
	}
	return brandSpecSKU, nil
}

func GetBBSIDListByID(brandID, skuID, specID int64) ([]int64, error) {
	idList := make([]int64, 0)

	query := database.DB.Model(&BrandSpecSKU{})

	if brandID > 0 {
		query = query.Where("brand_id = ?", brandID)
	}
	if skuID > 0 {
		query = query.Where("sku_id = ?", skuID)
	}
	if specID > 0 {
		query = query.Where("spec_id = ?", specID)
	}

	err := query.Pluck("id", &idList).Error
	if err != nil {
		return nil, err
	}
	return idList, nil
}
