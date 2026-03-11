package router

import (
	"vodpay/controller"

	"github.com/gin-gonic/gin"
)

func InitSupplierRouter(r *gin.Engine) {
	supplierController := &controller.SupplierController{}
	supplier := r.Group("/supplier")
	{
		// 供应商列表
		supplier.GET("", supplierController.SupplierList)
		// 创建供应商
		supplier.POST("", supplierController.CreateSupplier)
		// 更新供应商
		supplier.PUT("", supplierController.UpdateSupplier)

		// 获取某个供应商的产品列表
		supplier.GET(":id/product", supplierController.SupplierProduct)

		product := supplier.Group("/product")
		{
			// 供应商产品列表
			product.GET("", supplierController.SupplierProductList)
			// 创建供应商产品
			product.POST("", supplierController.CreateSupplierProduct)
			// 更新供应商产品
			product.PUT("", supplierController.UpdateSupplierProduct)
		}
	}
}
