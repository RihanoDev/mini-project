package routes

import (
	"api-mini-project/controller"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// func New() *echo.Echo {
// 	e := echo.New()

// 	//routing
// 	e.POST("/Register", controller.RegisterUser)

// 	return e
// }

func SetupRoute(server *echo.Echo) {
	//routes for auth
	server.POST("/api/v1/users/register", controller.Register)
	server.POST("/api/v1/users/login", controller.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))
	//private route
	// privateRoutes.GET("/test", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "authenticated!")
	// })

	//routes for products
	privateRoutes.GET("/api/v1/products", controller.GetAll)
	privateRoutes.GET("/api/v1/products/:id", controller.GetByID)
	privateRoutes.POST("/api/v1/products", controller.Create)
	privateRoutes.PUT("/api/v1/products/:id", controller.Update)
	privateRoutes.DELETE("/api/v1/products/:id", controller.Delete)
	privateRoutes.POST("/api/v1/products/:id", controller.Restore)
	privateRoutes.DELETE("/api/v1/products/force/:id", controller.ForceDelete)
}
