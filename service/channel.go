package service

import (
	"errors"
	"log"
	"vodpay/dto"
	"vodpay/form"
	"vodpay/repository"
	"vodpay/utils"

	"github.com/google/uuid"
)

func CreateChannel(channel *form.CreateChannelForm) error {
	secretKey, err := utils.GenerateSecret()
	if err != nil {
		return err
	}
	return repository.CreateChannel(&repository.Channel{
		Name:        channel.Name,
		AppID:       uuid.NewString(),
		SecretKey:   secretKey,
		WhiteList:   channel.WhiteList,
		Status:      1,
		CreditLimit: channel.CreditLimit * 100, // 单位：分
	})
}

func GetChannelList() ([]dto.Channel, error) {
	channelList, err := repository.GetChannelList()
	if err != nil {
		return nil, err
	}
	channelDTOList := make([]dto.Channel, 0, len(channelList))
	for _, channel := range channelList {
		channelDTOList = append(channelDTOList, dto.Channel{
			ID:            channel.ID,
			Name:          channel.Name,
			AppID:         channel.AppID,
			SecretKey:     channel.SecretKey,
			WhiteList:     channel.WhiteList,
			Status:        channel.Status,
			Balance:       channel.Balance,
			CreditLimit:   channel.CreditLimit,
			CreditBalance: channel.CreditBalance,
			CreatedAt:     channel.CreatedAt,
		})
	}
	return channelDTOList, nil
}

func GetChannelByID(id int) (*dto.Channel, error) {
	channelModel, err := repository.GetChannelByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.Channel{
		ID:            channelModel.ID,
		Name:          channelModel.Name,
		AppID:         channelModel.AppID,
		SecretKey:     channelModel.SecretKey,
		WhiteList:     channelModel.WhiteList,
		Status:        channelModel.Status,
		Balance:       channelModel.Balance,
		CreditLimit:   channelModel.CreditLimit,
		CreditBalance: channelModel.CreditBalance,
		CreatedAt:     channelModel.CreatedAt,
	}, nil
}

func UpdateChannel(channel *form.UpdateChannelForm) error {
	oldChannelModel, err := repository.GetChannelByID(channel.ID)
	if err != nil {
		return err
	}
	if oldChannelModel.Name != channel.Name {
		return errors.New("channel name is not same")
	}
	oldChannelModel.WhiteList = channel.WhiteList
	oldChannelModel.Status = *channel.Status
	oldChannelModel.CreditLimit = channel.CreditLimit * 100 // 单位：分
	return repository.UpdateChannel(oldChannelModel)
}

func GetProjectList(query *form.ProjectQueryForm) (*dto.ProjectListResp, error) {
	projectList, err := repository.GetProjectList(&repository.ProjectQuery{
		ChannelID: query.ChannelID,
	})
	if err != nil {
		return nil, err
	}

	channel, err := repository.GetChannelByID(*query.ChannelID)
	if err != nil {
		return nil, err
	}

	resp := &dto.ProjectListResp{
		ChannelName:  channel.Name,
		ProjectsList: make([]dto.Project, 0, len(projectList)),
	}

	for _, project := range projectList {
		resp.ProjectsList = append(resp.ProjectsList, dto.Project{
			ID:        project.ID,
			Name:      project.Name,
			ChannelID: project.ChannelID,
			Status:    project.Status,
			CreatedAt: project.CreatedAt,
		})
	}
	return resp, nil
}

func CreateProjectProduct(form *form.CreateProjectProductForm) error {
	name, err := projectProductName(form.ProjectID, form.SKUID, form.BrandID, form.SpecID)
	if err != nil {
		log.Printf("match supplier product name failed, err: %v", err)
		return err
	}
	// 转换为分
	facePrice := int(form.FacePrice * 100)
	price := int(form.Price * 100)

	productID, err := repository.CreateProjectProduct(&repository.ProjectProduct{
		Name:      name,
		Status:    1,
		ProjectID: form.ProjectID,
		BrandID:   form.BrandID,
		SpecID:    form.SpecID,
		SKUID:     form.SKUID,
		FacePrice: facePrice,
		Price:     price,
	})
	if err != nil {
		return err
	}

	if len(form.SupplierProductIDList) == 0 {
		return nil
	}

	relations := make([]repository.ProductRelation, 0, len(form.SupplierProductIDList))
	for _, supplierProductID := range form.SupplierProductIDList {
		relations = append(relations, repository.ProductRelation{
			ChannelProductID:  productID,
			SupplierProductID: supplierProductID,
			Status:            1,
		})
	}
	return repository.CreateProductRelation(relations)
}
