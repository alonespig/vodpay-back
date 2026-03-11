package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// 全局Redis客户端
var Redis *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123",
		DB:       0,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接Redis失败: %v", err)
	}
	log.Println("成功连接到Redis")
	Redis = client
}
