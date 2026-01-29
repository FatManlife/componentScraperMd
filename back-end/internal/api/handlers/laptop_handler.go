package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type LaptopHandler struct {
	service *service.LaptopService
}

func NewLaptopHandler(laptopService *service.LaptopService) *LaptopHandler {
	return &LaptopHandler{service: laptopService}
}

func (h *LaptopHandler) GetLaptops(c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.LaptopParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	laptops, err := h.service.GetLaptops(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve laptops"})
		return
	}

	c.JSON(200, laptops)
}