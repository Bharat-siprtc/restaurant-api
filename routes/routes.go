package routes

import (
	"example/restaurant-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/menu", controller.GetMenu)
	e.GET("/menu/pg", controller.GetMenuPg)
	e.GET("/menu/pg/:id", controller.GetMenuByIdPg)
	e.GET("/menu/pg", controller.GetMenuPg)
	e.POST("/menu/pg", controller.CreateMenuPg)
	e.PUT("/menu/pg/:id", controller.UpdateMenuPg)
	e.DELETE("/menu/pg/:id", controller.DeleteMenuPg)
	e.POST("/menu", controller.CreateMenu)
	e.DELETE("/menu/:id", controller.DeleteMenu)
	e.PUT("/menu/:id", controller.UpdateMenu)
}
