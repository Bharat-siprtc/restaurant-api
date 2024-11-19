package manager

import (
	"example/restaurant-api/models"
	"example/restaurant-api/request"
	"example/restaurant-api/response"
	"example/restaurant-api/services"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateStruct(input interface{}) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return fmt.Errorf("Validation failed!:%v", err)
	}
	return nil
}
func GetMenu() ([]response.AllMenuResponse, error) {
	return services.GetMenu()
}
func CreateMenu(menu request.CreateRequest) error {
	if err := ValidateStruct(menu); err != nil {
		return err
	}
	return services.CreateMenu(menu)
}
func DeleteMenu(id string) error {
	return services.DeleteMenu(id)
}
func UpdateMenu(id string, req request.UpdateRequest) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return services.UpdateMenu(objId, req)
}
func GetMenuPg() ([]models.MenuItem, error) {
	return services.GetMenuPg()
}
func GetMenuPage(page string,size string) ([]models.MenuItem, error) {
	return services.GetMenuPage(page,size)
}
func GetMenuByIdPg(id string) (models.MenuItem,error) {
	return services.GetMenuByIdPg(id)
}
func CreateMenuPg(menu request.CreateRequest)error{
	return services.CreateMenuPg(menu)
}
func UpdateMenuPg(id string,menu request.UpdateRequest)  error{
	return services.UpdateMenuPg(id,menu)
}
func DeleteMenuPg(id string) error {
	return services.DeleteMenuPg(id)
}
