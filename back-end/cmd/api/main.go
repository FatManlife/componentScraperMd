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

	//Website Initalization
	websiteRepo := repo.NewWebsiteRepository(db)

	//Product Initalization
	productRepo := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepo, websiteRepo)
	productHanlder := handler.NewProductHandler(productService)
		
	//Sdd Initalization
	ssdRepo := repo.NewSSDRepository(db)
	ssdService := service.NewSsdService(ssdRepo)
	ssdGenericService := service.NewComponentService[dto.SsdResponse, dto.SsdParams](ssdRepo.GetSsds, productRepo.GetProductByID,projUtils.SsdMapping) 
	ssdGenericHandler := handler.NewComponentHandler(ssdGenericService)
	ssdSpecHandler := handler.NewSpecHandler[dto.SsdSpecs](productService.GetDefaultSpecs,ssdService.GetSpecs,"ssd")

	//Ram Initalization
	ramRepo := repo.NewRamRepository(db)
	ramService := service.NewRamService(ramRepo)
	ramGenericService := service.NewComponentService[dto.RamResponse, dto.RamParams](ramRepo.GetRams, productRepo.GetProductByID,projUtils.RamMapping) 
	ramGenericHandler := handler.NewComponentHandler(ramGenericService)
	ramSpecHandler := handler.NewSpecHandler[dto.RamSpecs](productService.GetDefaultSpecs,ramService.GetSpecs,"ram")

	//PSU Initalization
	psuRepo := repo.NewPsuRepository(db)
	psuService := service.NewPsuService(psuRepo)
	psuGenericService := service.NewComponentService[dto.PsuResponse, dto.PsuParams](psuRepo.GetPsus, productRepo.GetProductByID,projUtils.PsuMapping) 
	psuGenericHandler := handler.NewComponentHandler(psuGenericService)
	psuSpecHandler := handler.NewSpecHandler[dto.PsuSpecs](productService.GetDefaultSpecs,psuService.GetSpecs,"psu")

	//Pc Mini Initalization
	pcMiniRepo := repo.NewPcMiniRepository(db)
	pcMiniService := service.NewPcMiniService(pcMiniRepo)
	pcMiniGenericService := service.NewComponentService[dto.PcMiniResponse, dto.PcParams](pcMiniRepo.GetPcMinis, productRepo.GetProductByID,projUtils.PcMiniMapping) 
	pcMiniGenericHandler := handler.NewComponentHandler(pcMiniGenericService)
	pcMiniSpecHandler := handler.NewSpecHandler[dto.PcSpecs](productService.GetDefaultSpecs,pcMiniService.GetSpecs,"pc_mini")
	
	//Pc Initalization
	pcRepo := repo.NewPcRepository(db)
	pcService := service.NewPcService(pcRepo)
	pcGenericService := service.NewComponentService[dto.PcResponse, dto.PcParams](pcRepo.GetPcs, productRepo.GetProductByID,projUtils.PcMapping) 
	pcGenericHandler := handler.NewComponentHandler(pcGenericService)
	pcSpecHandler := handler.NewSpecHandler[dto.PcSpecs](productService.GetDefaultSpecs,pcService.GetSpecs,"pc")

	//Motherboard Initalization
	mbRepo := repo.NewMotherboardRepository(db)
	mbService := service.NewMotherboardService(mbRepo)
	mbGenericService := service.NewComponentService[dto.MotherboardResponse, dto.MotherboardParams](mbRepo.GetMotherboards, productRepo.GetProductByID,projUtils.MbMapping) 
	mbGenericHandler := handler.NewComponentHandler(mbGenericService)
	mbSpecHandler := handler.NewSpecHandler[dto.MotherboardSpecs](productService.GetDefaultSpecs,mbService.GetSpecs,"motherboard")

	//Laptop Initalization
	laptopRepo := repo.NewLaptopRepository(db)
	laptopService := service.NewLaptopService(laptopRepo)
	laptopGenericService := service.NewComponentService[dto.LaptopResponse, dto.LaptopParams](laptopRepo.GetLaptops, productRepo.GetProductByID,projUtils.LaptopMapping) 
	laptopGenericHandler := handler.NewComponentHandler(laptopGenericService)
	laptopSpecHandler := handler.NewSpecHandler[dto.LaptopSpecs](productService.GetDefaultSpecs,laptopService.GetSpecs,"laptop")

	//Gpus Initalization
	gpuRepo := repo.NewGpuRepository(db)
	gpuService := service.NewGpuService(gpuRepo)
	gpuGenericService := service.NewComponentService[dto.GpuResponse, dto.GpuParams](gpuRepo.GetGpus, productRepo.GetProductByID,projUtils.GpuMapping)
	gpuGenericHandler := handler.NewComponentHandler(gpuGenericService)
	gpuSpecHandler := handler.NewSpecHandler[dto.GpuSpecs](productService.GetDefaultSpecs,gpuService.GetSpecs,"gpu")

	//Fans Initalization
	fanRepo := repo.NewFanRepository(db)
	fanService := service.NewFanService(fanRepo)
	fanGenericService := service.NewComponentService[dto.FanResponse, dto.FanParams](fanRepo.GetFans, productRepo.GetProductByID,projUtils.FanMapping)
	fanGenericHandler := handler.NewComponentHandler(fanGenericService)
	fanSpecHandler := handler.NewSpecHandler[dto.FanSpecs](productService.GetDefaultSpecs,fanService.GetSpecs,"fan")

	//Aio Initalization
	aioRepo := repo.NewAioRepository(db)
	aioService := service.NewAioService(aioRepo)
	aioGenericService := service.NewComponentService[dto.AioResponse, dto.AioParams](aioRepo.GetAios, productRepo.GetProductByID,projUtils.AioMapping) 
	aioGenericHandler := handler.NewComponentHandler(aioGenericService)
	aioSpecHandler:= handler.NewSpecHandler[dto.AioSpecs](productService.GetDefaultSpecs,aioService.GetSpecs,"aio")

	//Cpu Initalization
	cpuRepo := repo.NewCpuRepository(db)
	cpuService := service.NewCpuService(cpuRepo)
	cpuGenericService := service.NewComponentService[dto.CpuResponse, dto.CpuParams](cpuRepo.GetCpus, productRepo.GetProductByID,projUtils.CpuMapping) 
	cpuGenericHandler := handler.NewComponentHandler(cpuGenericService)
	cpuSpecHandler := handler.NewSpecHandler[dto.CpuSpecs](productService.GetDefaultSpecs,cpuService.GetSpecs,"cpu")
	
	//Case Initialization
	caseRepo := repo.NewCaseRepository(db)
	caseService := service.NewCaseService(caseRepo)
	caseGenericService := service.NewComponentService[dto.CaseResponse, dto.CaseParams](caseRepo.GetCases, productRepo.GetProductByID,projUtils.CaseMapping) 
	caseGenericHandler := handler.NewComponentHandler(caseGenericService)
	caseSpecHandler := handler.NewSpecHandler[dto.CaseSpecs](productService.GetDefaultSpecs,caseService.GetSpecs,"case")
	
	//HDD Initalization
	hddRepo := repo.NewHddRepository(db)
	hddService := service.NewHddService(hddRepo)
	hddGenericService := service.NewComponentService[dto.HddResponse, dto.HddParams](hddRepo.GetHdds, productRepo.GetProductByID,projUtils.HddMapping) 
	hddGenericHandler := handler.NewComponentHandler(hddGenericService)
	hddSpecHandler := handler.NewSpecHandler[dto.HddSpecs](productService.GetDefaultSpecs,hddService.GetSpecs,"hdd")

	//Cooler Initalization
	coolerRepo := repo.NewCoolerRepository(db)
	coolerGenericService := service.NewComponentService[dto.CoolerResponse, dto.CoolerParams](coolerRepo.GetCoolers, productRepo.GetProductByID,projUtils.CoolerMapping) 
	coolerService := service.NewCoolerService(coolerRepo)
	coolerGenericHandler := handler.NewComponentHandler(coolerGenericService)
	coolerSpecHandler := handler.NewSpecHandler[dto.CoolerSpecs](productService.GetDefaultSpecs,coolerService.GetSpecs,"cooler")

	//Initiazlie Gin Router
	r := gin.Default()
	r.Use(cors.Default())

	//Http Requests products getters
	products := r.Group("/product")
	{
		products.GET("", productHanlder.GetProducts)
		products.GET("/spec", productHanlder.GetDefaultSpecs)
	}

	//Http Requests component getters	
	//Cooler
	cooler := r.Group("/cooler")
	{
		cooler.GET("", coolerGenericHandler.GetComponents)
		cooler.GET("/spec", coolerSpecHandler.GetComponentSpecs)
		cooler.GET("/:id", coolerGenericHandler.GetComponentByID)
	}
	//Hdd
	hdd := r.Group("/hdd")
	{
		hdd.GET("", hddGenericHandler.GetComponents)
		hdd.GET("/spec", hddSpecHandler.GetComponentSpecs)
		hdd.GET("/:id", hddGenericHandler.GetComponentByID)
	}
	//SSD
	ssd := r.Group("/ssd")
	{
		ssd.GET("", ssdGenericHandler.GetComponents)
		ssd.GET("/spec", ssdSpecHandler.GetComponentSpecs)
		ssd.GET("/:id", ssdGenericHandler.GetComponentByID)
	}
	//RAM
	ram := r.Group("/ram")
	{
		ram.GET("", ramGenericHandler.GetComponents)
		ram.GET("/spec", ramSpecHandler.GetComponentSpecs)
		ram.GET("/:id", ramGenericHandler.GetComponentByID)
	}
	//PSU
	psu := r.Group("/psu")
	{
		psu.GET("", psuGenericHandler.GetComponents)
		psu.GET("/spec", psuSpecHandler.GetComponentSpecs)
		psu.GET("/:id", psuGenericHandler.GetComponentByID)
	}
	//PC Mini
	pcmini := r.Group("/pc-mini")
	{
		pcmini.GET("", pcMiniGenericHandler.GetComponents)
		pcmini.GET("/spec", pcMiniSpecHandler.GetComponentSpecs)
		pcmini.GET("/:id", pcMiniGenericHandler.GetComponentByID)
	}
	//PC
	pc := r.Group("/pc")
	{
		pc.GET("", pcGenericHandler.GetComponents)
		pc.GET("/spec", pcSpecHandler.GetComponentSpecs)
		pc.GET("/:id", pcGenericHandler.GetComponentByID)
	}
	//Motherboard
	motherboard := r.Group("/motherboard")
	{
		motherboard.GET("", mbGenericHandler.GetComponents)
		motherboard.GET("/spec", mbSpecHandler.GetComponentSpecs)
		motherboard.GET("/:id", mbGenericHandler.GetComponentByID)
	}
	//Laptop
	laptop := r.Group("/laptop")
	{
		laptop.GET("", laptopGenericHandler.GetComponents)
		laptop.GET("/spec", laptopSpecHandler.GetComponentSpecs)
		laptop.GET("/:id", laptopGenericHandler.GetComponentByID)
	}
	//GPU
	gpu := r.Group("/gpu")
	{
		gpu.GET("", gpuGenericHandler.GetComponents)
		gpu.GET("/spec", gpuSpecHandler.GetComponentSpecs)
		gpu.GET("/:id", gpuGenericHandler.GetComponentByID)
	}
	//Fan
	fan := r.Group("/fan")
	{
		fan.GET("", fanGenericHandler.GetComponents)
		fan.GET("/spec", fanSpecHandler.GetComponentSpecs)
		fan.GET("/:id", fanGenericHandler.GetComponentByID)
	}
	//AIO
	aio := r.Group("/aio")
	{
		aio.GET("", aioGenericHandler.GetComponents)	
		aio.GET("/spec", aioSpecHandler.GetComponentSpecs)
		aio.GET("/:id", aioGenericHandler.GetComponentByID)
	}
	//CPU
	cpu := r.Group("/cpu")
	{
		cpu.GET("", cpuGenericHandler.GetComponents)
		cpu.GET("/spec", cpuSpecHandler.GetComponentSpecs)
		cpu.GET("/:id", cpuGenericHandler.GetComponentByID)
	}
	//Case
	pcCase := r.Group("/case")
	{
		pcCase.GET("", caseGenericHandler.GetComponents)
		pcCase.GET("/spec", caseSpecHandler.GetComponentSpecs)
		pcCase.GET("/:id", caseGenericHandler.GetComponentByID)
	}
	
	// listen and serve on
	r.Run(":8080") 
}
