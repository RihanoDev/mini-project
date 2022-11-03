package service

import (
	"api-mini-project/model"
	"api-mini-project/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func New() ProductService {
	return ProductService{
		Repository: &repository.ProductRepositoryImpl{},
	}
}

func (n *ProductService) GetAll() []model.Product {
	return n.Repository.GetAll()
}

func (n *ProductService) GetByID(id string) model.Product {
	return n.Repository.GetByID(id)
}

func (n *ProductService) Create(input model.ProductInput) model.Product {
	return n.Repository.Create(input)
}

func (n *ProductService) Update(id string, input model.ProductInput) model.Product {
	return n.Repository.Update(id, input)
}

func (n *ProductService) Delete(id string) bool {
	return n.Repository.Delete(id)
}

func (n *ProductService) Restore(id string) model.Product {
	return n.Repository.Restore(id)
}

func (n *ProductService) ForceDelete(id string) bool {
	return n.Repository.ForceDelete(id)
}
