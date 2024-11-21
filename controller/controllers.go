package controller

import (
	"example/restaurant-api/manager"
	"example/restaurant-api/request"
	"example/restaurant-api/response"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateStruct(input interface{}) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return fmt.Errorf("validation failed!:%v", err)
	}
	return nil
}
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
	res, err := manager.GetMenuById(id, flag)
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
	if err := ValidateStruct(menu); err != nil {
		return fmt.Errorf("error in request data :%v",err)
	}
	res,err := manager.CreateMenu(flag, menu)
	if  err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,map[string]response.AllMenuResponse{"success":res})
}
func DeleteMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	id := c.Param("id")

	if err := manager.DeleteMenu(flag, id); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"sucess": fmt.Sprintf("menu item with id %v deleted successfully.",id)})
}
func UpdateMenu(c echo.Context) error {
	flag := c.QueryParam("flag")
	id := c.Param("id")
	var req request.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ValidateStruct(req); err != nil {
		return fmt.Errorf("error in request data :%v",err)
	}
	res,err := manager.UpdateMenu(flag, id, req)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]response.AllMenuResponse{"sucess":res})
}
// func GetMenuPg(c echo.Context) error {
// 	page := c.QueryParam("page")
// 	size := c.QueryParam("size")
// 	if page != "" && size != "" {
// 		menu, err := manager.GetMenuPage(page, size)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, menu)
// 	}
// 	menu, err := manager.GetMenu(page)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	// return c.JSON(http.StatusOK, menu)
// 	return c.JSON(http.StatusOK, menu)
// }

//	func GetMenuByIdPg(c echo.Context) error {
//		id := c.Param("id")
//		res, err := manager.GetMenuByIdPg(id)
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("error in getting data: %v", err)})
//		}
//		return c.JSON(http.StatusOK, res)
//	}
// func CreateMenuPg(c echo.Context) error {
// 	var menu request.CreateRequest
// 	if err := c.Bind(&menu); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	if res,err := manager.CreateMenuPg(menu); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, "menu created successfully.")
// }
// func UpdateMenuPg(c echo.Context) error {
// 	id := c.Param("id")
// 	var menu request.CreateRequest
// 	if err := c.Bind(&menu); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	if err := manager.UpdateMenuPg(id, menu); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, "sucessfully updated")
// }
// func DeleteMenuPg(c echo.Context) error {
// 	id := c.Param("id")
// 	if err := manager.DeleteMenuPg(id); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("error in delete:%v", err)})
// 	}
// 	return c.JSON(http.StatusOK, map[string]string{"success": fmt.Sprintf("menuitem with id %v deleted successfull.", id)})
// }
