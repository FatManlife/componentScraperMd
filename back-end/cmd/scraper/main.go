package main

import (
	"github.com/FatManlife/component-finder/back-end/internal/scraper/xstore"
	_ "github.com/FatManlife/component-finder/back-end/internal/scraper/xstore"
)

func main() {
	xstore.Run()
	//test.TestColly()
}
