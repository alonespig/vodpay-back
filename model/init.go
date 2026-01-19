package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() {
	dsn := `root:Root123!@tcp(127.0.0.1:3306)/vodpay_db?charset=utf8mb4&parseTime=True&loc=Local`
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败 err = ", err)
	}
	log.Println("数据库连接成功")
}
