package products

import (
	"api-mini-project/businesses/products"

	"gorm.io/gorm"
)

type productRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) products.Repository {
	return &productRepository{
		conn: conn,
	}
}

func (pr *productRepository) GetAll() []products.Domain {
	var rec []Product

	pr.conn.Preload("Category").Find(&rec)

	productDomain := []products.Domain{}

	for _, product := range rec {
		productDomain = append(productDomain, product.ToDomain())
	}

	return productDomain
}

func (pr *productRepository) GetByID(id string) products.Domain {
	var product Product

	pr.conn.Preload("Category").First(&product, "id = ?", id)

	return product.ToDomain()
}

func (pr *productRepository) Create(productDomain *products.Domain) products.Domain {
	rec := FromDomain(productDomain)

	result := pr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (pr *productRepository) Update(id string, productDomain *products.Domain) products.Domain {
	var product products.Domain = pr.GetByID(id)

	updateProduct := FromDomain(&product)

	updateProduct.Name = productDomain.Name
	updateProduct.Description = productDomain.Description
	updateProduct.Price = productDomain.Price
	updateProduct.CategoryID = productDomain.CategoryID
	updateProduct.Stock = productDomain.Stock

	pr.conn.Save(&updateProduct)

	return updateProduct.ToDomain()
}

func (pr *productRepository) Delete(id string) bool {
	var product products.Domain = pr.GetByID(id)

	deleteProduct := FromDomain(&product)

	result := pr.conn.Delete(&deleteProduct)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (pr *productRepository) Restore(id string) products.Domain {
	var trashProduct products.Domain

	trashed := FromDomain(&trashProduct)

	pr.conn.Unscoped().First(&trashed, "id = ?", id)

	trashed.DeletedAt = gorm.DeletedAt{}

	pr.conn.Unscoped().Save(&trashed)

	return trashed.ToDomain()
}

func (pr *productRepository) ForceDelete(id string) bool {
	var product products.Domain = pr.GetByID(id)

	deleteProduct := FromDomain(&product)

	result := pr.conn.Unscoped().Delete(&deleteProduct)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
