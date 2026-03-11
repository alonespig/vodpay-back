package repository

import "time"

type Order struct {
	ID              int64     `gorm:"primaryKey"`
	ProductID       int64     `gorm:"column:product_id"`
	ProductName     string    `gorm:"column:product_name;type:varchar(36)"`
	SupplierID      int64     `gorm:"column:supplier_id"`
	SupplierCode    string    `gorm:"column:supplier_code;type:varchar(36)"`
	SupplierName    string    `gorm:"column:supplier_name;type:varchar(36)"`
	SupProductCode  string    `gorm:"column:supplier_product_code;type:varchar(255)"`
	SupplierPrice   int       `gorm:"column:supplier_price"`
	ChannelIP       string    `gorm:"column:channel_ip;type:varchar(36)"`
	ChannelName     string    `gorm:"column:channel_name;type:varchar(36)"`
	ChannelPrice    int       `gorm:"column:channel_price"`
	ChannelID       int       `gorm:"column:channel_id"`
	BrandSpecSKUID  int64     `gorm:"column:brand_spec_sku_id"`
	PlatFromOrderNo string    `gorm:"column:plat_from_order_no;type:varchar(255)"`
	SelfOrderNo     string    `gorm:"column:self_order_no;type:varchar(255)"`
	ChannelOrderNo  string    `gorm:"column:channel_order_no;uniqueIndex;type:varchar(255)"`
	AccountID       string    `gorm:"column:account_id;type:varchar(255)"`
	Msg             string    `gorm:"column:msg;type:varchar(255)"`
	Status          int       `gorm:"column:status"`
	Remark          string    `gorm:"column:remark;type:varchar(255)"`
	CallBack        string    `gorm:"column:call_back;type:varchar(255)"`
	CreatedAt       time.Time `gorm:"column:create_time"`
	UpdatedAt       time.Time `gorm:"column:update_time"`
	RetryID         int64     `gorm:"column:retry_id"`
}
