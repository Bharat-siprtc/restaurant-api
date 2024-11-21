package response

type AllMenuResponse struct {
	ID       int     `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Desc     string  `json:"description" bson:"description"`
	Price    float64 `json:"price" bson:"price"`
}
type IdResponse struct{
	ID int `bson:"seq"`
}
// type CreateResponse struct {
// 	ID       primitive.ObjectID `json:"id" bson:"_id"`
// 	Name     string             `json:"name" bson:"name"`
// 	Category string             `json:"category" bson:"category"`
// 	Desc     string             `json:"desc" bson:"desc"`
// 	Price    float64            `json:"price" bson:"price"`
// }
