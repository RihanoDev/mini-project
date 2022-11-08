package model

import "github.com/go-playground/validator/v10"

type CategoryInput struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (input *CategoryInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
