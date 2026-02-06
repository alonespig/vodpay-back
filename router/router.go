package router

import (
	"vodpay/controller"
	"vodpay/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// CORS中间件
	r.Use(middleware.CORS())

	// 静态文件服务
	r.Static("/uploads", "./uploads")

	// 上传文件
	r.POST("/upload", controller.Upload)

	r.GET("/sku", controller.GetSkuList)
	r.GET("/spec", controller.GetSpecList)
	r.GET("/brand", controller.GetBrandList)

	// 供应商API端点
	supplier := r.Group("/supplier")
	{
		supplierController := &controller.SupplierController{}
		// 供应商列表
		supplier.GET("", supplierController.SupplierList)
		// 创建供应商
		supplier.POST("", supplierController.CreateSupplier)
		// 更新供应商
		supplier.PUT("", supplierController.UpdateSupplier)
		// 供应商产品列表
		supplier.GET("/product", supplierController.SupplierProductList)
		// 创建供应商产品
		supplier.POST("/product", supplierController.CreateSupplierProduct)
		// 更新供应商产品
		supplier.PUT("/product", supplierController.UpdateSupplierProduct)
		// 供应商充值
		supplier.POST("/recharge", supplierController.RechargeSupplier)
		supplier.GET("/recharge", supplierController.GetSupplierRechargeList)
		supplier.PUT("/recharge", supplierController.UpdateSupplierRecharge)
		supplier.GET("/recharge/history", supplierController.GetSupplierRechargeHistoryList)
	}

	// 品牌、规格、SKU API端点
	r.POST("/project", controller.CreateModel)
	// r.GET("/project", controller.GetModelList)

	// 渠道API端点
	channelController := &controller.ChannelController{}
	channel := r.Group("/channel")
	{
		// 创建渠道
		channel.POST("", channelController.CreateChannel)
		// 渠道列表
		channel.GET("", channelController.GetChannelList)
		// 更新渠道
		channel.PUT("", channelController.UpdateChannel)

		newproject := channel.Group("/project")
		{
			// 创建项目
			newproject.GET("", channelController.GetProjectList)
			newproject.POST("", channelController.CreateProject)
			newproject.PUT("", channelController.UpdateProject)

			newproduct := newproject.Group("/product")
			{
				// 创建项目产品
				newproduct.POST("", channelController.CreateProjectProduct)
				newproduct.GET("", channelController.GetProjectProductList)
				newproduct.PUT("", channelController.UpdateProjectProduct)
			}
		}

	}
	// 创建渠道供应商产品
	r.GET("/product", channelController.GetChannelSupplierProductList)
	r.POST("/product", channelController.CreateChannelSupplierProduct)

}
