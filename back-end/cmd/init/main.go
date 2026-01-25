package main

import (
	"fmt"
	"log"

	"github.com/FatManlife/component-finder/back-end/internal/config"
	"github.com/FatManlife/component-finder/back-end/internal/migrations"
	"github.com/FatManlife/component-finder/back-end/internal/seed"
)

func main() {
	db := config.ConnDb()

	//Delete tables
	// if err := migrations.DelteTables(db); err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }
	
	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	if err := seed.SeedWeb(db); err != nil {
		fmt.Println("Initial website creation failed:", err)
	}

	if err := seed.SeedCategory(db); err != nil {
		fmt.Println("Initial category creation failed:", err)
	}
}
