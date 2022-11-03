package model

type ProductInput struct {
	Name        string `json: "name" form: "name"`
	Price       int    `json: "price" form: "pricec"`
	Description string `json: "description" form: "description"`
}
