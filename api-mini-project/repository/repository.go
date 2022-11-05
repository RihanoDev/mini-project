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
