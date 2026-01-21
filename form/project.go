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
