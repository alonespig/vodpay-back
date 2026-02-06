package controller

import (
	"log"
	"vodpay/form"
	"vodpay/repository"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type SupplierController struct {
}

func (s *SupplierController) SupplierList(c *gin.Context) {
	suppliers, err := service.SupplierList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, suppliers)
}

func (s *SupplierController) CreateSupplier(c *gin.Context) {
	var supplier form.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.CreateSupplier(&repository.Supplier{
		Name: supplier.Name,
		Code: supplier.Code,
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

func (s *SupplierController) UpdateSupplier(c *gin.Context) {
	var supplier form.SupplierUpdateForm
	if err := c.ShouldBindJSON(&supplier); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.UpdateSupplier(&repository.Supplier{
		ID:     supplier.ID,
		Name:   supplier.Name,
		Status: *supplier.Status,
	}); err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, nil)
}

// SupplierProductList 获取供应商产品列表
func (s *SupplierController) SupplierProductList(c *gin.Context) {
	var req form.SupplierProductReq
	if err := c.ShouldBindQuery(&req); err != nil {
		ServerError(c, err.Error())
		return
	}
	log.Printf("supplier product list by info, req: %v", req)
	products, err := service.SupplierProductList(&req)
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, products)
}

// CreateSupplierProduct 创建供应商产品
func (s *SupplierController) CreateSupplierProduct(c *gin.Context) {
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

// UpdateSupplierProduct 更新供应商产品
func (s *SupplierController) UpdateSupplierProduct(c *gin.Context) {
	var form form.UpdateSupplierProductForm
	if err := c.ShouldBindJSON(&form); err != nil {
		ServerError(c, err.Error())
		return
	}
	if err := service.UpdateSupplierProduct(&repository.SupplierProduct{
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

// RechargeSupplier 供应商充值
func (s *SupplierController) RechargeSupplier(c *gin.Context) {
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

// GetSupplierRechargeList 获取供应商充值列表
func (s *SupplierController) GetSupplierRechargeList(c *gin.Context) {
	recharges, err := service.GetSupplierRechargeList(1)
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, recharges)
}

// UpdateSupplierRecharge 更新供应商充值
func (s *SupplierController) UpdateSupplierRecharge(c *gin.Context) {
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

// GetSupplierRechargeHistoryList 获取供应商充值历史列表
func (s *SupplierController) GetSupplierRechargeHistoryList(c *gin.Context) {
	recharges, err := service.GetSupplierRechargeHistoryList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, recharges)
}

func GetSkuList(c *gin.Context) {
	skus, err := service.GetSkuList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, skus)
}
func GetBrandList(c *gin.Context) {
	brands, err := service.GetBrandList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, brands)
}
func GetSpecList(c *gin.Context) {
	specs, err := service.GetSpecList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, specs)
}

func CreateModel(c *gin.Context) {
	var req form.CreateModelForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.CreateModel(req.Type, req.Name); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}
