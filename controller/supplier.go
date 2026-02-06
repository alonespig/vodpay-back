package controller

import (
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

// func SupplierList(c *gin.Context) {
// 	suppliers, err := service.SupplierList()
// 	if err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, suppliers)
// }

// func CreateSupplier(c *gin.Context) {
// 	var supplier form.Supplier
// 	if err := c.ShouldBindJSON(&supplier); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	if err := service.CreateSupplier(&model.Supplier{
// 		Name: supplier.Name,
// 		Code: supplier.Code,
// 	}); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func UpdateSupplier(c *gin.Context) {
// 	var supplier form.SupplierUpdateForm
// 	if err := c.ShouldBindJSON(&supplier); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	if err := service.UpdateSupplier(&model.Supplier{
// 		ID:     supplier.ID,
// 		Name:   supplier.Name,
// 		Status: *supplier.Status,
// 	}); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func UpdateSupplierProduct(c *gin.Context) {
// 	var form form.UpdateSupplierProductForm
// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	if err := service.UpdateSupplierProduct(&model.SupplierProduct{
// 		ID:        form.ID,
// 		Code:      form.Code,
// 		Status:    *form.Status,
// 		Price:     int(form.Price * 100),
// 		FacePrice: int(form.FacePrice * 100),
// 	}); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func CreateSupplierProduct(c *gin.Context) {
// 	var form form.SupplierProduct
// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	if err := service.CreateSupplierProduct(&form); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func RechargeSupplier(c *gin.Context) {
// 	var req form.RechargeSupplierForm
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	if err := service.RechargeSupplier(&req); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func SupplierProductList(c *gin.Context) {
// 	var req form.SupplierProductReq
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	log.Printf("supplier product list by info, req: %v", req)
// 	products, err := service.SupplierProductList(&req)
// 	if err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, products)
// }

// func CreateModel(c *gin.Context) {
// 	var req form.CreateModelForm
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		BadRequest(c, err.Error())
// 		return
// 	}
// 	if err := service.CreateModel(req.Type, req.Name); err != nil {
// 		Fail(c, 500, err.Error())
// 		return
// 	}
// 	Success(c, nil)
// }

// func GetModelList(c *gin.Context) {
// 	modelName := c.Query("type")
// 	if modelName == "brands" || modelName == "specs" || modelName == "skus" {
// 		list, err := service.GetModelList(modelName)
// 		if err != nil {
// 			Fail(c, 500, err.Error())
// 			return
// 		}
// 		Success(c, list)
// 	} else {
// 		BadRequest(c, "invalid model type")
// 	}
// }

// func GetSupplierRechargeHistoryList(c *gin.Context) {
// 	recharges, err := service.GetSupplierRechargeHistoryList()
// 	if err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, recharges)
// }

// func GetSupplierRechargeList(c *gin.Context) {
// 	recharges, err := service.GetSupplierRechargeList(1)
// 	if err != nil {
// 		ServerError(c, err.Error())
// 		return
// 	}
// 	Success(c, recharges)
// }

func UpdateSupplierRecharge(c *gin.Context) {
	var req form.SupplierRecharge
	if err := c.ShouldBindJSON(&req); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.UpdateSupplierRecharge(&req); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}
