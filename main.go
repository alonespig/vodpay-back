package main

import (
	"vodpay/database"
	"vodpay/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDB()
	// 初始化Redis
	database.InitRedis()

	// database.DB.AutoMigrate(&repository.Order{})

	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8088")
}
