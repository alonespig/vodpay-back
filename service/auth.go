package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"vodpay/database"
	"vodpay/dto"
	"vodpay/form"
	"vodpay/repository"
	"vodpay/utils"
)

// TokenInfo 存储在Redis中的token信息
type TokenInfo struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

func UserLogin(username, password string) (*form.LoginResp, error) {
	// 查询用户
	user, err := repository.GetUserByName(username)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, UserNamePasswordError
		}
		return nil, err
	}

	// 验证密码
	hashPassword := utils.MD5(password)
	if user.Password != hashPassword {
		return nil, UserNamePasswordError
	}

	token, err := utils.GenerateToken(user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	tokenInfo := dto.User{
		ID:   user.ID,
		Name: user.Name,
	}

	data, err := json.Marshal(&tokenInfo)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("vodpay-token:%d", user.ID)

	if err := database.Redis.Set(context.Background(), key, data, 2*time.Hour).Err(); err != nil {
		return nil, err
	}

	return &form.LoginResp{
		Avatar: "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png",
		Token:  token,
		User:   tokenInfo,
	}, nil
}

func UserRegister(username, password string) error {
	// 检查用户名是否已存在
	exist, err := repository.CheckUserExist(username)
	if err != nil {
		return err
	}
	if exist {
		return UserNameExist
	}

	// 创建用户
	hashPassword := utils.MD5(password)
	user := &repository.User{
		Name:     username,
		Password: hashPassword,
		Status:   1,
	}

	return repository.CreateUser(user)
}
