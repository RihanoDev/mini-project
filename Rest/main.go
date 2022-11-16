package main

import (
	"rest/config"
	"rest/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
