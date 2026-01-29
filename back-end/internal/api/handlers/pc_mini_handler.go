package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type PcMiniHandler struct {
	service *service.PcMiniService
} 

func NewPcMiniHandler(service *service.PcMiniService) *PcMiniHandler {
	return &PcMiniHandler{service: service}
}

func (h *PcMiniHandler) GetPcMinis(c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.PcMiniParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}
	
	pcMinis, err := h.service.GetPcMinis(ctx, params)
	
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pcMinis)
}