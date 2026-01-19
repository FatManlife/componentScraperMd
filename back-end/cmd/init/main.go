package main

import (
	"log"

	"github.com/FatManlife/component-finder/back-end/internal/config"
	"github.com/FatManlife/component-finder/back-end/internal/migrations"
)

func main() {
	Db := config.ConnDb()
	
	if err := migrations.Migrate(Db); err != nil {
		log.Fatal("Migration failed:", err)
	}
}
