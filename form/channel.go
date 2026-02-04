package form

type CreateChannelForm struct {
	Name        string `json:"name" binding:"required"`
	WhiteList   string `json:"whiteList" binding:"required"`
	CreditLimit int    `json:"creditLimit" binding:"required,gt=0"`
}

type UpdateChannelForm struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	WhiteList   string `json:"whiteList" binding:"required"`
	Status      *int   `json:"status" binding:"required,oneof=0 1"`
	CreditLimit int    `json:"creditLimit" binding:"required,gt=0"`
}

type CreateChannelSupplierProductForm struct {
	ProjectProductID  int   `json:"projectProductID" binding:"required"`
	SupplierProductID []int `json:"supplierProductIDList" binding:"required"`
}
