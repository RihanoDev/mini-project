package route

import (
	"rest/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//routing
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.POST("/user", controller.RegisterController)

	return e
}
