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
	} else if flag == "false" || flag == "" {
		return services.GetMenuPg()
	} else {
		return nil, fmt.Errorf("flag should be true or false only")
	}

	// return services.GetMenu()
}
func GetMenuById(id string, flag string) (response.AllMenuResponse, error) {
	if flag == "true" {
		return services.GetMenuById(id)
	} else if flag == "false" || flag == "" {
		return services.GetMenuByIdPg(id)
	} else {
		return response.AllMenuResponse{}, fmt.Errorf("flag should be true or false only")
	}
}
func CreateMenu(flag string, menu request.CreateRequest) (response.AllMenuResponse, error) {
	if flag == "true" {
		return services.CreateMenu(menu)
	} else if flag == "false" || flag == "" {
		return services.CreateMenuPg(menu)
	} else {
		return response.AllMenuResponse{}, fmt.Errorf("flag should be true or false only")
	}
}
func DeleteMenu(flag string, id string) error {
	if flag == "true" {
		fmt.Printf("mongo db is called.")
		return services.DeleteMenu(id)
	} else if flag == "false" || flag == "" {
		return services.DeleteMenuPg(id)
	} else {
		return fmt.Errorf("flag should be true or false only")
	}
}
func UpdateMenu(flag string, id string, req request.CreateRequest) (response.AllMenuResponse, error) {
	if flag == "true" {
		fmt.Printf("mongo db is called.")
		return services.UpdateMenu(id, req)
	} else if flag == "false" || flag == "" {
		return services.UpdateMenuPg(id, req)
	} else {
		return response.AllMenuResponse{}, fmt.Errorf("flag should be true or false only")
	}
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
