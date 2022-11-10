package routes

import (
	"api-mini-project/app/middlewares"
	"api-mini-project/controllers/categories"
	"api-mini-project/controllers/products"
	"api-mini-project/controllers/users"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	CategoryController categories.CategoryController
	ProductController  products.ProductController
}

func (cl *ControllerList) Route(server *echo.Echo) {
	server.Use(cl.LoggerMiddleware)

	users := server.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	product := server.Group("/api/v1/products", middleware.JWTWithConfig(cl.JWTMiddleware))
	product.Use(middlewares.CheckTokenMiddleware)

	product.GET("", cl.ProductController.GetAllProducts)
	product.GET("/:id", cl.ProductController.GetProductByID)
	product.POST("", cl.ProductController.CreateProduct)
	product.PUT("/:id", cl.ProductController.UpdateProduct)
	product.DELETE("/:id", cl.ProductController.DeleteProduct)
	product.POST("/:id", cl.ProductController.RestoreProduct)
	product.DELETE("/force/:id", cl.ProductController.ForceDeleteProduct)

	category := server.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))
	category.Use(middlewares.CheckTokenMiddleware)

	category.GET("", cl.CategoryController.GetAllCategories)
	category.GET("/:id", cl.CategoryController.GetCategoriesByID)
	category.POST("", cl.CategoryController.CreateCategories)
	category.PUT("/:id", cl.CategoryController.UpdateCategories)
	category.DELETE("/:id", cl.CategoryController.DeleteCategories)

	auth := server.Group("api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

}
