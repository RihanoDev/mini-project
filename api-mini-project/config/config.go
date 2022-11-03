package config

import (
	"api-mini-project/model"
	"api-mini-project/util"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = util.GetConfig("DB_PASSWORD")
	DB_NAME     string = util.GetConfig("DB_NAME")
	DB_HOST     string = util.GetConfig("DB_HOST")
	DB_PORT     string = util.GetConfig("DB_PORT")
)

func InitDB() {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}
	initMigrate()
	log.Println("connected to the database")
}

func initMigrate() {
	DB.AutoMigrate(&model.Product{})
}
