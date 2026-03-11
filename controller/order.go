package controller

import (
	"log"
	"net/http"
	"vodpay/client"
	"vodpay/form"
	"vodpay/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	form := &form.OrderForm{}
	if err := ctx.ShouldBind(form); err != nil {
		log.Printf("[CreateOrder] ShouldBind: %v", err)
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	if client.CheckTimestamp(form.Timestamp) != true {
		log.Printf("[CreateOrder] CheckTimestamp: %v", form.Timestamp)
		Fail(ctx, CodeParamError, "timestamp error")
		return
	}
	form.Ip = ctx.ClientIP()
	selfOrderNo, err := service.CreateOrder(form)
	if err != nil {
		log.Printf("[CreateOrder] CreateOrder: %v", err)
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":              0,
		"msg":               "收单成功",
		"order_status":      0,
		"platform_order_no": selfOrderNo,
	})
}

func (c *OrderController) QueryOrder(ctx *gin.Context) {
	queryForm := &form.OrderQueryForm{}
	if err := ctx.ShouldBind(queryForm); err != nil {
		log.Printf("[QueryOrder] ShouldBind: %v", err)
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	if client.CheckTimestamp(queryForm.Timestamp) != true {
		Fail(ctx, CodeParamError, "timestamp error")
		return
	}
	resp, err := service.QueryOrder(queryForm)
	if err != nil {
		log.Printf("[QueryOrder] QueryOrder: %v", err)
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *OrderController) GetOrderList(ctx *gin.Context) {
	queryForm := &form.OrderListQueryForm{}
	if err := ctx.ShouldBind(queryForm); err != nil {
		log.Printf("[GetOrderList] ShouldBind: %v", err)
		Fail(ctx, CodeParamError, err.Error())
		return
	}
	resp, err := service.GetOrderList(queryForm)
	if err != nil {
		log.Printf("[GetOrderList] GetOrderList: %v", err)
		Fail(ctx, CodeServerError, err.Error())
		return
	}
	Success(ctx, resp)
}
