package router

import (
	"vodpay/controller"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(r *gin.Engine) {
	orderController := &controller.OrderController{}
	order := r.Group("/api/order")
	{
		order.POST("", orderController.CreateOrder)
		order.GET("", orderController.GetOrderList)
		order.POST("/query", orderController.QueryOrder)

		// 渠道折线图
		order.GET("/channel-line-chart", orderController.GetChannelLineChart)

		order.GET("/supplier-order-list", orderController.GetSupplierOrderList)
	}
}
