package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

const (
	CodeSuccess      = 200
	CodeParamError   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeConflict     = 409
	CodeServerError  = 500
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
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
