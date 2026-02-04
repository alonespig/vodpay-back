package view

import "vodpay/model"

type ChannelSupplierProductView struct {
	ProjectProduct  model.ProjectProduct    `json:"projectProduct"`
	SupplierProduct []model.SupplierProduct `json:"supplierProduct"`
}
