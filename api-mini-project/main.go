package main

import (
	"api-mini-project/config"
	"api-mini-project/middlewares"
	"api-mini-project/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()

	server := echo.New()

	middlewares.LogMiddleware(server)

	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":9090"))
}
