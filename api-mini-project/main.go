package main

import (
	"api-mini-project/config"
	"api-mini-project/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	server := echo.New()
	routes.SetupRoute(server)
	server.Logger.Fatal(server.Start(":9090"))
}
