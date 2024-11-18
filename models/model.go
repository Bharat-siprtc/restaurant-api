package models

type MenuItem struct {
	ID       int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Desc     string  `json:"description" bson:"description"`
	Price    float32 `json:"price" bson:"price"`
}
