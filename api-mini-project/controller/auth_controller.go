package controller

import (
	"api-mini-project/auth"
	"api-mini-project/model"
	"api-mini-project/service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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

	registeredEmail := authService.CheckData(*userInput).Email

	if userInput.Email == registeredEmail {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email Already Used",
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

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := auth.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid Token",
		})
	}

	auth.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout Success",
	})
}
