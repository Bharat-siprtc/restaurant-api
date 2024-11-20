package controller

import (
	"example/restaurant-api/manager"
	"example/restaurant-api/request"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	menu, err := manager.GetMenu(flag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, menu)
}
func GetMenuById(c echo.Context) error {
	flag := c.QueryParam("flag")
	id := c.Param("id")
	res, err := manager.GetMenuById(id,flag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("error in getting data: %v", err)})
	}
	return c.JSON(http.StatusOK, res)
}
func CreateMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	var menu request.CreateRequest
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.CreateMenu(flag,menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "menu created successfully.")
}
func DeleteMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	id := c.Param("id")

	if err := manager.DeleteMenu(flag,id); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"sucess": "food item with given id is deleted successfully"})
}
func UpdateMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	id := c.Param("id")
	var req request.UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.UpdateMenu(flag,id, req); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"sucess": fmt.Sprintf("food item with id %v is updated successfully", id)})
}
func GetMenuPg(c echo.Context) error {
	page := c.QueryParam("page")
	size := c.QueryParam("size")
	if page != "" && size != "" {
		menu, err := manager.GetMenuPage(page, size)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, menu)
	}
	menu, err := manager.GetMenu(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// return c.JSON(http.StatusOK, menu)
	return c.JSON(http.StatusOK, menu)
}
// func GetMenuByIdPg(c echo.Context) error {
// 	id := c.Param("id")
// 	res, err := manager.GetMenuByIdPg(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("error in getting data: %v", err)})
// 	}
// 	return c.JSON(http.StatusOK, res)
// }
func CreateMenuPg(c echo.Context) error {
	var menu request.CreateRequest
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.CreateMenuPg(menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "menu created successfully.")
}
func UpdateMenuPg(c echo.Context) error {
	id := c.Param("id")
	var menu request.UpdateRequest
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := manager.UpdateMenuPg(id, menu); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "sucessfully updated")
}
func DeleteMenuPg(c echo.Context) error {
	id := c.Param("id")
	if err := manager.DeleteMenuPg(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("error in delete:%v", err)})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": fmt.Sprintf("menuitem with id %v deleted successfull.", id)})
}
