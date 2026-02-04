package repository

import "time"

type Supplier struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Code      string    `gorm:"column:code"`
	Balance   int       `gorm:"column:balance"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Supplier) TableName() string {
	return "vodpay_supplier"
}

type SupplierProduct struct {
	ID           int       `gorm:"primaryKey"`
	Name         string    `gorm:"column:name"`
	Code         string    `gorm:"column:code"`
	SupplierID   int       `gorm:"column:supplier_id"`
	SupplierName string    `gorm:"column:supplier_name"`
	SupplierCode string    `gorm:"column:supplier_code"`
	FacePrice    int       `gorm:"column:face_price"`
	SpecID       int       `gorm:"column:spec_id"`
	SKUID        int       `gorm:"column:sku_id"`
	BrandID      int       `gorm:"column:brand_id"`
	Price        int       `gorm:"column:price"`
	Status       int       `gorm:"column:status"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

func (SupplierProduct) TableName() string {
	return "vodpay_supplier_product"
}

type BaseModel struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type Sku struct {
	BaseModel
}

func (Sku) TableName() string {
	return "vodpay_sku"
}

type Spec struct {
	BaseModel
}

func (Spec) TableName() string {
	return "vodpay_spec"
}

type Brand struct {
	BaseModel
}

func (Brand) TableName() string {
	return "vodpay_brand"
}
