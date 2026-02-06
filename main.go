package main

import (
	"vodpay/database"
	"vodpay/model"
	"vodpay/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDB()
	model.InitDB()
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8088")
}
