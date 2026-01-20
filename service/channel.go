package service

import (
	"vodpay/form"
	"vodpay/model"
	"vodpay/utils"

	"github.com/google/uuid"
)

func CreateChannel(channel *form.CreateChannelForm) error {
	secretKey, err := utils.GenerateSecret()
	if err != nil {
		return err
	}
	return model.CreateChannel(&model.Channel{
		Name:        channel.Name,
		AppID:       uuid.NewString(),
		SecretKey:   secretKey,
		WhiteList:   channel.WhiteList,
		Status:      1,
		CreditLimit: channel.CreditLimit,
	})
}
