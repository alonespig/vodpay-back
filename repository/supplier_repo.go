package repository

import (
	"errors"
	"vodpay/database"
	"vodpay/form"

	"gorm.io/gorm"
)

func CreateSupplier(supplier *Supplier) error {
	return database.DB.Create(supplier).Error
}

func GetSupplierByID(id int) (*Supplier, error) {
	var supplier Supplier
	if err := database.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSupplierNotFound
		}
		return nil, err
	}
	return &supplier, nil
}

func SupplierList() ([]Supplier, error) {
	var suppliers []Supplier
	if err := database.DB.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func GetSupplierByCode(code string) (*Supplier, error) {
	var supplier Supplier
	if err := database.DB.Where("code = ?", code).First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSupplierNotFound
		}
		return nil, err
	}
	return &supplier, nil
}

func UpdateSupplier(supplier *Supplier) error {
	return database.DB.Updates(supplier).Error
}

func CreateSupplierRecharge(recharge *SupplierRecharge) error {
	return database.DB.Create(recharge).Error
}

type SupplierRechargeQuery struct {
	SupplierID *int
	Status     *int
}

func GetSupplierRechargeList(query *SupplierRechargeQuery) ([]SupplierRecharge, error) {
	var recharge []SupplierRecharge
	db := database.DB.Model(&SupplierRecharge{})

	if query == nil {
		query = &SupplierRechargeQuery{}
	}

	if query.SupplierID != nil {
		db = db.Where("supplier_id = ?", *query.SupplierID)
	}
	if query.Status != nil {
		db = db.Where("status >= ?", *query.Status)
	}

	if err := db.Find(&recharge).Error; err != nil {
		return nil, err
	}
	return recharge, nil
}

func GetSupplierRechargeByID(id int) (*SupplierRecharge, error) {
	var recharge SupplierRecharge
	if err := database.DB.Where("id = ?", id).First(&recharge).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSupplierRechargeNotFound
		}
		return nil, err
	}
	return &recharge, nil
}

func UpdateSupplierRecharge(recharge *SupplierRecharge) error {
	return database.DB.Updates(recharge).Error
}

func SupplierProductList(req *form.SupplierProductReq) ([]SupplierProduct, error) {
	if req == nil {
		req = &form.SupplierProductReq{}
	}

	var products []SupplierProduct
	db := database.DB

	if req.SupplierID != 0 {
		db = db.Where("supplier_id = ?", req.SupplierID)
	}
	if req.SpecID != 0 {
		db = db.Where("spec_id = ?", req.SpecID)
	}
	if req.SKUID != 0 {
		db = db.Where("sku_id = ?", req.SKUID)
	}
	if req.BrandID != 0 {
		db = db.Where("brand_id = ?", req.BrandID)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	err := db.Find(&products).Error
	return products, err
}

func CreateBrand(brand *Brand) error {
	return database.DB.Create(brand).Error
}

func CreateSpec(spec *Spec) error {
	return database.DB.Create(spec).Error
}

func CreateSku(sku *Sku) error {
	return database.DB.Create(sku).Error
}

func GetSpecList() ([]Spec, error) {
	var specs []Spec
	if err := database.DB.Find(&specs).Error; err != nil {
		return nil, err
	}
	return specs, nil
}

func GetSpecByID(specID int) (*Spec, error) {
	var spec Spec
	if err := database.DB.Where("id = ?", specID).First(&spec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSpecNotFound
		}
		return nil, err
	}
	return &spec, nil
}

func GetSkuByID(skuID int) (*Sku, error) {
	var sku Sku
	if err := database.DB.Where("id = ?", skuID).First(&sku).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSkuNotFound
		}
		return nil, err
	}
	return &sku, nil
}

func GetSkuList() ([]Sku, error) {
	var skus []Sku
	if err := database.DB.Find(&skus).Error; err != nil {
		return nil, err
	}
	return skus, nil
}

func GetBrandByID(brandID int) (*Brand, error) {
	var brand Brand
	if err := database.DB.Where("id = ?", brandID).First(&brand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBrandNotFound
		}
		return nil, err
	}
	return &brand, nil
}

func GetBrandList() ([]Brand, error) {
	var brands []Brand
	if err := database.DB.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}

func CreateSupplierProduct(product *SupplierProduct) error {
	return database.DB.Create(product).Error
}

func GetSupplierProductByID(productID int) (*SupplierProduct, error) {
	var supplierProduct SupplierProduct
	err := database.DB.First(&supplierProduct, "id = ?", productID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSupplierProductNotFound
		}
		return nil, err
	}
	return &supplierProduct, nil
}

func UpdateSupplierProduct(supplierProduct *SupplierProduct) error {
	return database.DB.Updates(supplierProduct).Error
}
