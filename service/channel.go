package service

import (
	"errors"
	"vodpay/form"
	"vodpay/model"
	"vodpay/utils"
	"vodpay/view"

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

func CreateChannelProjectProduct(channelProjectProductID int, supplierProductID []int) error {
	return model.CreateChannelProjectProduct(channelProjectProductID, supplierProductID)
}

func GetChannelSupplierProductList() ([]view.ChannelSupplierProductView, error) {
	productList, err := model.GetChannelSupplierRelation()
	if err != nil {
		return nil, err
	}
	viewList := make([]view.ChannelSupplierProductView, 0, len(productList))
	// 其中的一个产品 id
	for _, product := range productList {
		// 获取他所有的渠道的供应商
		projectProduct, err := model.GetProjectProductByID(product.ChannelProductID)
		if err != nil {
			return nil, err
		}
		// 遍历所有的供应商 id
		supplierProductList := make([]model.SupplierProduct, 0, len(product.SupplierProductID))
		for _, supplierProductID := range product.SupplierProductID {
			supplierProduct, err := model.GetSupplierProductByID(supplierProductID)
			if err != nil {
				return nil, err
			}
			supplierProductList = append(supplierProductList, *supplierProduct)
		}
		viewList = append(viewList, view.ChannelSupplierProductView{
			ProjectProduct:  *projectProduct,
			SupplierProduct: supplierProductList,
		})
	}

	return viewList, nil
}
