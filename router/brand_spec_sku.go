package router

import (
	"vodpay/controller"

	"github.com/gin-gonic/gin"
)

func InitBrandSkuSpecRouter(r *gin.Engine) {
	controller := &controller.BSSController{}
	api := r.Group("/brand-spec-sku")
	{
		api.GET("", controller.GetList)
		api.POST("", controller.Create)
		api.GET("/brand", controller.GetBrandList)
		api.GET("/sku", controller.GetSkuList)
		api.GET("/spec", controller.GetSpecList)
	}
}
