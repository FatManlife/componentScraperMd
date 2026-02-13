package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDb() *gorm.DB {
	err := godotenv.Load()
	
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("db Connected Successfully")

	return db
}
