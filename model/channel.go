package model

import "time"

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
	sqlStr := `INSERT INTO channels 
	(name, app_id, secret_key, white_list, status, balance, credit_limit, credit_balance) 
	VALUES (:name, :app_id, :secret_key, :white_list, :status, :balance, :credit_limit, :credit_balance)`
	_, err := db.NamedExec(sqlStr, channel)
	if err != nil {
		return err
	}
	return nil
}
