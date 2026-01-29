package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type GpuHandler struct {
	service *service.GpuService
}

func NewGpuHandler(service *service.GpuService) *GpuHandler {
	return &GpuHandler{service: service}
}

func (h *GpuHandler) GetGpus(c *gin.Context){
	ctx := c.Request.Context()

	var params dto.GpuParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	gpus, err := h.service.GetGpus(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve GPUs"})
		return
	}

	c.JSON(200, gpus)	
}