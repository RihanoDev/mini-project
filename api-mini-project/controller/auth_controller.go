package controller

import (
	"api-mini-project/model"
	"api-mini-project/service"
	"net/http"

	"github.com/labstack/echo"
)

var authService service.AuthService = service.NewAuthService()

func Register(c echo.Context) error {
	var userInput *model.UserInput = new(model.UserInput)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messaage": "Invalid Request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	user := authService.Register(*userInput)

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var userInput *model.UserInput = new(model.UserInput)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messaage": "Invalid Request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := authService.Login(*userInput)

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid Email or Password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
