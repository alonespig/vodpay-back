package router

import (
	"vodpay/controller"

	"github.com/gin-gonic/gin"
)

func InitProductRouter(r *gin.Engine) {
	productController := &controller.ProductController{}
	product := r.Group("/product")
	{
		// 创建项目产品
		product.POST("", productController.CreateProduct)
		// 获取项目产品列表
		product.GET("/list", productController.GetProductList)
		// 获取项目产品的供应商产品列表
		product.GET("/:productID/supplier", productController.GetProductSupplierList)

		// 更新项目产品关联
		product.POST("/relation", productController.UpdateProductRelation)
		// 更新项目产品
		product.PUT("", productController.UpdateProduct)
	}
}
