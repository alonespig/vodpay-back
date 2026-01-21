package controller

import (
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

func CreateChannel(c *gin.Context) {
	var req form.CreateChannelForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.CreateChannel(&req); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func GetChannelList(c *gin.Context) {
	list, err := service.GetChannelList()
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, list)
}

func UpdateChannel(c *gin.Context) {
	var req form.UpdateChannelForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.UpdateChannel(&req); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}
