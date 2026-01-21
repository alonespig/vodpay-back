package controller

import (
	"strconv"
	"vodpay/form"
	"vodpay/model"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	// 从URL参数中获取channelID
	channelIDStr := c.Param("channelID")
	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil || channelID <= 0 {
		BadRequest(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.CreateProject(channelID, &req.Name); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func GetProjectListByChannelID(c *gin.Context) {
	// 从URL参数中获取channelID
	channelIDStr := c.Param("channelID")
	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil || channelID <= 0 {
		BadRequest(c, err.Error())
		return
	}
	projects, channel, err := service.GetProjectListByChannelID(channelID)
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, gin.H{
		"channel":  channel,
		"projects": projects,
	})
}

func UpdateProjectStatus(c *gin.Context) {
	var req struct {
		ID     int  `json:"id" binding:"required"`
		Status *int `json:"status" binding:"required,oneof=0 1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.UpdateProjectStatus(req.ID, *req.Status); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func CreateProjectProduct(c *gin.Context) {
	var req form.ProjectProductForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}

	channelIDStr := c.Param("channelID")
	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil || channelID <= 0 {
		BadRequest(c, err.Error())
		return
	}

	projectIDStr := c.Param("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil || projectID <= 0 {
		BadRequest(c, err.Error())
		return
	}

	projectProduct := &model.ProjectProduct{
		ProjectID: projectID,
		BrandID:   req.BrandID,
		SpecID:    req.SpecID,
		SKUID:     req.SKUID,
		FacePrice: int(req.FacePrice * 100),
		Price:     int(req.Price * 100),
	}

	if err := service.CreateProjectProduct(projectProduct); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func GetProductListByProjectID(c *gin.Context) {
	// 从URL参数中获取projectID
	projectIDStr := c.Param("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil || projectID <= 0 {
		BadRequest(c, err.Error())
		return
	}
	products, err := service.GetProjectProductListByProjectID(projectID)
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, products)
}
