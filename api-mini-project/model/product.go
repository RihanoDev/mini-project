package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json: "id" form: "id" gorm: "primaryKey"`
	CreatedAt   time.Time      `json: "created_at" form: "created_at"`
	UpdatedAt   time.Time      `json: "updated_at" form: "updated_at"`
	DeletedAt   gorm.DeletedAt `json: "deleted_at" form: "deleted_at"`
	Name        string         `json: "name" form: "name"`
	Price       int            `json: "price" form: "price"`
	Description string         `json: "description" form: "description"`
}
