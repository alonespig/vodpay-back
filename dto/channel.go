package dto

import "time"

type Supplier struct {
	ID          int
	Name        string
	AppID       string
	SecretKey   string
	WhiteList   string
	Status      int
	Balance     int
	CreditLimit int
	CreatedAt   time.Time
}

type SupplierListResp struct {
	Suppliers []Supplier
}

type Project struct {
	ID        int       `json:"id"`
	ChannelID int       `json:"channelID"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProjectListResp struct {
	ChannelName  string    `json:"channelName"`
	ProjectsList []Project `json:"projects"`
}

type Channel struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	AppID         string    `json:"appID"`
	SecretKey     string    `json:"secretKey"`
	WhiteList     string    `json:"whiteList"`
	Status        int       `json:"status"`
	Balance       int       `json:"balance"`
	CreditLimit   int       `json:"creditLimit"`
	CreditBalance int       `json:"creditBalance"`
	CreatedAt     time.Time `json:"createdAt"`
}
