package router

import (
	"vodpay/controller"
	"vodpay/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// CORS中间件
	r.Use(middleware.CORS())

	// 静态文件服务
	r.Static("/uploads", "./uploads")

	// 上传文件
	r.POST("/upload", controller.Upload)

	// 供应商API端点
	r.GET("/supplier", controller.SupplierList)
	r.POST("/supplier", controller.CreateSupplier)
	r.PUT("/supplier", controller.UpdateSupplier)
	r.GET("/supplier/product", controller.SupplierProductList)
	r.POST("/supplier/product", controller.CreateSupplierProduct)
	r.PUT("/supplier/product", controller.UpdateSupplierProduct)
	r.POST("/supplier/recharge", controller.RechargeSupplier)
	r.GET("/supplier/recharge", controller.GetSupplierRechargeList)
	r.PUT("/supplier/recharge", controller.UpdateSupplierRecharge)
	r.GET("/supplier/recharge/history", controller.GetSupplierRechargeHistoryList)

	// 品牌、规格、SKU API端点
	r.POST("/project", controller.CreateModel)
	r.GET("/project", controller.GetModelList)

	// 渠道API端点
	channel := r.Group("/channel")
	{
		// 创建渠道
		channel.POST("", controller.CreateChannel)
		// 渠道列表
		channel.GET("", controller.GetChannelList)
		// 更新渠道
		channel.PUT("", controller.UpdateChannel)
		project := channel.Group("/:channelID/project")
		{
			// 项目列表
			project.GET("", controller.GetProjectListByChannelID)
			// 创建项目
			project.POST("", controller.CreateProject)
			// 更新项目状态
			project.PUT("/:projectID", controller.UpdateProjectStatus)

			product := project.Group("/:projectID/product")
			{
				// // 产品列表
				product.GET("", controller.GetProductListByProjectID)
				// // 创建产品
				product.POST("", controller.CreateProjectProduct)
				// // 更新产品状态
				// product.PUT("/:productID", controller.UpdateProductStatus)
				// 更新项目产品
				product.PUT("/:productID", controller.UpdateProjectProduct)
			}
		}

	}
	return r
}
