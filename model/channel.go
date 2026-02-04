package model

import (
	"log"
	"time"
)

type Channel struct {
	ID            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	AppID         string    `db:"app_id" json:"appID"`
	SecretKey     string    `db:"secret_key" json:"secretKey"`
	WhiteList     string    `db:"white_list" json:"whiteList"`
	Status        int       `db:"status" json:"status"`
	Balance       int       `db:"balance" json:"balance"`
	CreditLimit   int       `db:"credit_limit" json:"creditLimit"`
	CreditBalance int       `db:"credit_balance" json:"creditBalance"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}

type ChannelSupplierProduct struct {
	ID                int       `db:"id" json:"id"`
	ChannelProductID  int       `db:"channel_product_id" json:"channelProductID"`
	SupplierProductID int       `db:"supplier_product_id" json:"supplierProductID"`
	CreatedAt         time.Time `db:"created_at" json:"createdAt"`
}

func GetChannelSupplierProductList() ([]ChannelSupplierProduct, error) {
	var channelSupplierProducts []ChannelSupplierProduct
	err := db.Select(&channelSupplierProducts, "SELECT * FROM channel_supplier_products")
	if err != nil {
		return nil, err
	}
	return channelSupplierProducts, nil
}

type ChannelSupplierRelation struct {
	ChannelProductID  int   `db:"channel_product_id" json:"channelProductID"`
	SupplierProductID []int `db:"supplier_product_id" json:"supplierProductID"`
}

func GetChannelSupplierRelation() ([]ChannelSupplierRelation, error) {
	var rows []struct {
		ChannelProductID  int `db:"channel_product_id"`
		SupplierProductID int `db:"supplier_product_id"`
	}

	err := db.Select(&rows, `
		SELECT channel_product_id, supplier_product_id
		FROM channel_supplier_products
	`)
	if err != nil {
		return nil, err
	}

	m := make(map[int][]int)

	for _, r := range rows {
		m[r.ChannelProductID] = append(
			m[r.ChannelProductID],
			r.SupplierProductID,
		)
	}

	var result []ChannelSupplierRelation
	for k, v := range m {
		result = append(result, ChannelSupplierRelation{
			ChannelProductID:  k,
			SupplierProductID: v,
		})
	}

	return result, nil
}

func GetChannelSupplierProductListByChannelProductID(channelProductID int) ([]ChannelSupplierProduct, error) {
	var channelSupplierProducts []ChannelSupplierProduct
	err := db.Select(&channelSupplierProducts, "SELECT * FROM channel_supplier_products WHERE channel_product_id = ?", channelProductID)
	if err != nil {
		return nil, err
	}
	return channelSupplierProducts, nil
}

func CreateChannelProjectProduct(channelProjectProductID int, supplierProductID []int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	for _, productID := range supplierProductID {
		sqlStr := `INSERT INTO channel_supplier_products 
		(channel_product_id, supplier_product_id) 
		VALUES (?, ?)`
		_, err := tx.Exec(sqlStr, channelProjectProductID, productID)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateChannel(channel *Channel) error {
	log.Printf("create channel, %v", channel)
	sqlStr := `INSERT INTO channels 
	(name, app_id, secret_key, white_list, status, balance, credit_limit, credit_balance) 
	VALUES (:name, :app_id, :secret_key, :white_list, :status, :balance, :credit_limit, :credit_balance)`
	_, err := db.NamedExec(sqlStr, channel)
	if err != nil {
		return err
	}
	return nil
}

func GetChannelList() ([]Channel, error) {
	var channels []Channel
	err := db.Select(&channels, "SELECT * FROM channels")
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func GetChannelByID(id int) (*Channel, error) {
	var channel Channel
	err := db.Get(&channel, "SELECT * FROM channels WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func UpdateChannel(channel *Channel) error {
	log.Printf("update channel, %v", channel)
	sqlStr := `UPDATE channels SET name = :name, white_list = :white_list, status = :status, credit_limit = :credit_limit WHERE id = :id`
	_, err := db.NamedExec(sqlStr, channel)
	if err != nil {
		return err
	}
	return nil
}
