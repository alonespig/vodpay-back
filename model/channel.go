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
