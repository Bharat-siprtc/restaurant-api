package models

type MenuItem struct {
	ID       int     `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Desc     string  `json:"description" bson:"description"`
	Price    float32 `json:"price" bson:"price"`
}
type Config struct {
    MongoDB MongoDBConfig `toml:"mongodb"`
}

type MongoDBConfig struct {
    URI      string `toml:"uri"`      // MongoDB connection string
    Database string `toml:"database"` // Database name
}
