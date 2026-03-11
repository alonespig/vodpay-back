package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt密钥
var jwtSecret = []byte("vodpay_secret_key")

// Claims JWT声明结构
type Claims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		UserID: userID,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "vodpay",
			IssuedAt:  nowTime.Unix(),
			NotBefore: nowTime.Unix(),
			Id:        fmt.Sprintf("%d_%d", userID, nowTime.Unix()), // 唯一标识
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseToken 从Redis中验证token
func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		// 检查token是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, fmt.Errorf("token已过期")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("无效的token")
}
