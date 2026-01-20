package form

type CreateModelForm struct {
	Type string `json:"type" binding:"required,oneof=brands specs skus"`
	Name string `json:"name" binding:"required"`
}
