package model

import "time"

// 供应商
type Supplier struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	SupplierCode string    `db:"supplier_code" json:"supplier_code"`
	Status       int       `db:"status" json:"status"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

func CreateSupplier(supplier *Supplier) error {
	sqlStr := "INSERT INTO suppliers (name, supplier_code) VALUES (:name, :supplier_code)"
	_, err := db.NamedExec(sqlStr, supplier)
	if err != nil {
		return err
	}
	return nil
}

func SupplierList() ([]Supplier, error) {
	var suppliers []Supplier
	err := db.Select(&suppliers, "SELECT * FROM suppliers")
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}
