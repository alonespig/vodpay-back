package controller

import (
	"strconv"
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (*ProductController) UpdateProductRelation(ctx *gin.Context) {
	var form form.UpdateProductRelationForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	if err := service.UpdateProductRelation(&form); err != nil {
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	Success(ctx, nil)
}

func (*ProductController) UpdateProduct(ctx *gin.Context) {
	var form form.UpdateProductForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	if err := service.UpdateProduct(&form); err != nil {
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	Success(ctx, nil)
}

// 获取产品列表
// @Summary 获取产品列表
// @Description 获取产品列表
// @Tags 产品
// @Accept json
// @Produce json
// @Param channelID query int false "渠道ID"
// @Param projectID query int true "项目ID"
// @Param page query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} form.ProductListQueryResp
// @Failure 400 {object} form.ErrorResp
// @Failure 500 {object} form.ErrorResp
// @Router /product [get]
func (*ProductController) GetProductList(ctx *gin.Context) {
	var form form.ProductListQueryForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	resp, err := service.GetProductList(&form)
	if err != nil {
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	Success(ctx, resp)
}

// 获取产品的供应商产品列表
func (*ProductController) GetProductSupplierList(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		Fail(ctx, CodeParamError, err.Error())
		return
	}

	list, err := service.GetProductSupplierList(productID)
	if err != nil {
		Fail(ctx, 500, err.Error())
		return
	}
	Success(ctx, list)
}

func (*ProductController) CreateProduct(ctx *gin.Context) {
	var form form.CreateProductForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	if err := service.CreateProduct(&form); err != nil {
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	Success(ctx, nil)
}
