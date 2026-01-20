package controller

import (
	"vodpay/form"
	"vodpay/model"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

func SupplierList(c *gin.Context) {
	suppliers, err := service.SupplierList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, suppliers)
}

func CreateSupplier(c *gin.Context) {
	var supplier form.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.CreateSupplier(&model.Supplier{
		Name: supplier.Name,
		Code: supplier.Code,
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func UpdateSupplier(c *gin.Context) {
	var supplier form.SupplierUpdateForm
	if err := c.ShouldBindJSON(&supplier); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.UpdateSupplier(&model.Supplier{
		ID:     supplier.ID,
		Name:   supplier.Name,
		Status: *supplier.Status,
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func UpdateSupplierProduct(c *gin.Context) {
	var form form.UpdateSupplierProductForm
	if err := c.ShouldBindJSON(&form); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.UpdateSupplierProduct(&model.SupplierProduct{
		ID:        form.ID,
		Code:      form.Code,
		Status:    *form.Status,
		Price:     int(form.Price * 100),
		FacePrice: int(form.FacePrice * 100),
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func CreateSupplierProduct(c *gin.Context) {
	var form form.SupplierProduct
	if err := c.ShouldBindJSON(&form); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.CreateSupplierProduct(&form); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func RechargeSupplier(c *gin.Context) {
	var req form.RechargeSupplierForm
	if err := c.ShouldBindJSON(&req); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.RechargeSupplier(&req); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func SupplierProductList(c *gin.Context) {
	products, err := service.SupplierProductList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, products)
}
