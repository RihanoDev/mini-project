package response

import (
	"api-mini-project/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint           `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
	Name         string         `json:"name" form:"name" faker:"word"`
	Price        int            `json:"price" form:"price"`
	Description  string         `json:"description" form:"description" faker:"sentence"`
	CategoryName string         `json:"category" form:"category"`
	CategoryID   uint           `json:"category_id" form:"category_id"`
	Stock        uint           `json:"stock" form:"stock"`
}

func FromDomain(domain products.Domain) Product {
	return Product{
		ID:           domain.ID,
		Name:         domain.Name,
		Price:        domain.Price,
		Description:  domain.Description,
		CategoryName: domain.CategoryName,
		CategoryID:   domain.CategoryID,
		Stock:        domain.Stock,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}
