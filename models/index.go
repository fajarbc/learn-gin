package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseLink() string {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	NAME := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, NAME)
}
func ConnectDatabase() *gorm.DB {
	LINK := GetDatabaseLink()
	db, err := gorm.Open(mysql.Open(LINK))
	if err != nil {
		panic(err.Error())
	}

	return db
}

func AutoMigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&Author{}, &Article{})
}
