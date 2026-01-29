package main

import (
	handler "github.com/FatManlife/component-finder/back-end/internal/api/handlers"
	"github.com/FatManlife/component-finder/back-end/internal/config"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnDb()

	//Product Initalization
	productRepo := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHanlder := handler.NewProductHandler(productService)

	//Aio Initalization
	aioRepo := repo.NewAioRepository(db)
	aioService := service.NewAioService(aioRepo)
	aioHandler := handler.NewAioHandler(aioService)

	//Ssd Initalization
	ssdRepo := repo.NewSSDRepository(db)
	ssdService := service.NewSSDService(ssdRepo)
	ssdHandler := handler.NewSSDHandler(ssdService)

	//Ram Initalization
	ramRRepo := repo.NewRamRepository(db)
	ramService := service.NewRamService(ramRRepo)
	ramHandler := handler.NewRamHandler(ramService)

	//PSU Initalization
	psuRepo := repo.NewPsuRepository(db)
	psuService := service.NewPsuService(psuRepo)
	psuHandler := handler.NewPsuHandler(psuService)

	//Initiazlie Gin Router
	r := gin.Default()

	//Http Requests products
	r.GET("/products", productHanlder.GetProducts)
	r.GET("/aio", aioHandler.GetAios)
	r.GET("/ssd", ssdHandler.GetSsdsHandler) 
	r.GET("/ram", ramHandler.GetRams)
	r.GET("/psu", psuHandler.GetPsusHandler)

	// listen and serve on
	r.Run(":8080") 
}
