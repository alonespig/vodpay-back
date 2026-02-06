package service

import (
	"fmt"
	"log"
	"vodpay/dto"
	"vodpay/form"
	"vodpay/repository"
)

func CreateProject(form *form.CreateProjectForm) error {
	_, err := repository.GetChannelByID(form.ChannelID)
	if err != nil {
		return err
	}
	return repository.CreateProject(&repository.Project{
		ChannelID: form.ChannelID,
		Name:      *form.Name,
		Status:    1,
	})
}

func GetProjectListByChannelID(id int) ([]dto.Project, *dto.Channel, error) {
	channel, err := repository.GetChannelByID(id)
	if err != nil {
		return nil, nil, err
	}
	projects, err := repository.GetProjectList(&repository.ProjectQuery{ChannelID: &id})
	if err != nil {
		return nil, nil, err
	}
	var projectDTOs []dto.Project
	for _, project := range projects {
		projectDTOs = append(projectDTOs, dto.Project{
			ID:        project.ID,
			ChannelID: project.ChannelID,
			Name:      project.Name,
			Status:    project.Status,
			CreatedAt: project.CreatedAt,
		})
	}
	channelDTO := &dto.Channel{
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
	}
	return projectDTOs, channelDTO, nil
}

func UpdateProjectStatus(projectID, status int) error {
	project, err := repository.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	if project.Status == status {
		return nil
	}
	project.Status = status
	log.Println(project)
	return repository.UpdateProject(project)
}

func projectProductName(projectID, skuID, brandID, specID int) (string, error) {
	products, err := repository.GetProjectProductList(&repository.ProjectProductQuery{
		ProjectID: &projectID,
		SKUID:     &skuID,
		BrandID:   &brandID,
		SpecID:    &specID,
	})
	if err != nil {
		log.Printf("get project product name failed, err: %v", err)
		return "", err
	}
	sku, err := repository.GetSkuByID(skuID)
	if err != nil {
		log.Printf("get sku by id failed, err: %v", err)
		return "", err
	}
	brand, err := repository.GetBrandByID(brandID)
	if err != nil {
		log.Printf("get brand by id failed, err: %v", err)
		return "", err
	}
	spec, err := repository.GetSpecByID(specID)
	if err != nil {
		log.Printf("get spec by id failed, err: %v", err)
		return "", err
	}

	return fmt.Sprintf("%s%s%s-%d", brand.Name, spec.Name, sku.Name, len(products)+1), nil
}

// func CreateProjectProduct(product *model.ProjectProduct) error {
// 	name, err := ProjectProductName(product.ProjectID, product.SKUID, product.BrandID, product.SpecID)
// 	if err != nil {
// 		log.Printf("match supplier product name failed, err: %v", err)
// 		return err
// 	}
// 	product.Name = name
// 	return repository.CreateProjectProduct(&repository.ProjectProduct{
// 		Name:      product.Name,
// 		Status:    product.Status,
// 		ProjectID: product.ProjectID,
// 		BrandID:   product.BrandID,
// 		SpecID:    product.SpecID,
// 		SKUID:     product.SKUID,
// 		FacePrice: product.FacePrice,
// 		Price:     product.Price,
// 	})
// }

func GetProjectProductListByProjectID(id int) ([]dto.ProjectProduct, error) {
	ProjectProduct, err := repository.GetProjectProductList(&repository.ProjectProductQuery{ProjectID: &id})
	if err != nil {
		return nil, err
	}
	var projectProductDTOs []dto.ProjectProduct
	for _, product := range ProjectProduct {
		projectProductDTOs = append(projectProductDTOs, dto.ProjectProduct{
			ID:        product.ID,
			Name:      product.Name,
			Status:    product.Status,
			ProjectID: product.ProjectID,
			BrandID:   product.BrandID,
			SpecID:    product.SpecID,
			SKUID:     product.SKUID,
			FacePrice: product.FacePrice,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
			Version:   product.Version,
		})
	}
	return projectProductDTOs, nil
}

func UpdateProjectProduct(form *form.UpdateProjectProductForm) error {
	projectProduct, err := repository.GetProjectProductByID(form.ID)
	if err != nil {
		log.Printf("get project product by id failed, err: %v", err)
		return err
	}
	projectProduct.Status = *form.Status
	projectProduct.FacePrice = int(form.FacePrice * 100)
	projectProduct.Price = int(form.Price * 100)
	return repository.UpdateProjectProduct(projectProduct)
}

func GetSupplierRechargeList(status int) ([]dto.SupplierRecharge, error) {
	rechargeList, err := repository.GetSupplierRechargeList(nil)
	if err != nil {
		log.Printf("get supplier recharge list failed, err: %v", err)
		return nil, err
	}
	rechargeDTOList := make([]dto.SupplierRecharge, 0, len(rechargeList))
	for _, recharge := range rechargeList {
		rechargeDTOList = append(rechargeDTOList, dto.SupplierRecharge{
			ID:            recharge.ID,
			SupplierName:  recharge.SupplierName,
			SupplierCode:  recharge.SupplierCode,
			Amount:        recharge.Amount,
			Status:        recharge.Status,
			ApplyUserName: recharge.ApplyUserName,
			AuditUserName: recharge.AuditUserName,
			ImageURL:      recharge.ImageURL,
			Remark:        recharge.Remark,
			CreatedAt:     recharge.CreatedAt,
			PassAt:        recharge.PassAt,
		})
	}
	return rechargeDTOList, nil
}

func UpdateSupplierRecharge(form *form.SupplierRecharge) error {
	recharge, err := repository.GetSupplierRechargeByID(form.ID)
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
	return repository.UpdateSupplierRecharge(recharge)
}

func GetProjectProductList(q *form.ProjectProductQueryForm) (*dto.ProjectProductResp, error) {
	project, err := repository.GetProjectByID(*q.ProjectID)
	if err != nil {
		log.Printf("get project by id failed, err: %v", err)
		return nil, err
	}
	projectProductList, err := repository.GetProjectProductList(&repository.ProjectProductQuery{
		ProjectID: q.ProjectID,
	})
	if err != nil {
		log.Printf("get project product list failed, err: %v", err)
		return nil, err
	}
	resp := &dto.ProjectProductResp{
		ProjectName: project.Name,
		ProductList: make([]dto.ProjectProductItem, 0, len(projectProductList)),
	}

	for _, product := range projectProductList {
		// 获取这个产品的供应商的产品id
		relations, err := repository.GetProductRelationList(product.ID)
		if err != nil {
			log.Printf("get product relation list failed, err: %v", err)
			return nil, err
		}
		dtoProduct := dto.ProjectProduct{
			ID:        product.ID,
			Name:      product.Name,
			Status:    product.Status,
			ProjectID: product.ProjectID,
			BrandID:   product.BrandID,
			SpecID:    product.SpecID,
			SKUID:     product.SKUID,
			FacePrice: product.FacePrice,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
			Version:   product.Version,
		}
		resp.ProductList = append(resp.ProductList, dto.ProjectProductItem{
			ProjectProduct:      dtoProduct,
			SupplierProductList: make([]dto.SupplierProduct, 0, len(relations)),
		})
		for _, relation := range relations {
			supplierProduct, err := repository.GetSupplierProductByID(relation.SupplierProductID)
			if err != nil {
				log.Printf("get supplier product by id failed, err: %v", err)
				return nil, err
			}

			dtoSuProduct := dto.SupplierProduct{
				ID:           supplierProduct.ID,
				Name:         supplierProduct.Name,
				Code:         supplierProduct.Code,
				SupplierID:   supplierProduct.SupplierID,
				SupplierName: supplierProduct.SupplierName,
				SupplierCode: supplierProduct.SupplierCode,
				Status:       supplierProduct.Status,
				BrandID:      supplierProduct.BrandID,
				SpecID:       supplierProduct.SpecID,
				SKUID:        supplierProduct.SKUID,
				FacePrice:    supplierProduct.FacePrice,
				Price:        supplierProduct.Price,
				CreatedAt:    supplierProduct.CreatedAt,
			}
			resp.ProductList[len(resp.ProductList)-1].SupplierProductList = append(resp.ProductList[len(resp.ProductList)-1].SupplierProductList, dtoSuProduct)
		}
	}
	return resp, nil
}
