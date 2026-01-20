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

	// 供应商API端点
	r.GET("/supplier", controller.SupplierList)
	r.POST("/supplier", controller.CreateSupplier)
	r.PUT("/supplier", controller.UpdateSupplier)
	r.GET("/supplier/product", controller.SupplierProductList)
	r.POST("/supplier/product", controller.CreateSupplierProduct)
	r.PUT("/supplier/product", controller.UpdateSupplierProduct)
	r.POST("/supplier/recharge", controller.RechargeSupplier)

	// 品牌、规格、SKU API端点
	r.POST("/project", controller.CreateModel)
	r.GET("/project", controller.GetModelList)
	return r
}
