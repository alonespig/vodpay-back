package main

import (
	"vodpay/database"
	"vodpay/mq"
	"vodpay/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDB()
	// 初始化Redis
	database.InitRedis()

	err := mq.InitProducer()
	if err != nil {
		panic(err)
	}

	go func() {
		err := mq.StartOrderConsumer()
		if err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8088")
}
