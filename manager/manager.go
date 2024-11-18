package manager

import (
	"fmt"
	"restaurant-api/models"
	"restaurant-api/request"
	"restaurant-api/response"
	"restaurant-api/services"

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
func CreateMovie(menu request.CreateRequest) error {
	if err := ValidateStruct(menu); err != nil {
		return err
	}
	return services.CreateMovie(menu)
}
func DeleteMovie(id string) error {
	return services.DeleteMovie(id)
}
func UpdateMovie(id string, req request.UpdateRequest) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return services.UpdateMovie(objId, req)
}
func GetMenuPg() ([]models.MenuItem, error) {
	return services.GetMenuPg()
}
