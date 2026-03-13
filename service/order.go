package service

import (
	"errors"
	"log"
	"strings"
	"time"
	"vodpay/client"
	"vodpay/common"
	"vodpay/form"
	"vodpay/mq"
	"vodpay/repository"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func CreateOrder(form *form.OrderForm) (string, error) {
	channel, err := repository.GetChannelByAppID(form.Appid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrAppIDNotFound
		}
		log.Printf("[CreateOrder] appID = %s: %v", form.Appid, err)
		return "", ErrSystemError
	}
	// 检查渠道是否启用
	if channel.Status == 0 {
		return "", ErrChannelNotEnabled
	}
	// 检查IP是否在白名单中
	if channel.WhiteList != common.WhiteListAll && !strings.Contains(channel.WhiteList, form.Ip) {
		return "", ErrChannelNotEnabled
	}

	// sign := client.CheckOrderSign(*form, channel.SecretKey)
	// if sign != form.Sign {
	// 	log.Printf("[CreateOrder] sign = %s, checkSign = %s", form.Sign, sign)
	// 	return "", ErrSignNotMatch
	// }

	// 检查产品是否存在
	product, err := repository.GetProductByID(int64(form.ProductCode))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrProductNotFound
		}
		log.Printf("[CreateOrder] productCode = %d: %v", form.ProductCode, err)
		return "", ErrSystemError
	}
	// 检查项目是否存在
	project, err := repository.GetProjectByID(product.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrProjectNotFound
		}
		log.Printf("[CreateOrder] projectID = %d: %v", product.ProjectID, err)
		return "", ErrSystemError
	}
	// 检查项目是否启用
	if project.Status == 0 {
		return "", ErrProjectNotEnabled
	}
	// 检查产品是否启用
	if product.Status == 0 {
		return "", ErrProductNotEnabled
	}

	// 检查用户已经下了多少单
	orderCount, err := repository.CountTodayOrderByAccountID(form.Mobile)
	if err != nil {
		log.Printf("[CreateOrder] accountID = %s: %v", form.Mobile, err)
		return "", ErrSystemError
	}
	if orderCount >= int64(product.LimitCount) && product.LimitCount != 0 {
		return "", ErrAccountOrderLimitExceeded
	}

	order := &repository.Order{
		ProductID:       int64(form.ProductCode),
		ProductName:     product.Name,
		SupplierID:      int64(product.SupplierID),
		SupplierName:    product.SupplierName,
		SupProductCode:  product.SupplierProductCode,
		ChannelIP:       form.Ip,
		ChannelID:       channel.ID,
		ChannelName:     channel.Name,
		ChannelPrice:    product.Price,
		ChannelOrderNo:  form.ChannelOrderNo,
		SelfOrderNo:     client.NewUuid(),
		PlatFromOrderNo: "no",
		AccountID:       form.Mobile,
		Msg:             "",
		Status:          common.StatusWait,
		CallBack:        form.Callback,
		RetryID:         0,
	}
	if channel.Balance+channel.CreditLimit < product.Price {
		return "", ErrChannelBalanceNotEnough
	}
	// 创建订单
	err = repository.CreateOrder(order)
	if err != nil {
		log.Printf("[CreateOrder] order = %v: %v", order, err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return "", ErrChannelOrderNoExist
			}
		}
		return "", ErrSystemError
	}
	// 发送订单创建消息
	if err := mq.SendOrderCreated(order.ID); err != nil {
		log.Printf("[CreateOrder] SendOrderCreated: %v", err)
	}
	return order.SelfOrderNo, nil
}

func QueryOrder(queryForm *form.OrderQueryForm) (*form.OrderQueryResponse, error) {
	channel, err := repository.GetChannelByAppID(queryForm.Appid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppIDNotFound
		}
		log.Printf("[QueryOrder] appID = %s: %v", queryForm.Appid, err)
		return nil, ErrSystemError
	}
	// TODO: 检查签名
	// sign := client.CheckOrderQuerySign(*queryForm, channel.SecretKey)
	// if sign != queryForm.Sign {
	// 	log.Printf("[QueryOrder] sign = %s, checkSign = %s", queryForm.Sign, channel.SecretKey)
	// 	return nil, ErrSignNotMatch
	// }

	order, err := repository.GetOrderByChannelOrder(channel.ID,
		queryForm.ChannelOrderNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		log.Printf("[QueryOrder] channelID = %d, orderNo = %s: %v", channel.ID, queryForm.ChannelOrderNo, err)
		return nil, ErrSystemError
	}
	resp := form.OrderQueryResponse{
		Code:            0,
		Msg:             "查询成功",
		OrderStatus:     order.Status,
		PlatformOrderNo: order.SelfOrderNo,
	}
	return &resp, nil
}

func GetOrderList(queryForm *form.OrderListQueryForm) (*form.OrderListResp, error) {
	query := &repository.OrderListQuery{
		Page:               queryForm.Page,
		PageSize:           queryForm.PageSize,
		AccountID:          queryForm.AccountID,
		SupplierID:         int64(queryForm.SupplierID),
		ProjectID:          int64(queryForm.ProjectID),
		ProductID:          int64(queryForm.ProductID),
		ChannelOrderNo:     queryForm.ChannelOrderNo,
		Status:             queryForm.Status,
		BrandSkuSpecIDList: make([]int64, 0),
	}
	if queryForm.StartTime != 0 {
		startTime := time.UnixMilli(int64(queryForm.StartTime))
		query.StartTime = &startTime
	}
	if queryForm.EndTime != 0 {
		endTime := time.UnixMilli(int64(queryForm.EndTime) + 24*60*60*1000)
		query.EndTime = &endTime
	}
	if queryForm.BrandSkuSpecID != 0 {
		query.BrandSkuSpecIDList = append(query.BrandSkuSpecIDList, int64(queryForm.BrandSkuSpecID))
	} else if queryForm.BrandID != 0 || queryForm.SKUID != 0 || queryForm.SpecID != 0 {
		brandSkuSpecIDList, err := repository.GetBBSIDListByID(int64(queryForm.BrandID), int64(queryForm.SKUID), int64(queryForm.SpecID))
		if err != nil {
			log.Printf("[GetOrderList] brandID = %d, skuID = %d, specID = %d: %v", queryForm.BrandID, queryForm.SKUID, queryForm.SpecID, err)
			return nil, ErrSystemError
		}
		query.BrandSkuSpecIDList = append(query.BrandSkuSpecIDList, brandSkuSpecIDList...)
	}
	orderCount, orders, err := repository.GetOrderList(query)
	if err != nil {
		log.Printf("[GetOrderList] page = %d, size = %d: %v", queryForm.Page, queryForm.PageSize, err)
		return nil, ErrSystemError
	}
	resp := form.OrderListResp{
		Total:  orderCount,
		Orders: make([]*form.Order, 0, len(orders)),
	}
	for _, order := range orders {
		product, err := repository.GetProductByID(order.ProductID)
		if err != nil {
			log.Printf("[GetOrderList] productID = %d: %v", order.ProductID, err)
		}
		project, err := repository.GetProjectByID(product.ProjectID)
		if err != nil {
			log.Printf("[GetOrderList] projectID = %d: %v", product.ProjectID, err)
		}
		resp.Orders = append(resp.Orders, &form.Order{
			ID:              order.ID,
			ProductID:       order.ProductID,
			ProductName:     order.ProductName,
			SupplierID:      order.SupplierID,
			SupplierName:    order.SupplierName,
			ProjectName:     project.Name,
			ChannelName:     order.ChannelName,
			Price:           order.ChannelPrice,
			PlatFromOrderNo: order.PlatFromOrderNo,
			SelfOrderNo:     order.SelfOrderNo,
			ChannelOrderNo:  order.ChannelOrderNo,
			AccountID:       order.AccountID,
			Msg:             order.Msg,
			Status:          order.Status,
			CreatedAt:       order.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       order.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &resp, nil
}

func GetChannelLineChart(queryForm *form.ChannelLineChartQueryForm) (*form.ChannelLineChartResp, error) {
	// TODO: 实现渠道折线图
	// status := common.StatusSuccess
	query := &repository.OrderListQuery{
		ChannelID: int64(queryForm.ChannelID),
		ProjectID: int64(queryForm.ProjectID),
		ProductID: int64(queryForm.ProductID),
		// Status:    &status,
	}
	if queryForm.Timestamp != 0 {
		startTime := time.UnixMilli(int64(queryForm.Timestamp))
		endTime := time.UnixMilli(int64(queryForm.Timestamp) + 24*60*60*1000)
		query.StartTime = &startTime
		query.EndTime = &endTime
	}

	_, orders, err := repository.GetOrderList(query)
	if err != nil {
		log.Printf("[GetChannelLineChart] channelID = %d, projectID = %d, productID = %d: %v", queryForm.ChannelID, queryForm.ProjectID, queryForm.ProductID, err)
		return nil, ErrSystemError
	}

	log.Printf("orders = %v", orders)

	mapStatus := map[int]int64{}

	for _, order := range orders {
		mapStatus[order.CreatedAt.Hour()]++
	}

	log.Printf("mapStatus = %v", mapStatus)

	resp := form.ChannelLineChartResp{
		Points: make([]*form.Point, 0),
	}

	for hour := 0; hour < 24; hour++ {
		resp.Points = append(resp.Points, &form.Point{
			X: hour,
			Y: int(mapStatus[hour]),
		})
	}

	return &resp, nil
}

func GetSupplierOrderList(queryForm *form.SupplierOrderListQueryForm) (*form.SupplierOrderListResp, error) {
	query := &repository.SupplierOrderListQuery{
		SupplierID:     int64(queryForm.SupplierID),
		Status:         queryForm.Status,
		BrandSkuSpecID: int64(queryForm.BrandSkuSpecID),
	}
	startTime := time.UnixMilli(int64(queryForm.StartTime))
	endTime := time.UnixMilli(int64(queryForm.EndTime) + 24*60*60*1000)
	query.StartTime = &startTime
	query.EndTime = &endTime
	orderCount, orders, err := repository.GetSupplierOrderList(query)
	if err != nil {
		log.Printf("[GetSupplierOrderList] supplierID = %d: %v", queryForm.SupplierID, err)
		return nil, ErrSystemError
	}

	resp := form.SupplierOrderListResp{
		Total:  orderCount,
		Orders: make([]form.SupplierOrder, 0, len(orders)),
	}

	for _, order := range orders {
		product, err := repository.GetSupplierProductCode(order.SupProductCode)
		if err != nil {
			log.Printf("[GetSupplierOrderList] supProductCode = %s: %v", order.SupProductCode, err)
			return nil, ErrSystemError
		}
		resp.Orders = append(resp.Orders, form.SupplierOrder{
			SupplierName:   product.SupplierName,
			Name:           product.Name,
			Price:          order.Price,
			SupProductCode: order.SupProductCode,
			Total:          order.Total,
			Status:         order.Status,
		})
	}

	return &resp, nil
}

// 将下单中的订单状态改成充值中
func HandleOrder(orderID int64) error {
	return nil
}
