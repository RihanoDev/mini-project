package categories

import (
	"api-mini-project/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll() []categories.Domain {
	var rec []Category

	cr.conn.Find(&rec)

	categoryDomain := []categories.Domain{}

	for _, category := range rec {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}

	return categoryDomain
}

func (cr *categoryRepository) GetByID(id string) categories.Domain {
	var category Category

	cr.conn.First(&category, "id = ?", id)

	return category.ToDomain()
}

func (cr *categoryRepository) Create(categoryDomain *categories.Domain) categories.Domain {
	rec := FromDomain(categoryDomain)

	result := cr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (cr *categoryRepository) Update(id string, categoryDomain *categories.Domain) categories.Domain {
	var category categories.Domain = cr.GetByID(id)

	updateCategory := FromDomain(&category)

	updateCategory.Name = categoryDomain.Name

	cr.conn.Save(&updateCategory)

	return updateCategory.ToDomain()
}

func (cr *categoryRepository) Delete(id string) bool {
	var category categories.Domain = cr.GetByID(id)

	deleteCategory := FromDomain(&category)

	result := cr.conn.Unscoped().Delete(&deleteCategory)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (cr *categoryRepository) CheckData(categoryDomain *categories.Domain) categories.Domain {
	var data Category

	cr.conn.Find(&data, "name = ?", categoryDomain.Name)

	return data.ToDomain()
}
