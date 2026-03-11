package service

import (
	"errors"
	"fmt"
	"log"
	"vodpay/form"
	"vodpay/repository"

	"gorm.io/gorm"
)

// 如果不存在，就创建 brand_spec_sku 记录
func GetOrCreateBrandSpecSKU(brandID, specID, skuID int) (*repository.BrandSpecSKU, error) {
	BrandSpecSKU, err := repository.GetBrandSpecSKUByIDInfo(brandID, specID, skuID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return createBrandSpecSKU(brandID, specID, skuID)
		}
		return nil, err
	}
	return BrandSpecSKU, nil
}

func createBrandSpecSKU(brandID, specID, skuID int) (*repository.BrandSpecSKU, error) {
	brand, err := repository.GetBrandByID(brandID)
	if err != nil {
		return nil, err
	}
	spec, err := repository.GetSpecByID(specID)
	if err != nil {
		return nil, err
	}
	sku, err := repository.GetSkuByID(skuID)
	if err != nil {
		return nil, err
	}
	return repository.CreateBrandSpecSKU(&repository.BrandSpecSKU{
		BrandID: brandID,
		SpecID:  specID,
		SKUID:   skuID,
		Name:    brand.Name + spec.Name + sku.Name,
	})
}

func GetList() (*form.BaseListResp, error) {
	list, err := repository.GetBrandSpecSKUList()
	if err != nil {
		return nil, err
	}
	resp := &form.BaseListResp{
		Total:         int64(len(list)),
		BaseModelList: make([]form.BaseModel, len(list)),
	}
	for i, brandSpecSKU := range list {
		resp.BaseModelList[i] = form.BaseModel{
			ID:   brandSpecSKU.ID,
			Name: brandSpecSKU.Name,
		}
	}
	return resp, nil
}

func CreateBaseBSS(bssType string, name string) error {
	switch bssType {
	case "BRAND":
		return repository.CreateBrand(&repository.Brand{BaseModel: repository.BaseModel{Name: name}})
	case "SPEC":
		return repository.CreateSpec(&repository.Spec{BaseModel: repository.BaseModel{Name: name}})
	case "SKU":
		return repository.CreateSku(&repository.Sku{BaseModel: repository.BaseModel{Name: name}})
	default:
		return fmt.Errorf("invalid bss type")
	}
}

func GetBrandList() ([]form.BaseModel, error) {
	brands, err := repository.GetBrandList()
	if err != nil {
		log.Printf("get brand list failed, err: %v", err)
		return nil, err
	}
	resp := make([]form.BaseModel, 0, len(brands))
	for _, brand := range brands {
		resp = append(resp, form.BaseModel{
			ID:        brand.ID,
			Name:      brand.Name,
			CreatedAt: brand.CreatedAt,
		})
	}
	return resp, nil
}
func GetSpecList() ([]form.BaseModel, error) {
	specs, err := repository.GetSpecList()
	if err != nil {
		log.Printf("get spec list failed, err: %v", err)
		return nil, err
	}
	resp := make([]form.BaseModel, 0, len(specs))
	for _, spec := range specs {
		resp = append(resp, form.BaseModel{
			ID:        spec.ID,
			Name:      spec.Name,
			CreatedAt: spec.CreatedAt,
		})
	}
	return resp, nil
}
func GetSkuList() ([]form.BaseModel, error) {
	skus, err := repository.GetSkuList()
	if err != nil {
		log.Printf("get sku list failed, err: %v", err)
		return nil, err
	}
	resp := make([]form.BaseModel, 0, len(skus))
	for _, sku := range skus {
		resp = append(resp, form.BaseModel{
			ID:        sku.ID,
			Name:      sku.Name,
			CreatedAt: sku.CreatedAt,
		})
	}
	return resp, nil
}
