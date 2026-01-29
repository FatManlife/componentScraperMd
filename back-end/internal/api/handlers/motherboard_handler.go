package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type MotherboardHandler struct {
	service *service.MotherboardService
}

func NewMotherboardHandler(mbService *service.MotherboardService) *MotherboardHandler {
	return &MotherboardHandler{service: mbService}
}	

func (h *MotherboardHandler) GetMotherboards (c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.MotherboardParams

	if err := c.BindQuery(&params);  err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	motherboards, err := h.service.GetMotherboards(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve motherboards"})
		return
	}
	
	c.JSON(200, motherboards)
}