package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	// Read environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "prefer"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, database, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func InitGorm() {
	var err error
	DB, err = OpenConnection()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if DB == nil {
		log.Fatalf("Database instance is nil after initialization")
	}

	log.Println("Database connection established successfully")
}
