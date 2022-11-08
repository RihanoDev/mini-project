package service

import (
	"api-mini-project/model"
	"api-mini-project/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func NewCategoryService() CategoryService {
	return CategoryService{
		Repository: &repository.CategoryRepositoryImpl{},
	}
}

func (c *CategoryService) GetAllCategories() []model.Category {
	return c.Repository.GetAllCategories()
}

func (c *CategoryService) GetCategoriesByID(id string) model.Category {
	return c.Repository.GetCategoriesByID(id)
}

func (c *CategoryService) CreateCategories(input model.CategoryInput) model.Category {
	return c.Repository.CreateCategories(input)
}

func (c *CategoryService) UpdateCategories(id string, input model.CategoryInput) model.Category {
	return c.Repository.UpdateCategories(id, input)
}

func (c *CategoryService) ForceDeleteCategories(id string) bool {
	return c.Repository.ForceDeleteCategories(id)
}

func (c *CategoryService) CheckDataCategories(input model.CategoryInput) model.Category {
	return c.Repository.CheckDataCategories(input)
}
