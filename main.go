package main

import (
	"example/restaurant-api/config"
	"example/restaurant-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.ConnectDB()
	config.CreateCounterSeq()
	config.PostgresConnect()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8282"))
}
