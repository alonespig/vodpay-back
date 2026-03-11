package service

import "errors"

var (
	ErrUserNotFound                   = errors.New("用户不存在")
	ErrProductStatusDisabled          = errors.New("该产品处于禁用状态")
	ErrSupplierProductUsing           = errors.New("该供应商产品正在被使用")
	ErrProductNotFoundSupplierProduct = errors.New("该产品未关联供应商产品")
	ErrProductRelationStatusDisabled  = errors.New("该产品关联的供应商产品处于禁用状态")
	ErrCreateProductRelation          = errors.New("创建产品关联失败")
	// ErrSystemError                    = errors.New("系统错误")
	ErrProductRelationNotFound = errors.New("产品关联不存在")
)

var (
	UserNamePasswordError = errors.New("用户名或密码错误")
	UserNameExist         = errors.New("用户名已存在")
)
