package users

import (
	"api-mini-project/app/middlewares"
	"api-mini-project/businesses/users"
	"api-mini-project/controllers/users/request"
	"api-mini-project/controllers/users/response"
	"net/http"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation failed",
		})
	}

	sameEmail := ctrl.authUseCase.CheckData(userInput.ToDomain()).Email

	if userInput.Email == sameEmail {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email Already Used",
		})
	}

	user := ctrl.authUseCase.Register(userInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomain(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}

	token := ctrl.authUseCase.Login(userInput.ToDomain())

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid Email or Password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ctrl *AuthController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid token",
		})
	}

	middlewares.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout Success",
	})
}
