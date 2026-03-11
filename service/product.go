package service

import (
	"strconv"
	"vodpay/database"
	"vodpay/form"
	"vodpay/repository"
)

func UpdateProductRelation(form *form.UpdateProductRelationForm) error {
	relation, err := repository.GetProductRelationByID(form.ID)
	if err != nil {
		return err
	}
	product, err := repository.GetProductByID(relation.ChannelProductID)
	if err != nil {
		return err
	}
	if product.Status == 0 {
		return ErrProductStatusDisabled
	}
	if *form.Status == 0 && product.SupplierProductID == relation.SupplierProductID {
		return ErrSupplierProductUsing
	}
	return repository.UpdateProductRelationStatus(form.ID, *form.Status)
}

func UpdateProduct(form *form.UpdateProductForm) error {
	product, err := repository.GetProductByID(form.ID)
	if err != nil {
		return err
	}
	if form.SupplierProductID != 0 && product.SupplierProductID != form.SupplierProductID {
		supplierProduct, err := repository.GetSupplierProductByID(form.SupplierProductID)
		if err != nil {
			return err
		}
		relation, err := repository.GetProductRelationByProductID(product.ID, form.SupplierProductID)
		if err != nil {
			return ErrProductNotFoundSupplierProduct
		}
		if relation.Status == 0 {
			return ErrProductRelationStatusDisabled
		}
		supplier, err := repository.GetSupplierByID(supplierProduct.SupplierID)
		if err != nil {
			return err
		}
		product.SupplierID = supplier.ID
		product.SupplierName = supplier.Name
		product.SupplierProductID = form.SupplierProductID
		product.SupplierProductCode = supplierProduct.Code
	}
	product.FacePrice = int(form.FacePrice * 100)
	product.Price = int(form.Price * 100)
	product.Status = *form.Status
	product.Model = *form.Model

	// 更新数据库
	if err := repository.UpdateProduct2(product); err != nil {
		return err
	}
	return nil
}

func GetProductList(q *form.ProductListQueryForm) (*form.ProductListQueryResp, error) {
	products, err := repository.GetProductListByProjectID(*q.ProjectID)
	if err != nil {
		return nil, err
	}
	project, err := repository.GetProjectByID(*q.ProjectID)
	if err != nil {
		return nil, err
	}
	resp := form.ProductListQueryResp{
		ProjectName: project.Name,
		ProductList: make([]form.Product, 0, len(products)),
	}
	for _, product := range products {
		resp.ProductList = append(resp.ProductList, form.Product{
			ID:                  product.ID,
			Name:                product.Name,
			Status:              product.Status,
			SupplierName:        product.SupplierName,
			SupplierPrice:       product.FacePrice,
			SupplierProductName: product.SupplierProductCode,
			SupplierProductID:   product.SupplierProductID,
			FacePrice:           product.FacePrice,
			Price:               product.Price,
			Model:               product.Model,
			CreatedAt:           product.CreatedAt,
		})
	}
	return &resp, nil
}

func CreateProduct(form *form.CreateProductForm) error {
	brandSpecSKU, err := repository.GetBrandSpecSKUByIDInfo(
		form.BrandID, form.SpecID, form.SKUID)

	if err != nil {
		return err
	}

	project, err := repository.GetProjectByID(form.ProjectID)
	if err != nil {
		return err
	}
	channel, err := repository.GetChannelByID(project.ChannelID)
	if err != nil {
		return err
	}

	count, err := repository.GetProductCount(form.ProjectID,
		brandSpecSKU.ID)
	if err != nil {
		return err
	}

	productName := brandSpecSKU.Name
	if count > 0 {
		productName = productName + "-" + strconv.Itoa(count+1)
	}

	product := &repository.Product{
		Name:           productName,
		ChannelID:      channel.ID,
		ProjectID:      form.ProjectID,
		LimitCount:     form.LimitNum,
		BrandSpecSKUID: brandSpecSKU.ID,
		FacePrice:      int(form.FacePrice * 100),
		Price:          int(form.Price * 100),
		Model:          *form.Model,
		Status:         1,
	}

	// 创建数据库
	if _, err = repository.CreateProduct(product); err != nil {
		return err
	}

	relations := make([]repository.ProductRelation, 0, len(form.SupplierProductIDList))
	for index, supplierProductID := range form.SupplierProductIDList {
		supplierProduct, err := repository.GetSupplierProductByID(supplierProductID)
		if err != nil {
			return err
		}
		relations = append(relations, repository.ProductRelation{
			ChannelProductID:  product.ID,
			SupplierProductID: supplierProduct.ID,
			Status:            1,
		})
		if index == 0 {
			product.SupplierID = supplierProduct.SupplierID
			product.SupplierName = supplierProduct.SupplierName
			product.SupplierProductID = supplierProduct.ID
			product.SupplierProductCode = supplierProduct.Code
		}
	}

	// 创建数据库
	if err = repository.CreateProductRelation(relations); err != nil {
		return err
	}

	if err = database.DB.Model(product).Updates(product).Error; err != nil {
		return err
	}

	return nil
}
