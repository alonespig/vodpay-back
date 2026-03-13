package service

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"vodpay/dto"
	"vodpay/form"
	"vodpay/repository"

	"gorm.io/gorm"
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

func GetProjectsList(channelID int) (*form.ProjectListResp, error) {
	projects, err := repository.GetProjectList(&repository.ProjectQuery{ChannelID: &channelID})
	if err != nil {
		log.Printf("[GetProjectsList] channelID = %d: %v", channelID, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrChannelNotExist
		}
		return nil, ErrSystemError
	}
	var projectList []form.Project
	for _, project := range projects {
		projectList = append(projectList, form.Project{
			ID:   project.ID,
			Name: project.Name,
		})
	}
	return &form.ProjectListResp{
		ProjectList: projectList,
	}, nil
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
		ID:          channel.ID,
		Name:        channel.Name,
		AppID:       channel.AppID,
		SecretKey:   channel.SecretKey,
		WhiteList:   channel.WhiteList,
		Status:      channel.Status,
		Balance:     channel.Balance,
		CreditLimit: channel.CreditLimit,
		CreatedAt:   channel.CreatedAt,
	}
	return projectDTOs, channelDTO, nil
}

func UpdateProjectStatus(projectID, status int) error {
	project, err := repository.GetProjectByID(int64(projectID))
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

// func productName(projectID, skuID, brandID, specID int) (string, error) {
// 	products, err := repository.GetProductList(&repository.ProjectProductQuery{
// 		ProjectID: &projectID,
// 		SKUID:     &skuID,
// 		BrandID:   &brandID,
// 		SpecID:    &specID,
// 	})
// 	if err != nil {
// 		log.Printf("get product name failed, err: %v", err)
// 		return "", err
// 	}
// 	sku, err := repository.GetSkuByID(skuID)
// 	if err != nil {
// 		log.Printf("get sku by id failed, err: %v", err)
// 		return "", err
// 	}
// 	brand, err := repository.GetBrandByID(brandID)
// 	if err != nil {
// 		log.Printf("get brand by id failed, err: %v", err)
// 		return "", err
// 	}
// 	spec, err := repository.GetSpecByID(specID)
// 	if err != nil {
// 		log.Printf("get spec by id failed, err: %v", err)
// 		return "", err
// 	}

// 	return fmt.Sprintf("%s%s%s-%d", brand.Name, spec.Name, sku.Name, len(products)+1), nil
// }

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

// func GetProjectProductListByProjectID(id int) ([]dto.ProjectProduct, error) {
// 	products, err := repository.GetProductList(&repository.ProjectProductQuery{ProjectID: &id})
// 	if err != nil {
// 		return nil, err
// 	}
// 	var projectProductDTOs []dto.ProjectProduct
// 	for _, product := range products {
// 		projectProductDTOs = append(projectProductDTOs, dto.ProjectProduct{
// 			ID:        product.ID,
// 			Name:      product.Name,
// 			Status:    product.Status,
// 			ProjectID: product.ProjectID,
// 			BrandID:   product.BrandID,
// 			SpecID:    product.SpecID,
// 			SKUID:     product.SKUID,
// 			FacePrice: product.FacePrice,
// 			Price:     product.Price,
// 			CreatedAt: product.CreatedAt,
// 			Version:   product.Version,
// 		})
// 	}
// 	return projectProductDTOs, nil
// }

func UpdateProjectProduct(form *form.UpdateProjectProductForm) error {
	product, err := repository.GetProductByID(int64(form.ID))
	if err != nil {
		log.Printf("get project product by id failed, err: %v", err)
		return err
	}
	product.Status = *form.Status
	product.FacePrice = int(form.FacePrice * 100)
	product.Price = int(form.Price * 100)
	return repository.UpdateProduct(product)
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

// func GetProjectProductList(q *form.ProjectProductQueryForm) (*dto.ProjectProductResp, error) {
// 	project, err := repository.GetProjectByID(*q.ProjectID)
// 	if err != nil {
// 		log.Printf("get project by id failed, err: %v", err)
// 		return nil, err
// 	}
// 	projectProductList, err := repository.GetProductList(&repository.ProjectProductQuery{
// 		ProjectID: q.ProjectID,
// 	})
// 	if err != nil {
// 		log.Printf("get project product list failed, err: %v", err)
// 		return nil, err
// 	}
// 	resp := &dto.ProjectProductResp{
// 		ProjectName: project.Name,
// 		ProductList: make([]dto.ProjectProductItem, 0, len(projectProductList)),
// 	}

// 	for _, product := range projectProductList {
// 		// 获取这个产品的供应商的产品id
// 		relations, err := repository.GetProductRelationList(product.ID)
// 		if err != nil {
// 			log.Printf("get product relation list failed, err: %v", err)
// 			return nil, err
// 		}
// 		supplierProduct, _ := repository.GetSupplierProductByCode(product.SupplierID, product.SupplierProductCode)
// 		dtoProduct := dto.ProjectProduct{
// 			ID:                  product.ID,
// 			Name:                product.Name,
// 			Status:              product.Status,
// 			ProjectID:           product.ProjectID,
// 			BrandID:             product.BrandID,
// 			SpecID:              product.SpecID,
// 			SKUID:               product.SKUID,
// 			FacePrice:           product.FacePrice,
// 			Price:               product.Price,
// 			SupplierID:          product.SupplierID,
// 			SupplierName:        product.SupplierName,
// 			SupplierPrice:       supplierProduct.Price,
// 			SupplierProductID:   product.SupplierProductID,
// 			SupplierProductName: supplierProduct.Name,
// 			Model:               product.Model,
// 			CreatedAt:           product.CreatedAt,
// 			Version:             product.Version,
// 		}
// 		resp.ProductList = append(resp.ProductList, dto.ProjectProductItem{
// 			ProjectProduct:      dtoProduct,
// 			SupplierProductList: make([]dto.RelatedSupplierProduct, 0, len(relations)),
// 		})
// 		for _, relation := range relations {
// 			supplierProduct, err := repository.GetSupplierProductByID(relation.SupplierProductID)
// 			if err != nil {
// 				log.Printf("get supplier product by id failed, err: %v", err)
// 				return nil, err
// 			}
// 			dtoSuProduct := dto.RelatedSupplierProduct{
// 				ID:                  relation.ID,
// 				Name:                supplierProduct.Name,
// 				Code:                supplierProduct.Code,
// 				SupplierID:          supplierProduct.SupplierID,
// 				SupplierName:        supplierProduct.SupplierName,
// 				SupplierCode:        supplierProduct.SupplierCode,
// 				SupplierProductID:   supplierProduct.ID,
// 				SupplierProductName: supplierProduct.Name,
// 				Status:              relation.Status,
// 				BrandID:             supplierProduct.BrandID,
// 				SpecID:              supplierProduct.SpecID,
// 				SKUID:               supplierProduct.SKUID,
// 				FacePrice:           supplierProduct.FacePrice,
// 				Price:               supplierProduct.Price,
// 				CreatedAt:           supplierProduct.CreatedAt,
// 			}
// 			resp.ProductList[len(resp.ProductList)-1].SupplierProductList = append(resp.ProductList[len(resp.ProductList)-1].SupplierProductList, dtoSuProduct)
// 		}
// 	}
// 	return resp, nil
// }

func AddSupplierProduct(form *form.AddSupplierProductForm) error {
	// 检查项目产品是否存在
	_, err := repository.GetProductByID(int64(form.ProjectProductID))
	if err != nil {
		log.Printf("get project product by id failed, err: %v", err)
		return err
	}
	// 检查供应商产品是否存在
	for _, supplierProductID := range form.SupplierProductIDList {
		_, err := repository.GetSupplierProductByID(int64(supplierProductID))
		if err != nil {
			log.Printf("get supplier product by id failed, err: %v", err)
			return err
		}
	}
	addList := make([]repository.ProductSupplier, 0, len(form.SupplierProductIDList))
	for _, supplierProductID := range form.SupplierProductIDList {
		addList = append(addList, repository.ProductSupplier{
			ProductID:         int64(form.ProjectProductID),
			SupplierProductID: int64(supplierProductID),
		})
	}
	if err := repository.CreateProductSupplierList(addList); err != nil {
		log.Printf("create product supplier list failed, err: %v", err)
		return err
	}
	return nil
}

// GetProductSupplierList 获取产品的供应商产品列表
func GetProductSupplierList(productID int64) (*form.ProductSupplierResp, error) {
	product, err := repository.GetProductByID(int64(productID))
	if err != nil {
		return nil, err
	}
	productSupplierList, err := repository.GetProductSupplierList(product.ID)
	if err != nil {
		return nil, err
	}
	channelName, _ := repository.GetChannelNameByID(product.ChannelID)

	projectName, _ := repository.GetProjectNameByID(product.ProjectID)

	resp := &form.ProductSupplierResp{
		ChannelName: channelName,
		ProjectName: projectName,
		ProductName: product.Name,
		Product: form.Product{
			ID:                product.ID,
			Name:              product.Name,
			Status:            product.Status,
			SupplierProductID: product.SupplierProductID,
			// SupplierProductCode: product.SupplierProductCode,
			FacePrice: product.FacePrice,
			Price:     product.Price,
			Model:     product.Model,
		},
		SupplierList: make([]form.ProductSupplier, 0, len(productSupplierList)),
	}
	for _, relation := range productSupplierList {
		supplierProduct, err := repository.GetSupplierProductByID(relation.SupplierProductID)
		if err != nil {
			return nil, err
		}
		resp.SupplierList = append(resp.SupplierList, form.ProductSupplier{
			ID:                  int64(relation.ID),
			SupplierID:          int64(supplierProduct.SupplierID),
			SupplierName:        supplierProduct.SupplierName,
			SupplierCode:        supplierProduct.SupplierCode,
			SupplierProductID:   int64(supplierProduct.ID),
			SupplierProductCode: supplierProduct.Code,
			SupplierProductName: supplierProduct.Name,
			FacePrice:           supplierProduct.FacePrice,
			Price:               supplierProduct.Price,
			Status:              relation.Status,
		})
	}

	sort.Slice(resp.SupplierList, func(i, j int) bool {
		return resp.SupplierList[i].Price < resp.SupplierList[j].Price
	})

	return resp, nil
}
