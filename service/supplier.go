package service

import (
	"fmt"
	"log"
	"sort"
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
			ID:        int64(supplier.ID),
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

// func RechargeSupplier(req *form.RechargeSupplierForm) error {
// 	supplier, err := repository.GetSupplierByID(int64(req.ID))
// 	if err != nil {
// 		log.Printf("get supplier by id failed, err: %v", err)
// 		return err
// 	}
// 	if supplier.Name != req.Name {
// 		return fmt.Errorf("supplier name not match")
// 	}
// 	recharge := &repository.SupplierRecharge{
// 		SupplierID:    supplier.ID,
// 		SupplierName:  supplier.Name,
// 		SupplierCode:  supplier.Code,
// 		Amount:        req.Amount,
// 		ImageURL:      req.ImageURL,
// 		Status:        0, // 0 是审核中
// 		ApplyUserID:   1, // 1 是管理员
// 		ApplyUserName: "admin",
// 		AuditUserID:   0,
// 		AuditUserName: "",
// 		Remark:        nil,
// 	}
// 	return repository.CreateSupplierRecharge(recharge)
// }

func UpdateSupplier(supplier *repository.Supplier) error {
	_, err := repository.GetSupplierByID(supplier.ID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
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

func CreateSupplierProduct(form *form.CreateSupplierProductReq) error {
	supplier, err := repository.GetSupplierByID(int64(form.SupplierID))
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	brandSpecSKU, err := GetOrCreateBrandSpecSKU(form.BrandID, form.SpecID, form.SKUID)
	if err != nil {
		log.Printf("get or create brand spec sku failed, err: %v", err)
		return err
	}
	name := brandSpecSKU.Name

	count, err := repository.GetSupplierProductCount(&repository.SupplierProductQuery{
		SupplierID:     form.SupplierID,
		BrandSpecSKUID: brandSpecSKU.ID,
	})
	if err != nil {
		log.Printf("get supplier product count failed, err: %v", err)
		return err
	}

	if count != 0 {
		name += fmt.Sprintf("-%d", count+1)
	}
	product := &repository.SupplierProduct{
		Name:           name,
		SupplierID:     int64(form.SupplierID),
		Code:           form.Code,
		SupplierName:   supplier.Name,
		SupplierCode:   supplier.Code,
		BrandSpecSKUID: int64(brandSpecSKU.ID),
		Status:         1,
		FacePrice:      int(form.FacePrice * 100),
		Price:          int(form.Price * 100),
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

func SupplierProduct(supplierID int) (*form.SupplierProductResp, error) {
	total, products, err := repository.SupplierProductList(&repository.SupplierProductQuery{
		SupplierID: supplierID,
	})
	if err != nil {
		log.Printf("get supplier product list failed, err: %v", err)
		return nil, err
	}
	resp := form.SupplierProductResp{
		Supplier: form.Supplier{},
		Total:    total,
		Items:    make([]form.SupplierProduct, 0, len(products)),
	}
	if len(products) > 0 {
		resp.Supplier = form.Supplier{
			ID:   int64(products[0].SupplierID),
			Code: products[0].SupplierCode,
			Name: products[0].SupplierName,
		}
	}
	for _, product := range products {
		resp.Items = append(resp.Items, form.SupplierProduct{
			ID:           int64(product.ID),
			Name:         product.Name,
			Code:         product.Code,
			SupplierID:   int64(product.SupplierID),
			SupplierName: product.SupplierName,
			SupplierCode: product.SupplierCode,
			FacePrice:    product.FacePrice,
			Price:        product.Price,
			Status:       product.Status,
			CreatedAt:    product.CreatedAt,
		})
	}
	return &resp, nil
}

func SupplierProductList(req *form.SupplierProductListReq) (*form.SupplierProductListResp, error) {
	var brandSpecSKU repository.BrandSpecSKU
	if req.BrandSpecSKUID != 0 {
		bss, err := repository.GetBrandSpecSKUByID(req.BrandSpecSKUID)
		if err != nil {
			log.Printf("get brand spec sku failed, err: %v", err)
			return nil, err
		}
		brandSpecSKU = *bss
	} else if req.BrandID != 0 && req.SpecID != 0 && req.SKUID != 0 {
		bss, err := repository.GetBrandSpecSKUByIDInfo(req.BrandID, req.SpecID, req.SKUID)
		if err != nil {
			log.Printf("get brand spec sku failed, err: %v", err)
			return nil, err
		}
		brandSpecSKU = *bss
	}
	total, products, err := repository.SupplierProductList(&repository.SupplierProductQuery{
		BrandSpecSKUID: brandSpecSKU.ID,
		Page:           req.Page,
		PageSize:       req.PageSize,
	})
	if err != nil {
		log.Printf("get supplier product list failed, err: %v", err)
		return nil, err
	}
	resp := form.SupplierProductListResp{
		Total: total,
		Items: make([]form.SupplierProduct, 0, len(products)),
	}
	for _, product := range products {
		resp.Items = append(resp.Items, form.SupplierProduct{
			ID:           product.ID,
			Name:         product.Name,
			Code:         product.Code,
			SupplierID:   product.SupplierID,
			SupplierName: product.SupplierName,
			SupplierCode: product.SupplierCode,
			FacePrice:    product.FacePrice,
			Price:        product.Price,
			Status:       product.Status,
			CreatedAt:    product.CreatedAt,
		})
	}
	sort.Slice(resp.Items, func(i, j int) bool {
		return resp.Items[i].Price < resp.Items[j].Price
	})
	return &resp, nil
}
