package model

import (
	"fmt"
	"log"
	"time"
)

// 供应商
type Supplier struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Code      string    `db:"code" json:"code"`
	Balance   int       `db:"balance" json:"balance"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type SupplierProduct struct {
	ID           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Code         string    `db:"code" json:"code"`
	SupplierID   int       `db:"supplier_id" json:"supplierID"`
	SupplierName string    `db:"supplier_name" json:"supplierName"`
	SupplierCode string    `db:"supplier_code" json:"supplierCode"`
	FacePrice    int       `db:"face_price" json:"facePrice"`
	SpecID       int       `db:"spec_id" json:"specID"`
	SKUID        int       `db:"sku_id" json:"skuID"`
	BrandID      int       `db:"brand_id" json:"brandID"`
	Price        int       `db:"price" json:"price"`
	Status       int       `db:"status" json:"status"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
}

type SupplierRecharge struct {
	ID            int        `db:"id" json:"id"`
	SupplierID    int        `db:"supplier_id" json:"supplierID"`
	SupplierName  string     `db:"supplier_name" json:"supplierName"`
	SupplierCode  string     `db:"supplier_code" json:"supplierCode"`
	Amount        int        `db:"amount" json:"amount"`
	Status        int        `db:"status" json:"status"`
	ApplyUserID   int        `db:"apply_user_id" json:"applyUserID"`
	ApplyUserName string     `db:"apply_user_name" json:"applyUserName"`
	AuditUserID   int        `db:"audit_user_id" json:"auditUserID"`
	AuditUserName string     `db:"audit_user_name" json:"auditUserName"`
	ImageURL      string     `db:"image_url" json:"imageURL"`
	Remark        *string    `db:"remark" json:"remark"`
	PassAt        *time.Time `db:"pass_at" json:"passAt"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
}

type BaseModel struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func CreateModel(modelName string, name string) error {
	sqlStr := fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", modelName)
	_, err := db.Exec(sqlStr, name)
	if err != nil {
		return err
	}
	return nil
}

func GetModelByID(modelName string, id int) (*BaseModel, error) {
	var data BaseModel
	sqlStr := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", modelName)
	err := db.Get(&data, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetModelList(modelName string) ([]BaseModel, error) {
	var data []BaseModel
	sqlStr := fmt.Sprintf("SELECT * FROM %s", modelName)
	err := db.Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateSupplier(supplier *Supplier) error {
	sqlStr := "INSERT INTO suppliers (name, code) VALUES (:name, :code)"
	_, err := db.NamedExec(sqlStr, supplier)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplierStatus(supplier *Supplier) error {
	sqlStr := "UPDATE suppliers SET status = :status WHERE id = :id"
	_, err := db.NamedExec(sqlStr, supplier)
	if err != nil {
		return err
	}
	return nil
}

func CreateSupplierRecharge(recharge *SupplierRecharge) error {
	recharge.Amount = recharge.Amount * 100
	sqlStr := `INSERT INTO supplier_recharges (supplier_id, supplier_name, supplier_code, amount, status, apply_user_id, 
								apply_user_name, audit_user_id, audit_user_name, image_url) 
				VALUES (:supplier_id, :supplier_name, :supplier_code, :amount, :status, 
				:apply_user_id, :apply_user_name, :audit_user_id, :audit_user_name, :image_url)`
	_, err := db.NamedExec(sqlStr, recharge)
	if err != nil {
		log.Printf("create supplier recharge failed, err: %v", err)
		return err
	}
	return nil
}

func GetSupplierRechargeList(status int) ([]SupplierRecharge, error) {
	var recharges []SupplierRecharge
	sqlStr := fmt.Sprintf("SELECT * FROM supplier_recharges WHERE status = %d order by id ", status)
	err := db.Select(&recharges, sqlStr)
	if err != nil {
		return nil, err
	}
	return recharges, nil
}

func GetSupplierRechargeHistoryList() ([]SupplierRecharge, error) {
	var recharges []SupplierRecharge
	sqlStr := "SELECT * FROM supplier_recharges WHERE status != 1 order by id desc"
	err := db.Select(&recharges, sqlStr)
	if err != nil {
		return nil, err
	}
	return recharges, nil
}

func GetSupplierRechargeByID(id int) (*SupplierRecharge, error) {
	var recharge SupplierRecharge
	sqlStr := fmt.Sprintf("SELECT * FROM supplier_recharges WHERE id = %d", id)
	err := db.Get(&recharge, sqlStr)
	if err != nil {
		return nil, err
	}
	return &recharge, nil
}

func UpdateSupplierRecharge(recharge *SupplierRecharge) error {
	sqlStr := `UPDATE supplier_recharges SET 
				status = :status, pass_at = :pass_at, audit_user_id = :audit_user_id, audit_user_name = :audit_user_name, remark = :remark 
				WHERE id = :id`
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = db.NamedExec(sqlStr, recharge)
	if err != nil {
		log.Printf("update supplier recharge failed, err: %v", err)
		return err
	}
	if recharge.Status != 2 {
		// 提交事务
		err = tx.Commit()
		if err != nil {
			log.Printf("commit supplier recharge failed, err: %v", err)
			return err
		}
		return nil
	}
	// 更新供应商余额
	sqlStr = "UPDATE suppliers SET balance = balance + ? WHERE id = ?"
	_, err = tx.Exec(sqlStr, recharge.Amount*100, recharge.SupplierID)
	if err != nil {
		log.Printf("update supplier balance failed, err: %v", err)
		return err
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Printf("commit supplier recharge failed, err: %v", err)
		return err
	}
	return nil
}

// func RechargeSupplier(supplierID int, amount int) error {
// 	// 检查供应商是否存在
// 	supplier, err := GetSupplierByID(supplierID)
// 	if err != nil {
// 		log.Printf("get supplier by id failed, err: %v", err)
// 		return err
// 	}
// 	if supplier == nil {
// 		return fmt.Errorf("supplier not found")
// 	}
// 	sqlStr := "UPDATE suppliers SET balance = balance + ? WHERE id = ?"
// 	_, err = db.Exec(sqlStr, amount*100, supplierID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func SupplierList() ([]Supplier, error) {
	var suppliers []Supplier
	err := db.Select(&suppliers, "SELECT * FROM suppliers")
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func GetSupplierByID(supplierID int) (*Supplier, error) {
	var supplier Supplier
	sqlStr := "SELECT * FROM suppliers WHERE id = ?"
	err := db.Get(&supplier, sqlStr, supplierID)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func CreateSupplierProduct(supplierProduct *SupplierProduct) error {
	log.Printf("create supplier product, %v", supplierProduct)
	sqlStr := `INSERT INTO supplier_products (name, code, supplier_id, supplier_name, supplier_code, face_price, price, status, spec_id, sku_id, brand_id) 
				VALUES (:name, :code, :supplier_id, :supplier_name, :supplier_code, :face_price, :price, :status, :spec_id, :sku_id, :brand_id)`
	_, err := db.NamedExec(sqlStr, supplierProduct)
	if err != nil {
		return err
	}
	return nil
}

func GetSupplierProductByID(productID int) (*SupplierProduct, error) {
	var supplierProduct SupplierProduct
	sqlStr := "SELECT * FROM supplier_products WHERE id = ?"
	err := db.Get(&supplierProduct, sqlStr, productID)
	if err != nil {
		return nil, err
	}
	return &supplierProduct, nil
}

func UpdateSupplierProduct(supplierProduct *SupplierProduct) error {
	sqlStr := `UPDATE supplier_products SET 
				status = :status, price = :price, face_price = :face_price 
				WHERE id = :id`
	_, err := db.NamedExec(sqlStr, supplierProduct)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplierProductStatus(supplierProduct *SupplierProduct) error {
	sqlStr := "UPDATE supplier_products SET status = :status WHERE id = :id"
	_, err := db.NamedExec(sqlStr, supplierProduct)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplierProductPrice(supplierProduct *SupplierProduct) error {
	sqlStr := "UPDATE supplier_products SET price = :price WHERE id = :id"
	_, err := db.NamedExec(sqlStr, supplierProduct)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplierProductFacePrice(supplierProduct *SupplierProduct) error {
	sqlStr := "UPDATE supplier_products SET face_price = :face_price WHERE id = :id"
	_, err := db.NamedExec(sqlStr, supplierProduct)
	if err != nil {
		return err
	}
	return nil
}

func SupplierProductListByInfo(supplierID, specID, skuID, brandID int) ([]SupplierProduct, error) {
	var supplierProducts []SupplierProduct
	sqlStr := "SELECT * FROM supplier_products WHERE 1=1 "
	if supplierID > 0 {
		sqlStr += fmt.Sprintf("AND supplier_id = %d ", supplierID)
	}
	if specID > 0 {
		sqlStr += fmt.Sprintf("AND spec_id = %d ", specID)
	}
	if skuID > 0 {
		sqlStr += fmt.Sprintf("AND sku_id = %d ", skuID)
	}
	if brandID > 0 {
		sqlStr += fmt.Sprintf("AND brand_id = %d ", brandID)
	}
	log.Printf("supplier product list by info, sql: %s", sqlStr)
	err := db.Select(&supplierProducts, sqlStr)
	if err != nil {
		return nil, err
	}
	return supplierProducts, nil
}

func SupplierProductListBySupplierID(supplierID int) ([]SupplierProduct, error) {
	var supplierProducts []SupplierProduct
	sqlStr := fmt.Sprintf("SELECT * FROM supplier_products WHERE supplier_id = %d", supplierID)
	err := db.Select(&supplierProducts, sqlStr)
	if err != nil {
		return nil, err
	}
	return supplierProducts, nil
}

func SupplierProductList() ([]SupplierProduct, error) {
	var supplierProducts []SupplierProduct
	sqlStr := "SELECT * FROM supplier_products"
	err := db.Select(&supplierProducts, sqlStr)
	if err != nil {
		return nil, err
	}
	return supplierProducts, nil
}

func SupplierProductName(spupplierID, skuID, brandID, specID int) (int64, error) {
	var total int64
	sqlStr := `SELECT COUNT(*) FROM supplier_products 
				WHERE supplier_id = ? AND sku_id = ? AND brand_id = ? AND spec_id = ?`
	err := db.Get(&total, sqlStr, spupplierID, skuID, brandID, specID)
	if err != nil {
		return 0, err
	}
	return total, nil
}
