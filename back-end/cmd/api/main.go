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

	//Initiazlie Gin Router
	r := gin.Default()

	//Http Requests products
	r.GET("/products", productHanlder.GetAllProducts)

	// listen and serve on
	r.Run(":8080") 
}
