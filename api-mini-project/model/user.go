package model

import (
	"time"
)

type User struct {
	ID          int       `json: "id" form: "id" gorm: "primary-key"`
	Name        string    `json: "name" form: "name"`
	Age         int       `json: "age" form: "age"`
	Gender      string    `json: "gender" form: "gender"`
	Email       string    `json: "email" form: "email"`
	Password    string    `json: "password" form: "password"`
	NoHandphone string    `json: "no_hp" form: "no_hp"`
	CreatedAt   time.Time `json: "created_at" form: "created_at"`
	UpdatedAt   time.Time `json: "updated_at" form: "updated_at"`
}
