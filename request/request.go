package request

type CreateRequest struct {
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Desc     string  `json:"description" bson:"description"`
	Price    float64 `json:"price" bson:"price"`
}

