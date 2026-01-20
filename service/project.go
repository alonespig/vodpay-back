package service

import "vodpay/model"

func CreateModel(modelName string, name string) error {
	return model.CreateModel(modelName, name)
}

func GetModelByID(modelName string, id int) (*model.BaseModel, error) {
	return model.GetModelByID(modelName, id)
}

func GetModelList(modelName string) ([]model.BaseModel, error) {
	return model.GetModelList(modelName)
}
