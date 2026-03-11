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
		// order.GET("", orderController.GetOrderList)
		// order.PUT("", orderController.UpdateOrder)
	}
}
