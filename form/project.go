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

type CreateProjectProductForm struct {
	SKUID                 int     `json:"skuID" binding:"required"`
	BrandID               int     `json:"brandID" binding:"required"`
	SpecID                int     `json:"specID" binding:"required"`
	ProjectID             int     `json:"projectID" binding:"required"`
	FacePrice             float64 `json:"facePrice" binding:"required"`
	Price                 float64 `json:"price" binding:"required"`
	SupplierProductIDList []int   `json:"supplierProductIDList"`
}
