package repository

import (
	"api-mini-project/config"
	"api-mini-project/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{}

func (p *ProductRepositoryImpl) GetAll() []model.Product {
	var products []model.Product

	config.DB.Preload("Category").Find(&products)

	return products
}

func (b *ProductRepositoryImpl) GetByID(id string) model.Product {
	var product model.Product

	config.DB.Preload("Category").First(&product, "id = ?", id)

	return product
}

func (c *ProductRepositoryImpl) Create(input model.ProductInput) model.Product {
	var newProduct model.Product = model.Product{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		Stock:       input.Stock,
	}

	var addedProduct model.Product = model.Product{}

	result := config.DB.Create(&newProduct)

	result.Last(&addedProduct)

	return addedProduct
}

func (u *ProductRepositoryImpl) Update(id string, input model.ProductInput) model.Product {
	var product model.Product = u.GetByID(id)

	product.Name = input.Name
	product.Price = input.Price
	product.Description = input.Description
	product.CategoryID = input.CategoryID
	product.Stock = input.Stock

	config.DB.Save(&product)

	return product
}

func (d *ProductRepositoryImpl) Delete(id string) bool {
	var product model.Product = d.GetByID(id)

	result := config.DB.Delete(&product)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (r *ProductRepositoryImpl) Restore(id string) model.Product {
	var trashedProduct model.Product

	config.DB.Unscoped().Preload("Category").First(&trashedProduct, "id = ?", id)

	trashedProduct.DeletedAt = gorm.DeletedAt{}

	config.DB.Unscoped().Save(&trashedProduct)

	return trashedProduct
}

func (f *ProductRepositoryImpl) ForceDelete(id string) bool {
	var product model.Product = f.GetByID(id)

	result := config.DB.Unscoped().Delete(&product)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
