package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type PcHandler struct {
	service *service.PcService
}

func NewPcHandler(service *service.PcService) *PcHandler {
	return &PcHandler{service: service}
}

func (h *PcHandler) GetPcs(c *gin.Context){
	ctx := c.Request.Context()

	var params dto.PcParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}
	
	pcs, err := h.service.GetPcs(ctx, params)
	
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pcs)
}