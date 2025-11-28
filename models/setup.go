package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ESClient *elasticsearch.Client

const SearchIndex = "blogs"

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := func() int {
		port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		return port
	}()
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(Blog{})
}

func ESClientConnection() {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	ESClient = client

}

func ESCreateIndexIfNotExists() {
	_, err := esapi.IndicesExistsRequest{
		Index: []string{SearchIndex},
	}.Do(context.Background(), ESClient)

	if err != nil {
		ESClient.Indices.Create(SearchIndex)
	}

}
