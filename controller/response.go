package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}

func BadRequest(c *gin.Context, msg string) {
	Fail(c, 400, msg)
}

func Unauthorized(c *gin.Context, msg string) {
	Fail(c, 401, msg)
}

func Forbidden(c *gin.Context, msg string) {
	Fail(c, 403, msg)
}

func Conflict(c *gin.Context, msg string) {
	Fail(c, 409, msg)
}

func ServerError(c *gin.Context, msg string) {
	Fail(c, 500, msg)
}
