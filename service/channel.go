package service

import (
	"errors"
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
		CreditLimit: channel.CreditLimit * 100, // 单位：分
	})
}

func GetChannelList() ([]model.Channel, error) {
	return model.GetChannelList()
}

func GetChannelByID(id int) (*model.Channel, error) {
	return model.GetChannelByID(id)
}

func UpdateChannel(channel *form.UpdateChannelForm) error {
	oldChannelModel, err := model.GetChannelByID(channel.ID)
	if err != nil {
		return err
	}
	if oldChannelModel.Name != channel.Name {
		return errors.New("channel name is not same")
	}
	oldChannelModel.WhiteList = channel.WhiteList
	oldChannelModel.Status = *channel.Status
	oldChannelModel.CreditLimit = channel.CreditLimit * 100 // 单位：分
	return model.UpdateChannel(oldChannelModel)
}
