package form

import "time"

type Supplier struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type CreateModelForm struct {
	Type string `json:"type" binding:"required,oneof=brands specs skus"`
	Name string `json:"name" binding:"required"`
}

type SupplierUpdateForm struct {
	ID     int    `json:"id" binding:"required"`
	Status *int   `json:"status" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type SupplierProductListReq struct {
	BrandSpecSKUID int `form:"id"`
	BrandID        int `form:"brandID"`
	SpecID         int `form:"specID"`
	SKUID          int `form:"skuID"`
	Page           int `form:"page"`
	PageSize       int `form:"pageSize"`
}

type CreateSupplierProductReq struct {
	Code       string `json:"code" binding:"required"`
	SupplierID int    `json:"supplierID" binding:"required"`
	FacePrice  int    `json:"facePrice" binding:"required"`
	SpecID     int    `json:"specID" binding:"required"`
	SKUID      int    `json:"skuID" binding:"required"`
	BrandID    int    `json:"brandID" binding:"required"`
	Price      int    `json:"price" binding:"required"`
}

type SupplierProduct struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	SupplierID   int       `json:"supplierID"`
	SupplierName string    `json:"supplierName"`
	SupplierCode string    `json:"supplierCode"`
	FacePrice    int       `json:"facePrice"`
	SpecID       int       `json:"specID"`
	SKUID        int       `json:"skuID"`
	BrandID      int       `json:"brandID"`
	Price        int       `json:"price"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
}

type SupplierProductResp struct {
	Supplier Supplier          `json:"supplier"`
	Total    int64             `json:"total"`
	Items    []SupplierProduct `json:"items"`
}

type SupplierProductListResp struct {
	Total int64             `json:"total"`
	Items []SupplierProduct `json:"items"`
}

type UpdateSupplierProductForm struct {
	ID        int     `json:"id" binding:"required"`
	Code      string  `json:"code" binding:"required"`
	Status    *int    `json:"status" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	FacePrice float64 `json:"facePrice" binding:"required"`
}

type UpdateProductRelationForm struct {
	ID     int  `json:"id" binding:"required"`
	Status *int `json:"status" binding:"required"`
}

type UpdateProductForm struct {
	ID                int     `json:"productID" binding:"required"`
	FacePrice         float64 `json:"facePrice" binding:"required"`
	SupplierProductID int     `json:"supplierProductID"`
	Status            *int    `json:"status" binding:"required"`
	Model             *int    `json:"model" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
}

type RechargeSupplierForm struct {
	ID       int    `json:"id" binding:"required"`
	Amount   int    `json:"amount" binding:"required,gt=0"`
	Name     string `json:"name" binding:"required"`
	ImageURL string `json:"imageURL" binding:"required"`
}

type SupplierRecharge struct {
	ID         int    `json:"id" binding:"required"`
	SupplierID int    `json:"supplierID" binding:"required"`
	Amount     int    `json:"amount" binding:"required,gt=0"`
	Status     *int   `json:"status" binding:"required,oneof=0 2"`
	Remark     string `json:"remark"`
}

// 供应商
type SupplierResp struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Code      string    `db:"code" json:"code"`
	Balance   int       `db:"balance" json:"balance"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
