package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
