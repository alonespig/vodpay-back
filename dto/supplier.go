package dto

import "time"

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

type BaseModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Spec struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Sku struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Brand struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type SupplierRecharge struct {
	ID            int        `json:"id"`
	SupplierName  string     `json:"supplierName"`
	SupplierCode  string     `json:"supplierCode"`
	Amount        int        `json:"amount"`
	Status        int        `json:"status"`
	ApplyUserName string     `json:"applyUserName"`
	AuditUserName string     `json:"auditUserName"`
	ImageURL      string     `json:"imageURL"`
	Remark        *string    `json:"remark"`
	PassAt        *time.Time `json:"passAt"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type ProjectProduct struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	ProjectID int       `json:"projectID"`
	BrandID   int       `json:"brandID"`
	SpecID    int       `json:"specID"`
	SKUID     int       `json:"skuID"`
	FacePrice int       `json:"facePrice"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	Version   int       `json:"version"`
}

type ProjectProductItem struct {
	ProjectProduct      `json:"product"`
	SupplierProductList []SupplierProduct `json:"supplierProductList"`
}

type ProjectProductResp struct {
	ProjectName string               `json:"projectName"`
	ProductList []ProjectProductItem `json:"productList"`
}
