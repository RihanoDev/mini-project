package repository

import (
	"api-mini-project/auth"
	"api-mini-project/config"
	"api-mini-project/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct{}

func (r *AuthRepositoryImpl) Register(input model.UserInput) model.User {
	// Make Password parse
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser model.User = model.User{
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser model.User = model.User{}

	result := config.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (l *AuthRepositoryImpl) Login(input model.UserInput) string {
	var user model.User = model.User{}

	config.DB.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	//Generate JWT Token
	token := auth.CreateToken(user.ID)

	return token
}

func (c *AuthRepositoryImpl) CheckData(input model.UserInput) model.User {
	var data model.User = model.User{}

	config.DB.Find(&data, "email = ?", input.Email)

	return data
}
