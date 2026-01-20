package model

import (
	"fmt"
	"time"
)

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
