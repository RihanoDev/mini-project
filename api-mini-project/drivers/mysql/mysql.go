package mysql_driver

import (
	"api-mini-project/drivers/mysql/categories"
	"api-mini-project/drivers/mysql/products"
	"api-mini-project/drivers/mysql/users"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&products.Product{}, &categories.Category{}, &users.User{})
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("Database connection is closed")

	return nil
}

func SeedCategory(db *gorm.DB) categories.Category {
	var category categories.Category = categories.Category{
		Name: "sample",
	}

	if err := db.Create(&category).Error; err != nil {
		panic(err)
	}

	var createdCategory categories.Category

	db.Last(&createdCategory)

	return createdCategory
}

func SeedProduct(db *gorm.DB) products.Product {
	category := SeedCategory(db)

	var product products.Product = products.Product{
		Name:        "test",
		Price:       123,
		Description: "test desc",
		CategoryID:  category.ID,
		Stock:       1000,
	}

	if err := db.Create(&product).Error; err != nil {
		panic(err)
	}

	var createdProduct products.Product

	db.Last(&createdProduct)

	return createdProduct
}

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user users.User = users.User{
		Email:    "test@gmail.com",
		Password: string(password),
	}

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser users.User

	db.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func CleanSeeders(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	categoryResult := db.Exec("DELETE FROM categories")
	itemResult := db.Exec("DELETE FROM products")
	userResult := db.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}
	log.Println("Seeders are cleaned up successfully")
}
