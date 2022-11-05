package service

import (
	"api-mini-project/model"
	"api-mini-project/repository"
)

type AuthService struct {
	Repository repository.AuthRepository
}

func NewAuthService() AuthService {
	return AuthService{
		Repository: &repository.AuthRepositoryImpl{},
	}
}

func (a *AuthService) Register(input model.UserInput) model.User {
	return a.Repository.Register(input)
}

func (a *AuthService) Login(input model.UserInput) string {
	return a.Repository.Login(input)
}

func (a *AuthService) CheckData(input model.UserInput) model.User {
	return a.Repository.CheckData(input)
}
