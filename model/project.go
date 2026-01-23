package model

import (
	"time"
)

type Project struct {
	ID        int       `db:"id" json:"id"`
	ChannelID int       `db:"channel_id" json:"channelID"`
	Name      string    `db:"name" json:"name"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type ProjectProduct struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Status    int       `db:"status" json:"status"`
	ProjectID int       `db:"project_id" json:"projectID"`
	BrandID   int       `db:"brand_id" json:"brandID"`
	SpecID    int       `db:"spec_id" json:"specID"`
	SKUID     int       `db:"sku_id" json:"skuID"`
	FacePrice int       `db:"face_price" json:"facePrice"`
	Price     int       `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	Version   int       `db:"version" json:"version"`
}

func CreateProject(project *Project) error {
	sqlStr := `INSERT INTO projects (channel_id, name, status)
	 VALUES (:channel_id, :name, :status)`
	_, err := db.NamedExec(sqlStr, project)
	return err
}
func GetProjectByID(id int) (*Project, error) {
	project := &Project{}
	sqlStr := `SELECT id, channel_id, name, status, created_at
	 FROM projects WHERE id = ?`
	err := db.Get(project, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func GetProjectListByChannelID(id int) ([]Project, error) {
	projects := []Project{}
	sqlStr := `SELECT id, channel_id, name, status, created_at
	 FROM projects WHERE channel_id = ?`
	err := db.Select(&projects, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func UpdateProjectStatus(status, id int) error {
	sqlStr := `UPDATE projects SET status = ? WHERE id = ?`
	_, err := db.Exec(sqlStr, status, id)
	return err
}

func CreateProjectProduct(product *ProjectProduct) error {
	sqlStr := `INSERT INTO project_products 
			(name, status, project_id, brand_id, spec_id, sku_id, face_price, price) 
			VALUES (:name, :status, :project_id, :brand_id, :spec_id, :sku_id, :face_price, :price)`
	_, err := db.NamedExec(sqlStr, product)
	if err != nil {
		return err
	}
	return nil
}

func GetProjectProductListByProjectID(id int) ([]ProjectProduct, error) {
	products := []ProjectProduct{}
	sqlStr := `SELECT id, name, status, project_id, brand_id, spec_id, sku_id, face_price, price, created_at, version
	 FROM project_products WHERE project_id = ?`
	err := db.Select(&products, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func ProjectProductName(projectID, skuID, brandID, specID int) (int64, error) {
	var total int64
	sqlStr := `SELECT COUNT(*) FROM project_products 
				WHERE project_id = ? AND sku_id = ? AND brand_id = ? AND spec_id = ?`
	err := db.Get(&total, sqlStr, projectID, skuID, brandID, specID)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetProjectProductByID(projectProductID int) (*ProjectProduct, error) {
	product := &ProjectProduct{}
	sqlStr := `SELECT id, name, status, project_id, brand_id, spec_id, sku_id, face_price, price, created_at, version
	 FROM project_products WHERE id = ?`
	err := db.Get(product, sqlStr, projectProductID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProjectProduct(product *ProjectProduct) error {
	sqlStr := `UPDATE project_products SET status = :status, face_price = :face_price, price = :price, version = version + 1
	 WHERE id = :id AND version = :version`
	_, err := db.NamedExec(sqlStr, product)
	if err != nil {
		return err
	}
	return nil
}
