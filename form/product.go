package form

import "time"

type ProductListQueryForm struct {
	ProjectID *int `form:"projectID" binding:"required"`
}

type Product struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Status              int       `json:"status"`
	SupplierName        string    `json:"supplierName,omitempty"`
	SupplierPrice       int       `json:"supplierPrice,omitempty"`
	SupplierProductName string    `json:"supplierProductName,omitempty"`
	SupplierProductID   int       `json:"supplierProductID,omitempty"` //如果没有，不序列化
	FacePrice           int       `json:"facePrice"`
	Price               int       `json:"price"`
	Model               int       `json:"model"`
	CreatedAt           time.Time `json:"createdAt"`
}

type ProductListQueryResp struct {
	ChannelName string    `json:"channelName"`
	ProjectName string    `json:"projectName"`
	ProductList []Product `json:"productList"`
}

type ProductSupplier struct {
	ID                  int    `json:"id"`
	SupplierID          int    `json:"supplierID"`
	SupplierName        string `json:"supplierName"`
	SupplierCode        string `json:"supplierCode"`
	SupplierProductID   int    `json:"supplierProductID"`
	SupplierProductName string `json:"supplierProductName"`
	SupplierProductCode string `json:"supplierProductCode"`
	FacePrice           int    `json:"facePrice"`
	Price               int    `json:"price"`
	Status              int    `json:"status"`
}

type ProductSupplierResp struct {
	ChannelName  string            `json:"channelName"`
	ProjectName  string            `json:"projectName"`
	ProductName  string            `json:"productName"`
	Product      Product           `json:"product"`
	SupplierList []ProductSupplier `json:"supplierList"`
}
