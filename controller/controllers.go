package controller

import (
	"fmt"
	"net/http"
	"restaurant-api/manager"
	"restaurant-api/request"

	"github.com/labstack/echo/v4"
)

func GetMenu(c echo.Context) error {
	menu, err := manager.GetMenu()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, menu)
}
func CreateMovie(c echo.Context) error {
	var menu request.CreateRequest
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.CreateMovie(menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "movie created successfully.")
}
func DeleteMovie(c echo.Context) error {
	id := c.Param("id")

	if err := manager.DeleteMovie(id); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"sucess": "food item with given id is deleted successfully"})
}
func UpdateMovie(c echo.Context) error {
	id := c.Param("id")
	var req request.UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.UpdateMovie(id, req); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"sucess": fmt.Sprintf("food item with id %v is updated successfully", id)})
}
func GetMenuPg(c echo.Context) error {
	menu, err := manager.GetMenuPg()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, menu)
}
