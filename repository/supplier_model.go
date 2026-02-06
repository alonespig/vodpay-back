package repository

import "time"

type Supplier struct {
	ID        int       `gorm:"primaryKey;AutoIncrement"`
	Name      string    `gorm:"column:name"`
	Code      string    `gorm:"column:code"`
	Balance   int       `gorm:"column:balance"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Supplier) TableName() string {
	return "suppliers"
}

type SupplierProduct struct {
	ID           int       `gorm:"primaryKey;AutoIncrement"`
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
	CreatedAt    time.Time `gorm:"column:created_at;AutoCreateTime"`
}

func (SupplierProduct) TableName() string {
	return "supplier_products"
}

type BaseModel struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at;AutoCreateTime"`
}

type Sku struct {
	BaseModel
}

func (Sku) TableName() string {
	return "skus"
}

type Spec struct {
	BaseModel
}

func (Spec) TableName() string {
	return "specs"
}

type Brand struct {
	BaseModel
}

func (Brand) TableName() string {
	return "brands"
}

type SupplierRecharge struct {
	ID            int        `gorm:"primaryKey;AutoIncrement"`
	SupplierID    int        `gorm:"column:supplier_id"`
	SupplierName  string     `gorm:"column:supplier_name"`
	SupplierCode  string     `gorm:"column:supplier_code"`
	Amount        int        `gorm:"column:amount"`
	Status        int        `gorm:"column:status"`
	ApplyUserID   int        `gorm:"column:apply_user_id"`
	ApplyUserName string     `gorm:"column:apply_user_name"`
	AuditUserID   int        `gorm:"column:audit_user_id"`
	AuditUserName string     `gorm:"column:audit_user_name"`
	ImageURL      string     `gorm:"column:image_url"`
	Remark        *string    `gorm:"column:remark"`
	PassAt        *time.Time `gorm:"column:pass_at"`
	CreatedAt     time.Time  `gorm:"column:created_at;AutoCreateTime"`
}

func (SupplierRecharge) TableName() string {
	return "supplier_recharges"
}
