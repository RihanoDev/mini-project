package request

import (
	"api-mini-project/businesses/products"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	Stock       uint   `json:"stock" form:"stock" validate:"required"`
}

func (req *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Stock:       req.Stock,
	}
}

func (req *Product) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
