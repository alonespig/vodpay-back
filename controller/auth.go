package controller

import (
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}

	// 登录
	resp, err := service.UserLogin(req.Username, req.Password)
	if err != nil {
		if err == service.UserNamePasswordError {
			BadRequest(c, err.Error())
			return
		}
		ServerError(c, err.Error())
		return
	}

	Success(c, resp)
}

func Register(c *gin.Context) {
	var req RegisterForm
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err.Error())
		return
	}

	// 注册用户
	err := service.UserRegister(req.Username, req.Password)
	if err != nil {
		if err == service.UserNameExist {
			Conflict(c, err.Error())
			return
		}
		ServerError(c, err.Error())
		return
	}

	Success(c, nil)
}
