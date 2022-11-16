package model

type Product struct {
	Id    int    `json: "id" form: "id"`
	Name  string `json: "name" form: "name`
	Price int    `json: "price" form: "price"`
}
