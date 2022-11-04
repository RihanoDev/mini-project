package model

import "github.com/go-playground/validator/v10"

type ProductInput struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	Stock       uint   `json:"stock" form:"stock" validate:"required"`
}

func (input *ProductInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
