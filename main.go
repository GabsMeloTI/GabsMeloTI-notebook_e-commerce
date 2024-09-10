package main

import (
	"github.com/joho/godotenv"
	"log"
	"sales/config"
	"sales/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	err = config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db.InitGorm()

}
