package model

type ProductInput struct {
	Name        string `json: "name" form: "name"`
	Price       int    `json: "price" form: "price"`
	Description string `json: "description" form: "description"`
	CategoryID  uint   `json: "category_id" form: "category_id"`
	Stock       uint   `json: "stock" form: "stock"`
}
