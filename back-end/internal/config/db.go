package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDb() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=pc_components port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("db Connected Successfully")

	return db
}
