package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type RamHandler struct {
	service *service.RamService
}

func NewRamHandler(service *service.RamService) *RamHandler {
	return &RamHandler{service: service}
}

func (h *RamHandler) GetRams(c *gin.Context){
	ctx := c.Request.Context()

	var params dto.RamParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	products, err := h.service.GetRams(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}