package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type AioHandler struct {
	service *service.AioService	
}

func NewAioHandler(service *service.AioService) *AioHandler {
	return &AioHandler{service: service}
}

func (h *AioHandler) GetAios(c *gin.Context) {
	ctx := c.Request.Context()

	var aioParams dto.AioParams

	if err := extractProductParams(c, &aioParams.DefualtParams); err != nil {
		return
	}

	aioParams.Diagonal = c.DefaultQuery("diagonal","")
	aioParams.Ram = c.DefaultQuery("ram","")
	aioParams.Cpu = c.DefaultQuery("cpu","")
	aioParams.Gpu = c.DefaultQuery("gpu","")
	aioParams.Storage = c.DefaultQuery("storage","")

	products, err := h.service.GetAios(ctx, aioParams)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}