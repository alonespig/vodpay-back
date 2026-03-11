package repository

import (
	"log"
	"time"
	"vodpay/common"
	"vodpay/database"

	"gorm.io/gorm"
)

func CreateOrder(order *Order) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(order).Error
		if err != nil {
			return err
		}
		// 扣减渠道余额
		err = tx.Model(&Channel{}).Where("id = ?", order.ChannelID).
			Update("balance", gorm.Expr("balance - ?", order.ChannelPrice)).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func GetOrderByChannelOrder(channelID int, channelOrderNo string) (*Order, error) {
	var order Order
	err := database.DB.Where("channel_id = ? AND channel_order_no = ?", channelID, channelOrderNo).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

type OrderListQuery struct {
	Page               int
	PageSize           int
	AccountID          string
	SupplierID         int64
	ProjectID          int64
	ProductID          int64
	ChannelOrderNo     string
	BrandSkuSpecIDList []int64
	StartTime          *time.Time
	EndTime            *time.Time
	Status             *int
}

func GetOrderList(q *OrderListQuery) (int64, []*Order, error) {
	var orders []*Order
	query := database.DB.Model(&Order{})
	if q.AccountID != "" {
		query = query.Where("account_id = ?", q.AccountID)
	}
	if q.SupplierID != 0 {
		query = query.Where("supplier_id = ?", q.SupplierID)
	}
	if q.ProjectID != 0 {
		query = query.Where("project_id = ?", q.ProjectID)
	}
	if q.ProductID != 0 {
		query = query.Where("product_id = ?", q.ProductID)
	}
	if q.ChannelOrderNo != "" {
		query = query.Where("channel_order_no = ?", q.ChannelOrderNo)
	}
	if q.StartTime != nil {
		query = query.Where("create_time >= ?", *q.StartTime)
	}
	if q.EndTime != nil {
		query = query.Where("create_time <= ?", *q.EndTime)
	}
	if q.Status != nil {
		query = query.Where("status = ?", *q.Status)
	}
	if len(q.BrandSkuSpecIDList) > 0 {
		query = query.Where("brand_spec_sku_id IN ?", q.BrandSkuSpecIDList)
	}
	var orderCount int64

	if err := query.
		Count(&orderCount).Error; err != nil {
		log.Printf("[GetOrderList] page = %d, size = %d: %v", q.Page, q.PageSize, err)
		return 0, nil, err
	}

	if err := query.
		Offset((q.Page - 1) * q.PageSize).
		Limit(q.PageSize).
		Order("id DESC").
		Find(&orders).Error; err != nil {
		log.Printf("[GetOrderList] page = %d, size = %d: %v", q.Page, q.PageSize, err)
		return 0, nil, err
	}
	return orderCount, orders, nil
}

func CountTodayOrderByAccountID(accountID string) (int64, error) {
	var count int64
	startTime, endTime := common.GetTodayTimeRange()
	if err := database.DB.
		Model(&Order{}).
		Where("status = ? AND account_id = ? AND create_time >= ? AND create_time <= ?",
			common.StatusSuccess, accountID, startTime, endTime).
		Count(&count).Error; err != nil {
		log.Printf("[CountTodayOrderByAccountID] accountID = %s: %v", accountID, err)
		return 0, err
	}
	return count, nil
}

func GetOrderByStatus(status int) ([]*Order, error) {
	var orders []*Order
	if err := database.DB.
		Model(&Order{}).
		Where("status = ?", status).
		Find(&orders).Error; err != nil {
		log.Printf("[GetOrderByStatus] status = %d: %v", status, err)
		return nil, err
	}
	return orders, nil
}
