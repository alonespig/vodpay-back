package form

import "vodpay/dto"

type LoginResp struct {
	Avatar string   `json:"avatar"`
	Token  string   `json:"token"`
	User   dto.User `json:"user"`
}
