package repository

import (
	"api-mini-project/model"
)

type ProductRepository interface {
	GetAll() []model.Product
	GetByID(id string) model.Product
	Create(input model.ProductInput) model.Product
	Update(id string, input model.ProductInput) model.Product
	Delete(id string) bool
	Restore(id string) model.Product
	ForceDelete(id string) bool
}

type AuthRepository interface {
	Register(input model.UserInput) model.User
	Login(input model.UserInput) string
	CheckData(input model.UserInput) model.User
}

type CategoryRepository interface {
	GetAllCategories() []model.Category
	GetCategoriesByID(id string) model.Category
	CreateCategories(input model.CategoryInput) model.Category
	UpdateCategories(id string, input model.CategoryInput) model.Category
	ForceDeleteCategories(id string) bool
	CheckDataCategories(input model.CategoryInput) model.Category
}
