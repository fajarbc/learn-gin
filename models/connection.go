package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DB_NAME := "go_articles"
	LINK := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(LINK))
	if err != nil {
		panic(err.Error())
	}

	return db
}

func AutoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Author{}, &Article{})
}
