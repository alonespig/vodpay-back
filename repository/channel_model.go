package repository

import "time"

type Channel struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	AppID       string    `gorm:"column:app_id" json:"appID"`
	SecretKey   string    `gorm:"column:secret_key" json:"secretKey"`
	WhiteList   string    `gorm:"column:white_list" json:"whiteList"`
	Status      int       `gorm:"column:status" json:"status"`
	Balance     int       `gorm:"column:balance" json:"balance"`
	CreditLimit int       `gorm:"column:credit_limit" json:"creditLimit"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (c *Channel) TableName() string {
	return "vodpay_channel"
}

type Project struct {
	ID        int       `gorm:"primaryKey"`
	ChannelID int       `gorm:"column:channel_id"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (p *Project) TableName() string {
	return "vodpay_project"
}

type ProjectProduct struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`
	ProjectID int       `gorm:"column:project_id"`
	BrandID   int       `gorm:"column:brand_id"`
	SpecID    int       `gorm:"column:spec_id"`
	SKUID     int       `gorm:"column:sku_id"`
	FacePrice int       `gorm:"column:face_price"`
	Price     int       `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Version   int       `gorm:"column:version"`
}

func (p *ProjectProduct) TableName() string {
	return "vodpay_project_product"
}
