package controller

import (
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

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

func GetModelList(c *gin.Context) {
	modelName := c.Query("type")
	if modelName == "brands" || modelName == "specs" || modelName == "skus" {
		list, err := service.GetModelList(modelName)
		if err != nil {
			Fail(c, 500, err.Error())
			return
		}
		Success(c, list)
	} else {
		BadRequest(c, "invalid model type")
	}
}
