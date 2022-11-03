package routes

import (
	"api-mini-project/controller"

	"github.com/labstack/echo"
)

// func New() *echo.Echo {
// 	e := echo.New()

// 	//routing
// 	e.POST("/Register", controller.RegisterUser)

// 	return e
// }

func SetupRoute(server *echo.Echo) {
	server.GET("/api/v1/products", controller.GetAll)
	server.GET("/api/v1/products/:id", controller.GetByID)
	server.POST("/api/v1/products", controller.Create)
	server.PUT("/api/v1/products/:id", controller.Update)
	server.DELETE("/api/v1/products/:id", controller.Delete)
	server.POST("/api/v1/products/:id", controller.Restore)
	server.DELETE("/api/v1/products/force/:id", controller.ForceDelete)
}
