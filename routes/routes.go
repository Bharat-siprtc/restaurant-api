package routes

import (
	"example/restaurant-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/menu", controller.GetMenu)
	e.GET("/menus", controller.GetMenuPg)
	e.POST("/menu", controller.CreateMovie)
	e.DELETE("/menu/:id", controller.DeleteMovie)
	e.PUT("/menu/:id", controller.UpdateMovie)
}
