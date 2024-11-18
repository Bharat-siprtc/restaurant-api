package main

import (
	"restaurant-api/config"
	"restaurant-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.ConnectDB()
	config.PostgresConnect()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8282"))
}
