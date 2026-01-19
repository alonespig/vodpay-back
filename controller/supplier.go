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
		Name:         supplier.Name,
		SupplierCode: supplier.Code,
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}
