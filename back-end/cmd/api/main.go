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

	//PcMinis Initalization
	pcMiniRepo := repo.NewPcMiniRepository(db)
	pcMiniService := service.NewPcMiniService(pcMiniRepo)
	pcMiniHandler := handler.NewPcMiniHandler(pcMiniService)

	//Pc Initalization
	pcRepo := repo.NewPcRepository(db)
	pcService := service.NewPcService(pcRepo)
	pcHandler := handler.NewPcHandler(pcService)

	//Motherboard Initalization
	mbRepo := repo.NewMotherboardRepository(db)
	mbService := service.NewMotherboardService(mbRepo)
	mbHandler := handler.NewMotherboardHandler(mbService)

	//Laptop Initalization
	laptopRepo := repo.NewLaptopRepository(db)
	laptopService := service.NewLaptopService(laptopRepo)
	laptopHandler := handler.NewLaptopHandler(laptopService)

	//Gpus Initalization
	gpuRepo := repo.NewGpuRepository(db)
	gpuService := service.NewGpuService(gpuRepo)
	gpuHandler := handler.NewGpuHandler(gpuService)

	//Fans Initalization
	fanRepo := repo.NewFanRepository(db)
	fanService := service.NewFanService(fanRepo)
	fanHandler := handler.NewFanHandler(fanService)

	//Cpu Initalization
	cpuRepo := repo.NewCpuRepository(db)
	cpuService := service.NewCpuService(cpuRepo)
	cpuHandler := handler.NewCpuHandler(cpuService)

	//Case Initialization
	caseRepo := repo.NewCaseRepository(db)
	caseService := service.NewCaseService(caseRepo)
	caseHandler := handler.NewCaseHandler(caseService)

	//Initiazlie Gin Router
	r := gin.Default()

	//Http Requests products getters
	r.GET("/products", productHanlder.GetProducts)
	r.GET("/aio", aioHandler.GetAios)
	r.GET("/ssd", ssdHandler.GetSsds)
	r.GET("/ram", ramHandler.GetRams)
	r.GET("/psu", psuHandler.GetPsus)
	r.GET("/pcmini", pcMiniHandler.GetPcMinis)
	r.GET("/pc", pcHandler.GetPcs)
	r.GET("/motherboard", mbHandler.GetMotherboards)
	r.GET("/laptop", laptopHandler.GetLaptops)
	r.GET("/gpu", gpuHandler.GetGpus)
	r.GET("/fan", fanHandler.GetFans)
	r.GET("/cpu", cpuHandler.GetCpus)
	r.GET("/case", caseHandler.GetCases)
	
	// listen and serve on
	r.Run(":8080") 
}
