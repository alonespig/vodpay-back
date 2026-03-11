package repository

import "time"

type Channel struct {
	ID          int       `gorm:"primaryKey;AutoIncrement"`
	Name        string    `gorm:"column:name"`
	AppID       string    `gorm:"column:app_id"`
	SecretKey   string    `gorm:"column:secret_key"`
	WhiteList   string    `gorm:"column:white_list"`
	Status      int       `gorm:"column:status"`
	Balance     int       `gorm:"column:balance"`
	CreditLimit int       `gorm:"column:credit_limit"`
	CreatedAt   time.Time `gorm:"column:created_at;AutoCreateTime"`
}

func (c *Channel) TableName() string {
	return "channels"
}

type Project struct {
	ID        int       `gorm:"primaryKey;AutoIncrement"`
	ChannelID int       `gorm:"column:channel_id"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;AutoCreateTime"`
}

func (p *Project) TableName() string {
	return "projects"
}

type Product struct {
	ID                  int       `gorm:"primaryKey;AutoIncrement"`
	Name                string    `gorm:"column:name"`
	Status              int       `gorm:"column:status"`
	ChannelID           int       `gorm:"column:channel_id"`
	ProjectID           int       `gorm:"column:project_id"`
	LimitCount          int       `gorm:"column:limit_count"`
	SupplierID          int       `gorm:"column:supplier_id"`
	SupplierName        string    `gorm:"column:supplier_name"`
	SupplierProductCode string    `gorm:"column:supplier_product_code"`
	SupplierProductID   int       `gorm:"column:supplier_product_id"`
	BrandSpecSKUID      int       `gorm:"column:brand_spec_sku_id"`
	FacePrice           int       `gorm:"column:face_price"`
	Price               int       `gorm:"column:price"`
	Model               int       `gorm:"column:model"`
	CreatedAt           time.Time `gorm:"column:created_at;AutoCreateTime"`
	Version             int       `gorm:"column:version"`
}

func (p *Product) TableName() string {
	return "project_products"
}
