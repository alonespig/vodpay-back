package controller

import (
	"fmt"
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type RechargeController struct {
}

// RechargeSupplier 供应商充值
func (s *RechargeController) RechargeSupplier(c *gin.Context) {
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
func (s *RechargeController) GetSupplierRechargeList(c *gin.Context) {
	recharges, err := service.GetSupplierRechargeList(1)
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, recharges)
}

// UpdateSupplierRecharge 更新供应商充值
func (s *RechargeController) UpdateSupplierRecharge(c *gin.Context) {
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
func (s *RechargeController) GetSupplierRechargeHistoryList(c *gin.Context) {
	recharges, err := service.GetSupplierRechargeHistoryList()
	if err != nil {
		ServerError(c, err.Error())
		return
	}
	Success(c, recharges)
}

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"msg": "获取文件失败"})
		return
	}

	savePath := "./uploads/" + file.Filename
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(500, gin.H{"msg": "保存失败"})
		return
	}

	// 返回文件访问URL
	fileURL := "/uploads/" + file.Filename
	Success(c, gin.H{
		"msg":      "上传成功",
		"url":      fmt.Sprintf("http://%s%s", c.Request.Host, fileURL),
		"filename": file.Filename,
	})
}
