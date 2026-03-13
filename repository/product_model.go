package repository

import "time"

// type ProductRelation struct {
// 	ID                int       `gorm:"primaryKey;AutoIncrement"`
// 	ChannelProductID  int       `gorm:"column:channel_product_id"`
// 	SupplierProductID int       `gorm:"column:supplier_product_id"`
// 	Status            int       `gorm:"column:status"`
// 	CreatedAt         time.Time `gorm:"column:created_at"`
// }

type ProductSupplier struct {
	ID                int64     `gorm:"primaryKey;AutoIncrement"`
	ProductID         int64     `gorm:"column:product_id"`
	SupplierProductID int64     `gorm:"column:supplier_product_id"`
	Status            int       `gorm:"column:status"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}

func (*ProductSupplier) TableName() string {
	return "channel_supplier_products"
}
