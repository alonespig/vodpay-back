package form

type CreateChannelForm struct {
	Name        string `json:"name" binding:"required"`
	WhiteList   string `json:"whiteList" binding:"required"`
	CreditLimit int    `json:"creditLimit" binding:"required,gt=0"`
}
