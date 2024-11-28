package config

type MongoConfig struct {
	MongoDBURL  string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017"`
	MongoDBName string `env:"MONGODB_DB_NAME" envDefault:"restaurantDB"`
}
type Config struct {
    MongoDB MongoConfig `toml:"mongodb"`
}
