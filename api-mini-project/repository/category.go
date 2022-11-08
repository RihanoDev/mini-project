package repository

import (
	"api-mini-project/config"
	"api-mini-project/model"
)

type CategoryRepositoryImpl struct{}

func (c *CategoryRepositoryImpl) GetAllCategories() []model.Category {
	var categories []model.Category

	config.DB.Find(&categories)

	return categories
}

func (c *CategoryRepositoryImpl) GetCategoriesByID(id string) model.Category {
	var category model.Category

	config.DB.First(&category, "id = ?", id)

	return category
}

func (c *CategoryRepositoryImpl) CreateCategories(input model.CategoryInput) model.Category {
	var newCategory model.Category = model.Category{
		Name: input.Name,
	}

	var addedCategory model.Category = model.Category{}

	result := config.DB.Create(&newCategory)

	result.Last(&addedCategory)

	return addedCategory
}

func (c *CategoryRepositoryImpl) UpdateCategories(id string, input model.CategoryInput) model.Category {
	var category model.Category = c.GetCategoriesByID(id)

	category.Name = input.Name

	config.DB.Save(&category)

	return category
}

func (c *CategoryRepositoryImpl) ForceDeleteCategories(id string) bool {
	var category model.Category = c.GetCategoriesByID(id)

	result := config.DB.Delete(&category)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (c *CategoryRepositoryImpl) CheckDataCategories(input model.CategoryInput) model.Category {
	var data model.Category = model.Category{}

	config.DB.Find(&data, "name = ?", input.Name)

	return data
}
