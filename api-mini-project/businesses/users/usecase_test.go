package users_test

import (
	"api-mini-project/app/middlewares"
	"api-mini-project/businesses/users"
	_userMock "api-mini-project/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJWT{})

	usersDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123",
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		usersRepository.On("Register", &usersDomain).Return(usersDomain).Once()

		result := usersService.Register(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Register | InValid", func(t *testing.T) {
		usersRepository.On("Register", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Register(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}

func TestCheckData(t *testing.T) {
	t.Run("CheckData | Valid", func(t *testing.T) {
		usersRepository.On("CheckData", &usersDomain).Return(usersDomain).Once()

		result := usersService.CheckData(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("CheckData | InValid", func(t *testing.T) {
		usersRepository.On("CheckData", &usersDomain).Return(usersDomain).Once()

		result := usersService.CheckData(&usersDomain)

		assert.NotNil(t, result)
	})
}
