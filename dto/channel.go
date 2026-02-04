package dto

import "time"

type Channel struct {
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
