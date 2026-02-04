package service

import (
	"fmt"
	"log"
	"vodpay/form"
	"vodpay/model"
)

func CreateProject(channelID int, name *string) error {
	_, err := model.GetChannelByID(channelID)
	if err != nil {
		return err
	}
	return model.CreateProject(&model.Project{
		ChannelID: channelID,
		Name:      *name,
		Status:    1,
	})
}

func GetProjectListByChannelID(id int) ([]model.Project, *model.Channel, error) {
	channel, err := model.GetChannelByID(id)
	if err != nil {
		return nil, nil, err
	}
	projects, err := model.GetProjectListByChannelID(id)
	if err != nil {
		return nil, nil, err
	}
	return projects, channel, nil
}

func UpdateProjectStatus(projectID, status int) error {
	project, err := model.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if project.Status == status {
		return nil
	}
	return model.UpdateProjectStatus(status, projectID)
}

func ProjectProductName(projectID, skuID, brandID, specID int) (string, error) {
	total, err := model.ProjectProductName(projectID, skuID, brandID, specID)
	if err != nil {
		log.Printf("get project product name failed, err: %v", err)
		return "", err
	}
	sku, err := model.GetModelByID("skus", skuID)
	if err != nil {
		log.Printf("get sku by id failed, err: %v", err)
		return "", err
	}
	brand, err := model.GetModelByID("brands", brandID)
	if err != nil {
		log.Printf("get brand by id failed, err: %v", err)
		return "", err
	}
	spec, err := model.GetModelByID("specs", specID)
	if err != nil {
		log.Printf("get spec by id failed, err: %v", err)
		return "", err
	}

	return fmt.Sprintf("%s%s%s-%d", brand.Name, spec.Name, sku.Name, total+1), nil
}

func CreateProjectProduct(product *model.ProjectProduct) error {
	name, err := ProjectProductName(product.ProjectID, product.SKUID, product.BrandID, product.SpecID)
	if err != nil {
		log.Printf("match supplier product name failed, err: %v", err)
		return err
	}
	product.Name = name
	return model.CreateProjectProduct(product)
}

func GetProjectProductListByProjectID(id int) ([]model.ProjectProduct, error) {
	return model.GetProjectProductListByProjectID(id)
}

func UpdateProjectProduct(channelID, projectID int, form *form.UpdateProjectProductForm) error {
	// 检查项目是否存在
	project, err := model.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if project.ChannelID != channelID {
		return fmt.Errorf("project %d not belong to channel %d", projectID, channelID)
	}
	projectProduct, err := model.GetProjectProductByID(form.ID)
	if err != nil {
		log.Printf("get project product by id failed, err: %v", err)
		return err
	}
	if projectProduct.ProjectID != projectID {
		return fmt.Errorf("project product %d not belong to project %d", form.ID, projectID)
	}
	projectProduct.Status = *form.Status
	projectProduct.FacePrice = int(form.FacePrice * 100)
	projectProduct.Price = int(form.Price * 100)
	return model.UpdateProjectProduct(projectProduct)
}

func GetSupplierRechargeList(status int) ([]model.SupplierRecharge, error) {
	return model.GetSupplierRechargeList(status)
}

func UpdateSupplierRecharge(form *form.SupplierRecharge) error {
	recharge, err := model.GetSupplierRechargeByID(form.ID)
	if err != nil {
		log.Printf("get supplier recharge by id failed, err: %v", err)
		return err
	}
	if recharge.SupplierID != form.SupplierID {
		return fmt.Errorf("supplier recharge %d not belong to supplier %d", form.ID, form.SupplierID)
	}
	if recharge.Amount != form.Amount*100 {
		return fmt.Errorf("supplier recharge %d amount not match, expect %d, got %d", form.ID, recharge.Amount, form.Amount)
	}
	recharge.Status = *form.Status
	recharge.Remark = &form.Remark
	return model.UpdateSupplierRecharge(recharge)
}
