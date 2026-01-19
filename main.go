package main

import (
	"vodpay/model"
	"vodpay/router"
)

func main() {
	// 初始化数据库
	model.InitDB()
	r := router.InitRouter()
	r.Run(":8088")
}
