package drivers

import (
	categoryDomain "api-mini-project/businesses/categories"
	productDomain "api-mini-project/businesses/products"
	userDomain "api-mini-project/businesses/users"
	categoryDB "api-mini-project/drivers/mysql/categories"
	productDB "api-mini-project/drivers/mysql/products"
	userDB "api-mini-project/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewProductRepository(conn *gorm.DB) productDomain.Repository {
	return productDB.NewMySQLRepository(conn)
}
func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
