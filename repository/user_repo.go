package repository

import (
	"errors"
	"vodpay/database"

	"gorm.io/gorm"
)

func CreateUser(user *User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(name string) (*User, error) {
	var user User
	err := database.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
