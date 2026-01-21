package form

type Supplier struct {
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

type SupplierProduct struct {
	SupplierID int     `json:"supplierID" binding:"required"`
	Code       string  `json:"code" binding:"required"`
	FacePrice  float64 `json:"facePrice" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	SpecID     int     `json:"specID" binding:"required"`
	SKUID      int     `json:"skuID" binding:"required"`
	BrandID    int     `json:"brandID" binding:"required"`
}

type UpdateSupplierProductForm struct {
	ID        int     `json:"id" binding:"required"`
	Code      string  `json:"code" binding:"required"`
	Status    *int    `json:"status" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	FacePrice float64 `json:"facePrice" binding:"required"`
}

type RechargeSupplierForm struct {
	SupplierID   int    `json:"supplierID" binding:"required"`
	Amount       int    `json:"amount" binding:"required,gt=0"`
	SupplierName string `json:"supplierName" binding:"required"`
}
