package config

import (
	"context"
	"database/sql"
	"example/restaurant-api/models"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var PG *sql.DB

func ConnectDB() {
	var config models.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(config.MongoDB.Database)
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
	// log.Printf("Connecting to PostgreSQL with DSN: %s", "postgres", "user=postgres dbname=restaurant sslmode=disable")
	// Connect to PostgreSQL
	// PGDB, err := sql.Open("postgres", "user=postgres dbname=restaurant sslmode=disable password=2503")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer PGDB.Close()

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
	PG = PGDB
	log.Println("Connected to PostgreSQL successfully!")
}
