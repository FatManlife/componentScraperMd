package main

import (
	handler "github.com/FatManlife/component-finder/back-end/internal/api/handlers"
	"github.com/FatManlife/component-finder/back-end/internal/config"
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	repo "github.com/FatManlife/component-finder/back-end/internal/repositories"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	projUtils "github.com/FatManlife/component-finder/back-end/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnDb()

	//Product Initalization
	productRepo := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHanlder := handler.NewProductHandler(productService)
		
	//Sdd Initalization
	ssdRepo := repo.NewSSDRepository(db)
	ssdService := service.NewComponentService[dto.SsdResponse, dto.SsdParams](ssdRepo.GetSsds, productRepo.GetProductByID,projUtils.SsdMapping) 
	ssdHandler := handler.NewComponentHandler(ssdService)

	//Ram Initalization
	ramRepo := repo.NewRamRepository(db)
	ramService := service.NewComponentService[dto.RamResponse, dto.RamParams](ramRepo.GetRams, productRepo.GetProductByID,projUtils.RamMapping) 
	ramHandler := handler.NewComponentHandler(ramService)

	//PSU Initalization
	psuRepo := repo.NewPsuRepository(db)
	psuService := service.NewComponentService[dto.PsuResponse, dto.PsuParams](psuRepo.GetPsus, productRepo.GetProductByID,projUtils.PsuMapping) 
	psuHandler := handler.NewComponentHandler(psuService)

	//Pc Mini Initalization
	pcMiniRepo := repo.NewPcMiniRepository(db)
	pcMiniService := service.NewComponentService[dto.PcMiniResponse, dto.PcParams](pcMiniRepo.GetPcMinis, productRepo.GetProductByID,projUtils.PcMiniMapping) 
	pcMiniHandler := handler.NewComponentHandler(pcMiniService)
	
	//Pc Initalization
	pcRepo := repo.NewPcRepository(db)
	pcService := service.NewComponentService[dto.PcResponse, dto.PcParams](pcRepo.GetPcs, productRepo.GetProductByID,projUtils.PcMapping) 
	pcHandler := handler.NewComponentHandler(pcService)

	//Motherboard Initalization
	mbRepo := repo.NewMotherboardRepository(db)
	mbService := service.NewComponentService[dto.MotherboardResponse, dto.MotherboardParams](mbRepo.GetMotherboards, productRepo.GetProductByID,projUtils.MbMapping) 
	mbHandler := handler.NewComponentHandler(mbService)

	//Laptop Initalization
	laptopRepo := repo.NewLaptopRepository(db)
	laptopService := service.NewComponentService[dto.LaptopResponse, dto.LaptopParams](laptopRepo.GetLaptops, productRepo.GetProductByID,projUtils.LaptopMapping) 
	laptopHandler := handler.NewComponentHandler(laptopService)

	//Gpus Initalization
	gpuRepo := repo.NewGpuRepository(db)
	gpuService := service.NewComponentService[dto.GpuResponse, dto.GpuParams](gpuRepo.GetGpus, productRepo.GetProductByID,projUtils.GpuMapping)
	gpuHandler := handler.NewComponentHandler(gpuService)

	//Fans Initalization
	fanRepo := repo.NewFanRepository(db)
	fanService := service.NewComponentService[dto.FanResponse, dto.FanParams](fanRepo.GetFans, productRepo.GetProductByID,projUtils.FanMapping)
	fanHandler := handler.NewComponentHandler(fanService)

	//Aio Initalization
	aioRepo := repo.NewAioRepository(db)
	aioService := service.NewComponentService[dto.AioResponse, dto.AioParams](aioRepo.GetAios, productRepo.GetProductByID,projUtils.AioMapping) 
	aioHandler := handler.NewComponentHandler(aioService)

	//Cpu Initalization
	cpuRepo := repo.NewCpuRepository(db)
	cpuService := service.NewComponentService[dto.CpuResponse, dto.CpuParams](cpuRepo.GetCpus, productRepo.GetProductByID,projUtils.CpuMapping) 
	cpuHandler := handler.NewComponentHandler(cpuService)
	
	//Case Initialization
	caseRepo := repo.NewCaseRepository(db)
	caseService := service.NewComponentService[dto.CaseResponse, dto.CaseParams](caseRepo.GetCases, productRepo.GetProductByID,projUtils.CaseMapping) 
	caseHandler := handler.NewComponentHandler(caseService)
	
	//Case HDD Initalization
	hddRepo := repo.NewHddRepository(db)
	hddService := service.NewComponentService[dto.HddResponse, dto.HddParams](hddRepo.GetHdds, productRepo.GetProductByID,projUtils.HddMapping) 
	hddHandler := handler.NewComponentHandler(hddService)

	//Cooler Initalization
	coolerRepo := repo.NewCoolerRepository(db)
	coolerService := service.NewComponentService[dto.CoolerResponse, dto.CoolerParams](coolerRepo.GetCoolers, productRepo.GetProductByID,projUtils.CoolerMapping) 
	coolerHandler := handler.NewComponentHandler(coolerService)

	//Filter Initialization
	websiteRepo := repo.NewWebsiteRepository(db)
	filterService := service.NewFilterService(websiteRepo, productRepo)
	filterHandler := handler.NewFilterHandler(filterService)

	//Initiazlie Gin Router
	r := gin.Default()
	r.Use(cors.Default())

	//Http Requests products getters
	r.GET("/", productHanlder.GetProducts)
	r.GET("/count", productHanlder.GetProductsCount)

	//Http Requests component getters
	//Filters
	filters := r.Group("/filter")
	{
		filters.GET("", filterHandler.GetDefaultFilters)
	}
	//Cooler
	cooler := r.Group("/cooler")
	{
		cooler.GET("", coolerHandler.GetComponents)
		cooler.GET("/:id", coolerHandler.GetComponentByID)
	}
	//Hdd
	hdd := r.Group("/hdd")
	{
		hdd.GET("", hddHandler.GetComponents)
		hdd.GET("/:id", hddHandler.GetComponentByID)
	}
	//SSD
	ssd := r.Group("/ssd")
	{
		ssd.GET("", ssdHandler.GetComponents)
		ssd.GET("/:id", ssdHandler.GetComponentByID)
	}
	//RAM
	ram := r.Group("/ram")
	{
	ram.GET("", ramHandler.GetComponents)
	ram.GET("/:id", ramHandler.GetComponentByID)
	}
	//PSU
	psu := r.Group("/psu")
	{
		psu.GET("", psuHandler.GetComponents)
		psu.GET("/:id", psuHandler.GetComponentByID)
	}
	//PC Mini
	pcmini := r.Group("/pcmini")
	{
		pcmini.GET("", pcMiniHandler.GetComponents)
		pcmini.GET("/:id", pcMiniHandler.GetComponentByID)
	}
	//PC
	pc := r.Group("/pc")
	{
		pc.GET("", pcHandler.GetComponents)
		pc.GET("/:id", pcHandler.GetComponentByID)
	}
	//Motherboard
	motherboard := r.Group("/motherboard")
	{
		motherboard.GET("", mbHandler.GetComponents)
		motherboard.GET("/:id", mbHandler.GetComponentByID)
	}
	//Laptop
	laptop := r.Group("/laptop")
	{
		laptop.GET("", laptopHandler.GetComponents)
		laptop.GET("/:id", laptopHandler.GetComponentByID)
	}
	//GPU
	gpu := r.Group("/gpu")
	{
		gpu.GET("", gpuHandler.GetComponents)
		gpu.GET("/:id", gpuHandler.GetComponentByID)
	}
	//Fan
	fan := r.Group("/fan")
	{
		fan.GET("", fanHandler.GetComponents)
		fan.GET("/:id", fanHandler.GetComponentByID)
	}
	//AIO
	aio := r.Group("/aio")
	{
		aio.GET("", aioHandler.GetComponents)
		aio.GET("/:id", aioHandler.GetComponentByID)
	}
	//CPU
	cpu := r.Group("/cpu")
	{
		cpu.GET("", cpuHandler.GetComponents)
		cpu.GET("/:id", cpuHandler.GetComponentByID)
	}
	//Case
	pcCase := r.Group("/case")
	{
		pcCase.GET("", caseHandler.GetComponents)
		pcCase.GET("/:id", caseHandler.GetComponentByID)
	}
	
	// listen and serve on
	r.Run(":8080") 
}
