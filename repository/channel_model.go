package repository

import "time"

type Channel struct {
	ID            int       `gorm:"primaryKey;AutoIncrement"`
	Name          string    `gorm:"column:name"`
	AppID         string    `gorm:"column:app_id"`
	SecretKey     string    `gorm:"column:secret_key"`
	WhiteList     string    `gorm:"column:white_list"`
	Status        int       `gorm:"column:status"`
	Balance       int       `gorm:"column:balance"`
	CreditLimit   int       `gorm:"column:credit_limit"`
	CreditBalance int       `gorm:"column:credit_balance"`
	CreatedAt     time.Time `gorm:"column:created_at;AutoCreateTime"`
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

type ProjectProduct struct {
	ID        int       `gorm:"primaryKey;AutoIncrement"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`
	ProjectID int       `gorm:"column:project_id"`
	BrandID   int       `gorm:"column:brand_id"`
	SpecID    int       `gorm:"column:spec_id"`
	SKUID     int       `gorm:"column:sku_id"`
	FacePrice int       `gorm:"column:face_price"`
	Price     int       `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at;AutoCreateTime"`
	Version   int       `gorm:"column:version"`
}

func (p *ProjectProduct) TableName() string {
	return "project_products"
}
