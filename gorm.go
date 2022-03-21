package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func mainGorm() {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DB_NAME := "go_videos"
	LINK := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DB_NAME)

	// connection
	db, err := gorm.Open(mysql.Open(LINK), &gorm.Config{})
	if err != nil {
		panic("unable to connect")
	}

	// migrate schema
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Code:  "A12",
		Price: 99000,
	})
	db.Create(&Product{
		Code:  "B98",
		Price: 150000,
	})

	// read
	var product1 Product
	var product2 Product
	db.First(&product1, 1)                 // 1 is primary key
	db.First(&product2, "code = ?", "B98") // find with code = B98
	log.Println(product1)
	log.Println(product2)

	// update
	db.Model(&product2).Update("Price", 200000)
	db.Model(&product1).Updates(Product{
		Price: 300000,
		Code:  "C55",
	})
	// or
	db.Model(&product2).Updates(map[string]interface{}{
		"Price": 300000,
		"Code":  "C54",
	})

	// delete
	db.Delete(&product2, 2)
}
