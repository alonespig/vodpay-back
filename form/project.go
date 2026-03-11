package form

type ProjectForm struct {
	ChannelID int     `json:"channelID" binding:"required"`
	Status    *int    `json:"status" binding:"required,oneof=0 1"`
	Name      *string `json:"name" binding:"required"`
}

type ProjectProductForm struct {
	BrandID   int     `json:"brandID" binding:"required"`
	SpecID    int     `json:"specID" binding:"required"`
	SKUID     int     `json:"skuID" binding:"required"`
	FacePrice float64 `json:"facePrice" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type UpdateProjectProductForm struct {
	ID        int     `json:"id" binding:"required"`
	Status    *int    `json:"status" binding:"required,oneof=0 1"`
	FacePrice float64 `json:"facePrice" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type AddSupplierProductForm struct {
	ProjectProductID      int   `json:"projectProductID" binding:"required"`
	SupplierProductIDList []int `json:"supplierProductIDList" binding:"required"`
}

type CreateProductForm struct {
	SKUID                 int     `json:"skuID" binding:"required"`
	BrandID               int     `json:"brandID" binding:"required"`
	SpecID                int     `json:"specID" binding:"required"`
	ProjectID             int     `json:"projectID" binding:"required"`
	FacePrice             float64 `json:"facePrice" binding:"required"`
	Price                 float64 `json:"price" binding:"required"`
	LimitNum              int     `json:"limitNum"`
	Model                 *int    `json:"model" binding:"required,oneof=0 1 2"`
	SupplierProductIDList []int   `json:"supplierProductIDList"`
}

type ChannelSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectListResp struct {
	ProjectList []Project `json:"projectList"`
}

type BSSListResp struct {
	BrandSpecSKUList []BaseModel `json:"brandSpecSKUList"`
}
