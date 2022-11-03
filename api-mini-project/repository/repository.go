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
