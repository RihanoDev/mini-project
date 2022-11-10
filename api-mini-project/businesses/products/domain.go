package products

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Name         string
	Price        int
	Description  string
	CategoryName string
	CategoryID   uint
	Stock        uint
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(productDomain *Domain) Domain
	Update(id string, productDomain *Domain) Domain
	Delete(id string) bool
	Restore(id string) Domain
	ForceDelete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(productDomain *Domain) Domain
	Update(id string, productDomain *Domain) Domain
	Delete(id string) bool
	Restore(id string) Domain
	ForceDelete(id string) bool
}
