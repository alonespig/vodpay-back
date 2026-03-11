package common

type ErrorCode struct {
	Code    int
	Message string
}

func (e ErrorCode) Error() string {
	return e.Message
}

var (
	ErrServiceUnavailable = ErrorCode{2001, "服务不可用"}
	ErrRPCFailed          = ErrorCode{2002, "远程服务调用失败"}
	ErrNetwork            = ErrorCode{2003, "网络异常"}
)

var (
	ErrDB               = ErrorCode{3000, "数据库错误"}
	ErrDBConnect        = ErrorCode{3001, "数据库连接失败"}
	ErrDBQuery          = ErrorCode{3002, "数据库查询失败"}
	ErrDBInsert         = ErrorCode{3003, "数据库插入失败"}
	ErrDBDuplicateKey   = ErrorCode{3004, "数据重复"}
	ErrDBRecordNotFound = ErrorCode{3005, "数据不存在"}
)

var (
	ErrInvalidParam             = ErrorCode{4001, "参数不合法"}
	ErrBusinessFailed           = ErrorCode{4002, "业务处理失败"}
	ErrAlreadyExists            = ErrorCode{4003, "资源已存在"}
	ErrSupplierNotFound         = ErrorCode{4004, "供应商不存在"}
	ErrChannelNotFound          = ErrorCode{4005, "渠道不存在"}
	ErrSupplierProductNotFound  = ErrorCode{4006, "供应商商品不存在"}
	ErrSkuNotFound              = ErrorCode{4007, "SKU不存在"}
	ErrBrandNotFound            = ErrorCode{4008, "品牌不存在"}
	ErrSpecNotFound             = ErrorCode{4009, "规格不存在"}
	ErrSupplierRechargeNotFound = ErrorCode{4010, "供应商充值不存在"}
	ErrProjectNotFound          = ErrorCode{4011, "项目不存在"}
	ErrProductNotFound          = ErrorCode{4012, "商品不存在"}
	ErrUserNotFound             = ErrorCode{4013, "用户不存在"}
	ErrProductRelationNotFound  = ErrorCode{4014, "商品关联不存在"}
	ErrSystemError              = ErrorCode{5000, "系统错误"}
)

var (
	ErrUnauthorized = ErrorCode{5001, "未登录或登录失效"}
	ErrForbidden    = ErrorCode{5002, "无权限访问"}
	ErrTokenExpired = ErrorCode{5003, "Token 已过期"}
)
