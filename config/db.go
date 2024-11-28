package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var PG *sql.DB

// func loadConfig(filePath string, config *MongoConfig) error {
// 	// Viper is one option to load TOML config
// 	viper.SetConfigType("toml")
// 	viper.SetConfigFile(filePath)

// 	if err := viper.ReadInConfig(); err != nil {
// 		return fmt.Errorf("error reading config file, %s", err)
// 	}

// 	// Bind the values to the struct
// 	if err := viper.Unmarshal(config); err != nil {
// 		return fmt.Errorf("unable to decode into struct, %v", err)
// 	}

// 	return nil
// }
// func ConnectDB() {
// 	var config MongoConfig
// 	if err := loadConfig("config.toml", &config); err != nil {
// 		log.Fatalf("Error loading config file: %v", err)
// 	}
// 	fmt.Println("MongoDB URL from config:", config.MongoDBURL) // Debugging line
// 	clientOptions := options.Client().ApplyURI(config.MongoDBURL)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		DB = client.Database(config.MongoDBName)
//		fmt.Println(DB)
//	}
//
// var mongoCon MongoConfig
// var a=mongoCon.MongoDBURL
func ConnectDB() {
	cfg := MongoConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse env vars: %v", err)
	}
	fmt.Printf("Connecting to MongoDB at %s using database %s\n", cfg.MongoDBURL, cfg.MongoDBName)
	// var mongoConfig models.Config
	// if _, err := toml.DecodeFile("config.toml", &config); err != nil {
	// 	log.Fatalf("Error loading config file: %v", err)
	// }
	clientOptions := options.Client().ApplyURI(cfg.MongoDBURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(cfg.MongoDBName)
}

func PostgresConnect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// PostgreSQL connection details from .env
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")

	// Format PostgreSQL DSN string
	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB)

	PGDB, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}

	// Test PostgreSQL connection
	err = PGDB.Ping()
	if err != nil {
		log.Fatal("Error pinging PostgreSQL:", err)
	}
	if PGDB == nil {
		log.Println("Database connection is nil!")
	}
	err = createMenuTable(PGDB)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	PG = PGDB
	log.Println("Connected to PostgreSQL successfully!")
}
func createMenuTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS menu (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		category VARCHAR(50),
		description TEXT,
		price NUMERIC(10, 2)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	return nil
}

func CreateCounterSeq() error {
	collection := DB.Collection("counters") // Counter collection name

	// Check if the sequence document exists
	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{"_id": "restaurant_id"}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Sequence document doesn't exist, create it with an initial value of 0
		_, err = collection.InsertOne(context.Background(), bson.M{
			"_id": "restaurant_id", // Unique identifier for the sequence
			"seq": 0,               // Initial sequence value
		})
		if err != nil {
			return fmt.Errorf("error creating counter sequence document: %v", err)
		}
		fmt.Println("Counter sequence document created with initial value.")
	} else if err != nil {
		return fmt.Errorf("error checking counter sequence document: %v", err)
	} else {
		fmt.Println("Counter sequence document already exists.")
	}
	return nil
}
