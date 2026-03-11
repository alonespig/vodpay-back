package router

import (
	"vodpay/controller"

	"github.com/gin-gonic/gin"
)

func InitChannelRouter(r *gin.Engine) {
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

		project := channel.Group("/project")
		{
			// 创建项目
			project.POST("", channelController.CreateProject)
			// 项目列表
			project.GET("", channelController.GetProjectList)
			// 更新项目
			project.PUT("", channelController.UpdateProject)
		}
	}
}
