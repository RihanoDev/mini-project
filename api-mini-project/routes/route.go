package routes

import (
	"api-mini-project/controller"
	"api-mini-project/middlewares"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
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

	privateRoutes.Use(middlewares.CheckTokenMiddleware)

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

	//routes for categories
	// privateRoutes.GET("/api/v1/categories", controller.GetAllCategories)
	// privateRoutes.POST("/api/v1/categories", controller.AddCategory)
	// privateRoutes.PUT("/api/v1/categories/:id", controller.UpdateCategory)
	// privateRoutes.DELETE("/api/v1/categories", controller.DeleteCategory)

	//logout
	privateRoutes.POST("/api/v1/users/logout", controller.Logout)
}
