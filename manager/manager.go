package manager

import (
	"example/restaurant-api/models"
	"example/restaurant-api/request"
	"example/restaurant-api/response"
	"example/restaurant-api/services"
	"fmt"
)

func GetMenu(flag string) ([]response.AllMenuResponse, error) {
	if flag == "true" {
		fmt.Printf("mongo db is called.")
		return services.GetMenu()
	}
	// return services.GetMenu()
	return services.GetMenuPg()
}
func GetMenuById(id string, flag string) (response.AllMenuResponse, error) {
	if flag == "true" {
		return services.GetMenuById(id)
	}
	return services.GetMenuByIdPg(id)
}
func CreateMenu(flag string, menu request.CreateRequest) (response.AllMenuResponse,error) {
	if flag == "true" {
		return services.CreateMenu(menu)
	}
	return services.CreateMenuPg(menu)
}
func DeleteMenu(flag string, id string) error {
	if flag == "true" {
		fmt.Printf("mongo db is called.")
		return services.DeleteMenu(id)
	}
	return services.DeleteMenuPg(id)
}
func UpdateMenu(flag string, id string, req request.CreateRequest) (response.AllMenuResponse,error) {
	if flag == "true" {
		fmt.Printf("mongo db is called.")
		return services.UpdateMenu(id, req)
	}
	return services.UpdateMenuPg(id, req)
}

//	func GetMenuPg() ([]models.MenuItem, error) {
//		return services.GetMenuPg()
//	}
func GetMenuPage(page string, size string) ([]models.MenuItem, error) {
	return services.GetMenuPage(page, size)
}

//	func GetMenuByIdPg(id string) (models.MenuItem, error) {
//		return services.GetMenuByIdPg(id)
//	}
// func CreateMenuPg(menu request.CreateRequest) error {
// 	return services.CreateMenuPg(menu)
// }
// func UpdateMenuPg(id string, menu request.CreateRequest) error {
// 	return services.UpdateMenuPg(id, menu)
// }
// func DeleteMenuPg(id string) error {
// 	return services.DeleteMenuPg(id)
// }
