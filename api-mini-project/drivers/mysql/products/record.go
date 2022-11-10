package products

import (
	"api-mini-project/businesses/products"
	"api-mini-project/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint                `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt   time.Time           `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" form:"deleted_at"`
	Name        string              `json:"name" form:"name" faker:"word"`
	Price       int                 `json:"price" form:"price"`
	Description string              `json:"description" form:"description" faker:"sentence"`
	Category    categories.Category `json:"category" form:"category"`
	CategoryID  uint                `json:"category_id" form:"category_id"`
	Stock       uint                `json:"stock" form:"stock"`
}

func FromDomain(domain *products.Domain) *Product {
	return &Product{
		ID:          domain.ID,
		Name:        domain.Name,
		Price:       domain.Price,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Stock:       domain.Stock,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (rec *Product) ToDomain() products.Domain {
	return products.Domain{
		ID:           rec.ID,
		Name:         rec.Name,
		Price:        rec.Price,
		Description:  rec.Description,
		CategoryName: rec.Category.Name,
		CategoryID:   rec.Category.ID,
		Stock:        rec.Stock,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
