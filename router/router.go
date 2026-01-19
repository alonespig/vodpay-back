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
	r.GET("/api/supplier", controller.SupplierList)
	r.POST("/api/supplier", controller.CreateSupplier)
	return r
}
