package form

import "time"

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

type CreateChannelForm struct {
	Name        string `json:"name" binding:"required"`
	WhiteList   string `json:"whiteList" binding:"required"`
	CreditLimit int    `json:"creditLimit" binding:"required,gt=0"`
}

type UpdateChannelForm struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	WhiteList   string `json:"whiteList" binding:"required"`
	Status      *int   `json:"status" binding:"required,oneof=0 1"`
	CreditLimit int    `json:"creditLimit" binding:"required,gt=0"`
}

type CreateChannelSupplierProductForm struct {
	ProjectProductID  int   `json:"projectProductID" binding:"required"`
	SupplierProductID []int `json:"supplierProductIDList" binding:"required"`
}

type ProjectQueryForm struct {
	ChannelID *int `form:"channelID" binding:"required"`
}

type CreateProjectForm struct {
	ChannelID int     `json:"channelID" binding:"required"`
	Name      *string `json:"name" binding:"required"`
}

type UpdateProjectStatusForm struct {
	ID     int  `json:"id" binding:"required"`
	Status *int `json:"status" binding:"required,oneof=0 1"`
}

type ProjectProductQueryForm struct {
	ProjectID *int `form:"projectID" binding:"required"`
}
