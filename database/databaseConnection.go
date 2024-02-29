package database

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error reading environment file")
	}
}
