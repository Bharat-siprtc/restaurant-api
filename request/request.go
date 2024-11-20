package request

type CreateRequest struct {
	ID       int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Desc     string  `json:"description" bson:"desc"`
	Price    float64 `json:"price" bson:"price"`
}
type UpdateRequest struct {
	Name     string  `json:"name" bson:"name" validate="required"`
	Category string  `json:"category" bson:"category" validate="required"`
	Desc     string  `json:"description" bson:"desc" validate="required"`
	Price    float64 `json:"price" bson:"price" validate="required"`
}
