package service

type ServiceError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ServiceError) Error() string {
	return e.Msg
}

// 服务层错误
var (
	ErrAppIDNotFound             = &ServiceError{Code: 400, Msg: "appid不存在"}
	ErrChannelNotEnabled         = &ServiceError{Code: 401, Msg: "系统渠道未启用"}
	ErrSignNotMatch              = &ServiceError{Code: 402, Msg: "签名不匹配"}
	ErrProjectNotFound           = &ServiceError{Code: 403, Msg: "项目不存在"}
	ErrProjectNotEnabled         = &ServiceError{Code: 404, Msg: "项目未启用"}
	ErrProductNotFound           = &ServiceError{Code: 405, Msg: "产品不存在"}
	ErrProductNotEnabled         = &ServiceError{Code: 406, Msg: "产品未启用"}
	ErrOrderNotFound             = &ServiceError{Code: 407, Msg: "订单不存在"}
	ErrChannelNotExist           = &ServiceError{Code: 408, Msg: "渠道不存在"}
	ErrAccountOrderLimitExceeded = &ServiceError{Code: 409, Msg: "账号已下订单数已达上限"}
	ErrChannelOrderNoExist       = &ServiceError{Code: 410, Msg: "渠道订单号重复"}
	ErrChannelBalanceNotEnough   = &ServiceError{Code: 411, Msg: "渠道余额不足"}
)

var (
	ErrSystemError = &ServiceError{Code: 500, Msg: "系统错误"}
)
