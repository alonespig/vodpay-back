package controller

import (
	"log"
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
}

func (*ChannelController) CreateChannel(c *gin.Context) {
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

func (*ChannelController) GetChannelList(c *gin.Context) {
	list, err := service.GetChannelList()
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, list)
}

func (*ChannelController) UpdateChannel(c *gin.Context) {
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

func (*ChannelController) CreateChannelSupplierProduct(c *gin.Context) {
	var req form.CreateChannelSupplierProductForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.CreateChannelProjectProduct(req.ProjectProductID, req.SupplierProductID); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func (*ChannelController) GetChannelSupplierProductList(c *gin.Context) {
	list, err := service.GetChannelSupplierProductList()
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, list)
}

func (*ChannelController) GetProjectList(c *gin.Context) {
	var req form.ProjectQueryForm
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	list, err := service.GetProjectList(&req)
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, list)
}

func (*ChannelController) CreateProject(c *gin.Context) {
	var req form.CreateProjectForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.CreateProject(&req); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func (*ChannelController) UpdateProject(c *gin.Context) {
	var req form.UpdateProjectStatusForm
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

func (*ChannelController) GetProjectProductList(c *gin.Context) {
	var req form.ProjectProductQueryForm
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	list, err := service.GetProjectProductList(&req)
	if err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, list)
}

func (*ChannelController) UpdateProjectProduct(c *gin.Context) {
	var req form.UpdateProjectProductForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if err := service.UpdateProjectProduct(&req); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}

func (*ChannelController) CreateProjectProduct(c *gin.Context) {
	var req form.CreateProjectProductForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}
	log.Println(req)
	if err := service.CreateProjectProduct(&req); err != nil {
		Fail(c, 500, err.Error())
		return
	}
	Success(c, nil)
}
