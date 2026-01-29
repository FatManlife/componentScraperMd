package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type SSDHandler struct {
	service *service.SSDService
}

func NewSSDHandler(service *service.SSDService) *SSDHandler {
	return &SSDHandler{service: service}
}

func (h *SSDHandler) GetSsdsHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var ssdParams dto.SsdParams

	if err := c.BindQuery(&ssdParams); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	products, err := h.service.GetSsds(ctx, ssdParams)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}