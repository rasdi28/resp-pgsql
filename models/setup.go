package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dns := "host=localhost user=postgres password=postgress dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
