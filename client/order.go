package client

import "vodpay/repository"

type RespOrder struct {
	Code    int
	Msg     string
	OrderID string
	Resp    string
}

func CallPlatformOrder(order *repository.Order) (*RespOrder, error) {
	switch order.SupplierCode {
	case "DEBUG":
		return &RespOrder{
			Code: 1,
			Msg:  "DEBUG",
		}, nil
	}
	return nil, nil
}

// 根据具体订单中的供应商，查询订单结果
func CallPlatformOrderQuery(order *repository.Order) (*RespOrder, error) {
	switch order.SupProductCode {
	case "DEBUG":
		return &RespOrder{
			Code:    0,
			Msg:     "success",
			OrderID: order.PlatFromOrderNo,
			Resp:    "DEBUG",
		}, nil
	}
	return nil, nil
}
