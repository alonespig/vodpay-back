package service

import (
	"fmt"
	"log"
	"vodpay/form"
	"vodpay/model"
)

func SupplierList() ([]model.Supplier, error) {
	return model.SupplierList()
}

func CreateSupplier(supplier *model.Supplier) error {
	return model.CreateSupplier(supplier)
}

func RechargeSupplier(req *form.RechargeSupplierForm) error {
	supplier, err := model.GetSupplierByID(req.SupplierID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	if supplier.Name != req.SupplierName {
		return fmt.Errorf("supplier name not match")
	}
	return model.RechargeSupplier(req.SupplierID, req.Amount)
}

func UpdateSupplier(supplier *model.Supplier) error {
	oldSupplier, err := model.GetSupplierByID(supplier.ID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	if oldSupplier.Name != supplier.Name {
		return fmt.Errorf("supplier name not match")
	}
	return model.UpdateSupplierStatus(supplier)
}

func UpdateSupplierProduct(supplierProduct *model.SupplierProduct) error {
	oldProduct, err := model.GetSupplierProductByID(supplierProduct.ID)
	if err != nil {
		log.Printf("get supplier product by id failed, err: %v", err)
		return err
	}
	if oldProduct.Code != supplierProduct.Code {
		return fmt.Errorf("supplier product code not match")
	}
	return model.UpdateSupplierProduct(supplierProduct)
}

func matchSupProductName(spupplierID, skuID, brandID, specID int) (string, error) {
	total, err := model.SupplierProductName(spupplierID, skuID, brandID, specID)
	if err != nil {
		log.Printf("get supplier product name failed, err: %v", err)
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
	name := fmt.Sprintf("%s%s%s", brand.Name, spec.Name, sku.Name)

	if total != 0 {
		name += fmt.Sprintf("%d", total+1)
	}
	return name, nil
}

func CreateSupplierProduct(form *form.SupplierProduct) error {
	supplier, err := model.GetSupplierByID(form.SupplierID)
	if err != nil {
		log.Printf("get supplier by id failed, err: %v", err)
		return err
	}
	name, err := matchSupProductName(form.SupplierID, form.SKUID, form.BrandID, form.SpecID)
	if err != nil {
		log.Printf("match supplier product name failed, err: %v", err)
		return err
	}
	product := &model.SupplierProduct{
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
	err = model.CreateSupplierProduct(product)
	if err != nil {
		log.Printf("create supplier product failed, err: %v", err)
		return err
	}
	return nil
}

func SupplierProductList() ([]model.SupplierProduct, error) {
	return model.SupplierProductList()
}

func CreateModel(modelName string, name string) error {
	return model.CreateModel(modelName, name)
}

func GetModelByID(modelName string, id int) (*model.BaseModel, error) {
	return model.GetModelByID(modelName, id)
}

func GetModelList(modelName string) ([]model.BaseModel, error) {
	return model.GetModelList(modelName)
}
