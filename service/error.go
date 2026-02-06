package service

import "errors"

var (
	ErrCreateProductRelation = errors.New("创建产品关联失败")
	ErrChannelNotFound       = errors.New("channel not found")
)
