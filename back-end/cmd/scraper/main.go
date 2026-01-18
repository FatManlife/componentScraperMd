package main

import (
	"sync"

	"github.com/FatManlife/component-finder/back-end/internal/scraper/neocomputer"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/pcprime"
	"github.com/FatManlife/component-finder/back-end/internal/scraper/xstore"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func (){
		defer wg.Done()
		xstore.Run()
	}()

	go func (){
		defer wg.Done()
		pcprime.Run()
	}()

	go func (){
		defer wg.Done()
		neocomputer.Run()
	}()

	wg.Wait()	
}
