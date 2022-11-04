package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" form:"id" gorm:"primary-key"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
}
