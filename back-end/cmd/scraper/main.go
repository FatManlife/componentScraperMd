package main

import (
	_ "github.com/FatManlife/component-finder/back-end/internal/scraper/xstore"
	"github.com/FatManlife/component-finder/back-end/internal/test"
)

func main() {
	//xstore.Run()
	test.TestColly()
}
