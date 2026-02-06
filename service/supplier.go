package service

import (
	"fmt"
	"log"
	"vodpay/dto"
	"vodpay/form"
	"vodpay/repository"
)

func SupplierList() ([]form.SupplierResp, error) {
	suppliers, err := repository.SupplierList()
	if err != nil {
		return nil, err
	}
	resp := make([]form.SupplierResp, 0, len(suppliers))
	for _, supplier := range suppliers {
		resp = append(resp, form.SupplierResp{
			ID:        supplier.ID,
			Name:      supplier.Name,
			Code:      supplier.Code,
			Balance:   supplier.Balance,
			Status:    supplier.Status,
			CreatedAt: supplier.CreatedAt,
		})
	}
	return resp, nil
}

func CreateSupplier(supplier *repository.Supplier) error {
	return repository.CreateSupplier(supplier)
}

func RechargeSupplier(req *form.RechargeSupplierForm) error {
	supplier, err := repository.GetSupplierByID(req.ID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	if supplier.Name != req.Name {
		return fmt.Errorf("supplier name not match")
	}
	recharge := &repository.SupplierRecharge{
		SupplierID:    supplier.ID,
		SupplierName:  supplier.Name,
		SupplierCode:  supplier.Code,
		Amount:        req.Amount,
		ImageURL:      req.ImageURL,
		Status:        0, // 0 是审核中
		ApplyUserID:   1, // 1 是管理员
		ApplyUserName: "admin",
		AuditUserID:   0,
		AuditUserName: "",
		Remark:        nil,
	}
	return repository.CreateSupplierRecharge(recharge)
}

func UpdateSupplier(supplier *repository.Supplier) error {
	oldSupplier, err := repository.GetSupplierByID(supplier.ID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	if oldSupplier.Name != supplier.Name {
		return fmt.Errorf("supplier name not match")
	}
	if oldSupplier.Status != 0 {
		return fmt.Errorf("充值已审核过，请勿重复操作")
	}
	return repository.UpdateSupplier(supplier)
}

func UpdateSupplierProduct(supplierProduct *repository.SupplierProduct) error {
	oldProduct, err := repository.GetSupplierProductByID(supplierProduct.ID)
	if err != nil {
		log.Printf("get supplier product by id failed, err: %v", err)
		return err
	}
	if oldProduct.Code != supplierProduct.Code {
		return fmt.Errorf("supplier product code not match")
	}
	return repository.UpdateSupplierProduct(supplierProduct)
}

func matchSupProductName(spupplierID, skuID, brandID, specID int) (string, error) {
	req := &form.SupplierProductReq{
		SupplierID: spupplierID,
		SKUID:      skuID,
		BrandID:    brandID,
		SpecID:     specID,
	}
	products, err := repository.SupplierProductList(req)
	if err != nil {
		log.Printf("get supplier product name failed, err: %v", err)
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
	name := fmt.Sprintf("%s%s%s", brand.Name, spec.Name, sku.Name)

	if len(products) != 0 {
		name += fmt.Sprintf("%d", len(products)+1)
	}
	return name, nil
}

func CreateSupplierProduct(form *form.SupplierProduct) error {
	supplier, err := repository.GetSupplierByID(form.SupplierID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	name, err := matchSupProductName(form.SupplierID, form.SKUID, form.BrandID, form.SpecID)
	if err != nil {
		log.Printf("match supplier product name failed, err: %v", err)
		return err
	}
	product := &repository.SupplierProduct{
		Name:         name,
		SupplierID:   form.SupplierID,
		Code:         form.Code,
		SupplierName: supplier.Name,
		SupplierCode: supplier.Code,
		Status:       1,
		FacePrice:    int(form.FacePrice * 100),
		Price:        int(form.Price * 100),
		SpecID:       form.SpecID,
		SKUID:        form.SKUID,
		BrandID:      form.BrandID,
	}
	err = repository.CreateSupplierProduct(product)
	if err != nil {
		log.Printf("create supplier product failed, err: %v", err)
		return err
	}
	return nil
}

func GetSupplierRechargeHistoryList() ([]dto.SupplierRecharge, error) {
	status := 1
	rechargeList, err := repository.GetSupplierRechargeList(&repository.SupplierRechargeQuery{
		Status: &status,
	})
	if err != nil {
		log.Printf("get supplier recharge list failed, err: %v", err)
		return nil, err
	}
	resp := make([]dto.SupplierRecharge, 0, len(rechargeList))
	for _, recharge := range rechargeList {
		resp = append(resp, dto.SupplierRecharge{
			ID:            recharge.ID,
			SupplierName:  recharge.SupplierName,
			SupplierCode:  recharge.SupplierCode,
			Amount:        recharge.Amount,
			Status:        recharge.Status,
			ApplyUserName: recharge.ApplyUserName,
			AuditUserName: recharge.AuditUserName,
			ImageURL:      recharge.ImageURL,
			Remark:        recharge.Remark,
			PassAt:        recharge.PassAt,
			CreatedAt:     recharge.CreatedAt,
		})
	}
	return resp, nil
}

func SupplierProductList(req *form.SupplierProductReq) ([]dto.SupplierProduct, error) {
	products, err := repository.SupplierProductList(req)
	if err != nil {
		log.Printf("get supplier product list failed, err: %v", err)
		return nil, err
	}
	resp := make([]dto.SupplierProduct, 0, len(products))
	for _, product := range products {
		resp = append(resp, dto.SupplierProduct{
			ID:           product.ID,
			Name:         product.Name,
			Code:         product.Code,
			SupplierID:   product.SupplierID,
			SupplierName: product.SupplierName,
			SupplierCode: product.SupplierCode,
			FacePrice:    product.FacePrice,
			SpecID:       product.SpecID,
			SKUID:        product.SKUID,
			BrandID:      product.BrandID,
			Price:        product.Price,
			Status:       product.Status,
			CreatedAt:    product.CreatedAt,
		})
	}
	return resp, nil
}

func CreateModel(modelName string, name string) error {
	if modelName == "brands" {
		return repository.CreateBrand(&repository.Brand{BaseModel: repository.BaseModel{Name: name}})
	}
	if modelName == "specs" {
		return repository.CreateSpec(&repository.Spec{BaseModel: repository.BaseModel{Name: name}})
	}
	if modelName == "skus" {
		return repository.CreateSku(&repository.Sku{BaseModel: repository.BaseModel{Name: name}})
	}
	return fmt.Errorf("invalid model name")
}

func GetBrandList() ([]dto.Brand, error) {
	brands, err := repository.GetBrandList()
	if err != nil {
		log.Printf("get brand list failed, err: %v", err)
		return nil, err
	}
	resp := make([]dto.Brand, 0, len(brands))
	for _, brand := range brands {
		resp = append(resp, dto.Brand{
			ID:        brand.ID,
			Name:      brand.Name,
			CreatedAt: brand.CreatedAt,
		})
	}
	return resp, nil
}
func GetSpecList() ([]dto.Spec, error) {
	specs, err := repository.GetSpecList()
	if err != nil {
		log.Printf("get spec list failed, err: %v", err)
		return nil, err
	}
	resp := make([]dto.Spec, 0, len(specs))
	for _, spec := range specs {
		resp = append(resp, dto.Spec{
			ID:        spec.ID,
			Name:      spec.Name,
			CreatedAt: spec.CreatedAt,
		})
	}
	return resp, nil
}
func GetSkuList() ([]dto.Sku, error) {
	skus, err := repository.GetSkuList()
	if err != nil {
		log.Printf("get sku list failed, err: %v", err)
		return nil, err
	}
	resp := make([]dto.Sku, 0, len(skus))
	for _, sku := range skus {
		resp = append(resp, dto.Sku{
			ID:        sku.ID,
			Name:      sku.Name,
			CreatedAt: sku.CreatedAt,
		})
	}
	return resp, nil
}
