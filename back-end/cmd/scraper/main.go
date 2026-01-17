package main

import (
	"github.com/FatManlife/component-finder/back-end/internal/test"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func (){
	// 	defer wg.Done()
	// 	xstore.Run()
	// }()

	// go func (){
	// 	defer wg.Done()
	// 	pcprime.Run()
	// }()

	// wg.Wait()
	
	test.TestColly()
}
