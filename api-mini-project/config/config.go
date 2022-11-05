package config

import (
	"api-mini-project/model"
	"api-mini-project/util"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME  string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD  string = util.GetConfig("DB_PASSWORD")
	DB_NAME      string = util.GetConfig("DB_NAME")
	DB_HOST      string = util.GetConfig("DB_HOST")
	DB_PORT      string = util.GetConfig("DB_PORT")
	DB_TEST_NAME string = util.GetConfig("DB_TEST_NAME")
)

func InitDB() {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}
	log.Println("connected to the database")
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.Product{}, &model.Category{}, &model.User{})
}

func InitTestDB() {
	var err error

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_TEST_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}
	log.Println("connected to the database")
	initMigrateTest()
}

func initMigrateTest() {
	DB.AutoMigrate(&model.Product{}, &model.Category{}, &model.User{})
}

func SeedCategory() model.Category {
	var category model.Category = model.Category{
		Name: "sample",
	}

	if err := DB.Create(&category).Error; err != nil {
		panic(err)
	}

	var createdCategory model.Category

	DB.Last(&createdCategory)

	return createdCategory
}

func SeedProduct() model.Product {
	category := SeedCategory()

	var product model.Product = model.Product{
		Name:        "test",
		Price:       123,
		Description: "test desc",
		CategoryID:  category.ID,
		Stock:       1000,
	}

	if err := DB.Create(&product).Error; err != nil {
		panic(err)
	}

	var createdProduct model.Product

	DB.Last(&createdProduct)

	return createdProduct
}

func SeedUser() model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user model.User = model.User{
		Email:    "test@gmail.com",
		Password: string(password),
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser model.User

	DB.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func CleanSeeders() {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	categoryResult := DB.Exec("DELETE FROM categories")
	itemResult := DB.Exec("DELETE FROM products")
	userResult := DB.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}
	log.Println("Seeders are cleaned up successfully")
}
