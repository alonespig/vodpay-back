package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"vodpay/database"
	"vodpay/dto"
	"vodpay/pkg/response"
	"vodpay/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		auth := c.GetHeader("Authorization")
		if auth == "" {
			response.Unauthorized(c, "未授权")
			c.Abort()
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(auth, prefix) {
			response.Unauthorized(c, "非法 token")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(auth, prefix)
		if token == "" {
			response.Unauthorized(c, "非法 token")
			c.Abort()
			return
		}

		// 验证token
		claims, err := utils.ParseToken(token)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		key := fmt.Sprintf("vodpay-token:%d", claims.UserID)

		data, err := database.Redis.Get(c, key).Result()
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		userInfo := dto.User{}

		err = json.Unmarshal([]byte(data), &userInfo)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		// 将用户ID存储在上下文
		c.Set("userID", userInfo.ID)
		c.Next()
	}
}
