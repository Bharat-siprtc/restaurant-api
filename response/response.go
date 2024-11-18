package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllMenuResponse struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Category string             `json:"category" bson:"category"`
	Desc     string             `json:"desc" bson:"desc"`
	Price    float64            `json:"price" bson:"price"`
}

type CreateResponse struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Category string             `json:"category" bson:"category"`
	Desc     string             `json:"desc" bson:"desc"`
	Price    float64            `json:"price" bson:"price"`
}
