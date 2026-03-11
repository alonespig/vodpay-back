package controller

import (
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type BSSController struct {
}

func (c *BSSController) Create(ctx *gin.Context) {
	var req struct {
		Type string `json:"type" binding:"required"`
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ServerError(ctx, err.Error())
		return
	}
	if err := service.CreateBaseBSS(req.Type, req.Name); err != nil {
		ServerError(ctx, err.Error())
		return
	}
	Success(ctx, nil)
}

func (c *BSSController) GetList(ctx *gin.Context) {
	list, err := service.GetList()
	if err != nil {
		ServerError(ctx, err.Error())
		return
	}
	Success(ctx, list)
}

func (c *BSSController) GetBrandList(ctx *gin.Context) {
	list, err := service.GetBrandList()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	Success(ctx, list)
}

func (c *BSSController) GetSkuList(ctx *gin.Context) {
	list, err := service.GetSkuList()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	Success(ctx, list)
}

func (c *BSSController) GetSpecList(ctx *gin.Context) {
	list, err := service.GetSpecList()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	Success(ctx, list)
}
