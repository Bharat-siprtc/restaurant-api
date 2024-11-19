package services

import (
	"context"
	"example/restaurant-api/config"
	"example/restaurant-api/models"
	"example/restaurant-api/request"
	"example/restaurant-api/response"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMenu() ([]response.AllMenuResponse, error) {
	var menu []response.AllMenuResponse
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &menu); err != nil {
		return nil, err
	}
	return menu, nil
}
func CreateMenu(menu request.CreateRequest) error {
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// var resp response.CreateResponse
	_, err := collection.InsertOne(ctx, menu)
	if err != nil {
		return err
	}
	return nil
}
func DeleteMenu(id string) error {
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	intId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("please enter valid id")
	}
	res, err := collection.DeleteOne(ctx, bson.M{"id": intId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no id match with %v in database", id)
	}
	return nil
}

func UpdateMenu(id string, req request.UpdateRequest) error {
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{
			"name":     req.Name,
			"category": req.Category,
			"desc":     req.Desc,
			"price":    req.Price,
		},
	}
	intId, err := strconv.Atoi(id)

	res, err := collection.UpdateOne(ctx, bson.M{"id": intId}, update)
	if err != nil {
		return fmt.Errorf("there is an error in updation:%v", err)
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("no id matched with the given id %v", id)
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed! there is an problem in updation of id %v", id)
	}
	return nil
}
func GetMenuPg() ([]response.AllMenuResponse, error) {
	var menu []response.AllMenuResponse

	// Fetch all movies from the database
	rows, err := config.PG.Query("SELECT id, name, category, description, price FROM menu")
	if err != nil {
		log.Println("Error getting movies:", err)
		return nil, fmt.Errorf("Error getting movies: %v", err)
	}
	defer rows.Close()

	// Iterate through rows and append to the movies slice
	for rows.Next() {
		var menuItem response.AllMenuResponse
		// var id primitive.ObjectID
		if err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Category, &menuItem.Desc, &menuItem.Price); err != nil {
			log.Println("Error scanning menu item:", err)
			return nil, fmt.Errorf("Error scanning menu item: %v", err)
		}
		menu = append(menu, menuItem)
	}
	return menu, nil
}
func GetMenuPage(page string, size string) ([]models.MenuItem, error) {
	var menu []models.MenuItem
	pg, err := strconv.Atoi(page)
	if err != nil {
		return nil, fmt.Errorf("please enter valid page no")
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		return nil, fmt.Errorf("please enter valid size")
	}
	offset := (pg - 1) * sz
	// Fetch menu from the database
	// var totalCount int
	rows, err := config.PG.Query("SELECT id, name, category, description, price FROM menu LIMIT $1 OFFSET $2", sz, offset)
	if err != nil {
		log.Println("Error getting movies:", err)
		return nil, fmt.Errorf("Error getting movies: %v", err)
	}
	// _,er:=config.PG.Query("SELECT COUNT(*) FROM menu").Scan(&totalCount)

	// if err != nil {
	// 	log.Println("Error getting movies:", err)
	// 	return nil, fmt.Errorf("Error getting movies: %v", err)
	// }
	defer rows.Close()
	// Iterate through rows and append to the movies slice
	for rows.Next() {
		var menuItem models.MenuItem
		if err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Category, &menuItem.Desc, &menuItem.Price); err != nil {
			log.Println("Error scanning menu item:", err)
			return nil, fmt.Errorf("Error scanning menu item: %v", err)
		}
		menu = append(menu, menuItem)
	}
	return menu, nil
}
func GetMenuByIdPg(id string) (models.MenuItem, error) {
	// query:="SELECT * FROM menu WHERE id=$1"
	var menuItem models.MenuItem
	err := config.PG.QueryRow("SELECT id, name, category, description, price FROM menu WHERE id = $1", id).Scan(
		&menuItem.ID,
		&menuItem.Name,
		&menuItem.Category,
		&menuItem.Desc,
		&menuItem.Price,
	)
	if err != nil {
		return models.MenuItem{}, fmt.Errorf("error in fetching data:%v", err)
	}
	return menuItem, nil
}
func CreateMenuPg(menu request.CreateRequest) error {
	query := "INSERT INTO menu (name, category, description, price) VALUES ($1, $2, $3, $4)"
	_, err := config.PG.Exec(query, menu.Name, menu.Category, menu.Desc, menu.Price)
	if err != nil {
		log.Println("Error inserting menu item:", err)
		return fmt.Errorf("error inserting menu item: %v", err)
	}
	log.Println("Menu item created successfully")
	return nil
}
func UpdateMenuPg(id string, menu request.UpdateRequest) error {
	query := "UPDATE menu SET name=$1, category=$2, description=$3, price=$4 WHERE id=$5"
	result, err := config.PG.Exec(query, menu.Name, menu.Category, menu.Desc, menu.Price, id)
	if err != nil {
		log.Println("Error inserting menu item:", err)
		return fmt.Errorf("error updating menu item: %v", err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no menu item found with ID %v", id)
	}
	log.Println("Menu item updated successfully")
	return nil
}
func DeleteMenuPg(id string) error {
	query := "DELETE FROM menu WHERE id=$1"
	result, err := config.PG.Exec(query, id)
	if err != nil {
		log.Println("Error Deleting menu item:", err)
		return fmt.Errorf("Error updating menu item: %v", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("No menu item found with ID %v", id)
	}

	log.Println("Menu item Deleted successfully")
	return nil
}
