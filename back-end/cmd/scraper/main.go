package main

import (
	"fmt"
	"sync"

	"github.com/FatManlife/component-finder/back-end/internal/config"
	ormsql "github.com/FatManlife/component-finder/back-end/internal/db/orm_sql"
	rawsql "github.com/FatManlife/component-finder/back-end/internal/db/raw_sql"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/neocomputer"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/pcprime"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/xstore"
	"github.com/FatManlife/component-finder/back-end/internal/utils"
)

func main() {
	db := config.ConnDb()
	
	links, err := ormsql.GetAllProductLinks(db)
	
	if err != nil {
		fmt.Println("Error fetching product links:", err)
	} 
	
	utils.PopulateVisitedUrls(links)

	storage := rawsql.NewStorage(db)

	var wg sync.WaitGroup
	wg.Add(3)

	go func (){
		defer wg.Done()
		xstore.Run(storage)
	}()

	go func (){
		defer wg.Done()
		pcprime.Run(storage)
	}()

	go func (){
		defer wg.Done()
		neocomputer.Run(storage)
	}()

	wg.Wait()	
}
