package services

import (
	"context"
	"fmt"
	"log"
	"restaurant-api/config"
	"restaurant-api/models"
	"restaurant-api/request"
	"restaurant-api/response"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func CreateMovie(menu request.CreateRequest) error {
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
func DeleteMovie(id string) error {
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no id match with %v in database", id)
	}
	return nil
}

func UpdateMovie(objId primitive.ObjectID, req request.UpdateRequest) error {
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
	res, err := collection.UpdateByID(ctx, objId, update)
	if err != nil {
		fmt.Errorf("there is an error in updation:%v", err)
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("no id matched with the given id %v", objId)
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed! there is an problem in updation of id %v", objId)
	}
	return nil
}
func GetMenuPg() ([]models.MenuItem, error) {
	var menu []models.MenuItem

	// Fetch all movies from the database
	rows, err := config.PG.Query("SELECT * FROM menu")
	if err != nil {
		log.Println("Error getting movies:", err)
		return nil, fmt.Errorf("Error getting movies: %v", err)
	}
	defer rows.Close()

	// Iterate through rows and append to the movies slice
	for rows.Next() {
		var menuItem models.MenuItem
		if err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Category, &menuItem.Desc, &menuItem.Price); err != nil {
			log.Println("Error scanning movie:", err)
			return nil, fmt.Errorf("Error scanning movie: %v", err)
		}
		menu = append(menu, menuItem)
	}
	
	return menu, nil
}
