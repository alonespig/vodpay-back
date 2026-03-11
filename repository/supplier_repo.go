package repository

import (
	"errors"
	"log"
	"vodpay/common"
	"vodpay/database"

	"gorm.io/gorm"
)

func CreateSupplier(supplier *Supplier) error {
	return database.DB.Create(supplier).Error
}

func GetSupplierByID(id int) (*Supplier, error) {
	var supplier Supplier
	if err := database.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrSupplierNotFound
		}
		log.Println("[get supplier by id] err : ", err)
		return nil, common.ErrDBQuery
	}
	return &supplier, nil
}

func SupplierList() ([]Supplier, error) {
	var suppliers []Supplier
	if err := database.DB.Find(&suppliers).Error; err != nil {
		log.Println("[get supplier list] err : ", err)
		return nil, common.ErrDBQuery
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
	return database.DB.Model(&Supplier{}).
		Where("id = ?", supplier.ID).
		Updates(map[string]interface{}{
			"status": supplier.Status,
		}).Error
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

func GetSupplierProductListByID(supplierID int) ([]SupplierProduct, error) {
	var products []SupplierProduct
	if err := database.DB.Where("supplier_id = ?", supplierID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

type SupplierProductQuery struct {
	SupplierID     int
	BrandSpecSKUID int
	Status         *int
	Page           int
	PageSize       int
}

func SupplierProductList(req *SupplierProductQuery) (int64, []SupplierProduct, error) {
	if req == nil {
		req = &SupplierProductQuery{}
	}

	var products []SupplierProduct
	db := database.DB.Model(&SupplierProduct{})

	if req.SupplierID != 0 {
		db = db.Where("supplier_id = ?", req.SupplierID)
	}

	if req.BrandSpecSKUID != 0 {
		db = db.Where("brand_spec_sku_id = ?", req.BrandSpecSKUID)
	}

	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	if req.Page > 0 && req.PageSize > 0 {
		db = db.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize)
	}

	err := db.Find(&products).Error
	return total, products, err
}

func GetSupplierProductCount(req *SupplierProductQuery) (int, error) {
	if req == nil {
		req = &SupplierProductQuery{}
	}

	var count int64
	db := database.DB.Model(&SupplierProduct{})

	log.Println("[supplier product list] req : ", req)

	if req.SupplierID != 0 {
		db = db.Where("supplier_id = ?", req.SupplierID)
	}

	if req.BrandSpecSKUID != 0 {
		db = db.Where("brand_spec_sku_id = ?", req.BrandSpecSKUID)
	}

	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	err := db.Count(&count).Error
	return int(count), err
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
		log.Println("[get spec list] err : ", err)
		return nil, common.ErrDBQuery
	}
	return specs, nil
}

func GetSpecByID(specID int) (*Spec, error) {
	var spec Spec
	if err := database.DB.Where("id = ?", specID).First(&spec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrSpecNotFound
		}
		return nil, common.ErrDBQuery
	}
	return &spec, nil
}

func GetSkuByID(skuID int) (*Sku, error) {
	var sku Sku
	if err := database.DB.Where("id = ?", skuID).First(&sku).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrSkuNotFound
		}
		return nil, common.ErrDBQuery
	}
	return &sku, nil
}

func GetSkuList() ([]Sku, error) {
	var skus []Sku
	if err := database.DB.Find(&skus).Error; err != nil {
		log.Println("[get sku list] err : ", err)
		return nil, common.ErrDBQuery
	}
	return skus, nil
}

func GetBrandByID(brandID int) (*Brand, error) {
	var brand Brand
	if err := database.DB.Where("id = ?", brandID).First(&brand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrBrandNotFound
		}
		return nil, common.ErrDBQuery
	}
	return &brand, nil
}

func GetBrandList() ([]Brand, error) {
	var brands []Brand
	if err := database.DB.Find(&brands).Error; err != nil {
		log.Println("[get brand list] err : ", err)
		return nil, common.ErrDBQuery
	}
	return brands, nil
}

func CreateSupplierProduct(product *SupplierProduct) error {
	return database.DB.Create(product).Error
}

func GetSupplierProductByCode(supplierID int, code string) (SupplierProduct, error) {
	supplierProduct := SupplierProduct{}
	err := database.DB.Model(&supplierProduct).
		Where("supplier_id = ? AND code = ?", supplierID, code).
		First(&supplierProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return SupplierProduct{}, common.ErrSupplierProductNotFound
		}
		log.Printf("[GetSupplierProductByCode] supplierID = %d, code = %s: %v", supplierID, code, err)
		return SupplierProduct{}, common.ErrDBQuery
	}
	return supplierProduct, nil
}

func GetSupplierProductByID(productID int) (*SupplierProduct, error) {
	var supplierProduct SupplierProduct
	err := database.DB.First(&supplierProduct, "id = ?", productID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrSupplierProductNotFound
		}
		log.Printf("[GetSupplierProductByID] id = %d: %v", productID, err)
		return nil, common.ErrDBQuery
	}
	return &supplierProduct, nil
}

func UpdateSupplierProduct(supplierProduct *SupplierProduct) error {
	err := database.DB.Updates(supplierProduct).Error
	if err != nil {
		log.Printf("[UpdateSupplierProduct] id = %d: %v", supplierProduct.ID, err)
		return common.ErrDBQuery
	}
	return nil
}
