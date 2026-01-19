package service

import "vodpay/model"

func SupplierList() ([]model.Supplier, error) {
	return model.SupplierList()
}

func CreateSupplier(supplier *model.Supplier) error {
	return model.CreateSupplier(supplier)
}
