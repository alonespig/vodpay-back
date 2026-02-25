package repository

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;AutoIncrement"`
	Name      string    `gorm:"column:name;unique;not null"`
	Password  string    `gorm:"column:password;not null"`
	Status    int       `gorm:"column:status;default:1"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (User) TableName() string {
	return "users"
}
