package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type cpuHandler struct {
	service *service.CpuService
}

func NewCpuHandler(service *service.CpuService) *cpuHandler {
	return &cpuHandler{service: service}
}

func (h *cpuHandler) GetCpus(c *gin.Context){
	ctx := c.Request.Context()
	
	var params dto.CpuParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	cpus, err := h.service.GetCpus(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get cpus"})
		return
	}

	c.JSON(200, cpus)
}	