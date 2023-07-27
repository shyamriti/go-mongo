package main

import (
	"context"
	"go-mongo/pkg/database"
	"go-mongo/pkg/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load("/home/shyam/go-workspace/go-mongo/.env")
	if err != nil {
		log.Print(err.Error())
	}

	// Access the environment variables
	dbColl := os.Getenv("DB_COLLECTION")

	database.Connection(dbColl)
	defer database.Client.Disconnect(context.TODO())

	r := router.Router()

	// Start the server
	r.Run(":8080")
}
