package repository

import "errors"

var (
	ErrSupplierNotFound         = errors.New("供应商不存在")
	ErrChannelNotFound          = errors.New("渠道不存在")
	ErrSupplierProductNotFound  = errors.New("供应商商品不存在")
	ErrSkuNotFound              = errors.New("SKU不存在")
	ErrBrandNotFound            = errors.New("品牌不存在")
	ErrSpecNotFound             = errors.New("规格不存在")
	ErrSupplierRechargeNotFound = errors.New("供应商充值不存在")
	ErrProjectNotFound          = errors.New("项目不存在")
	ErrProjectProductNotFound   = errors.New("项目商品不存在")
)
