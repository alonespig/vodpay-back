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

	// 认证路由
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}

	InitOrderRouter(r)

	// 认证中间件
	// r.Use(middleware.AuthMiddleware())

	// 上传文件
	r.POST("/upload", controller.Upload)

	InitChannelRouter(r)

	InitSupplierRouter(r)

	InitProductRouter(r)

	InitBrandSkuSpecRouter(r)
}
