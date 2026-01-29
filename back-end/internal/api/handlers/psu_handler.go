package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type PsuHandler struct {
	service *service.PsuService
}

func NewPsuHandler(service *service.PsuService) *PsuHandler {
	return &PsuHandler{service: service}
}

func (h *PsuHandler) GetPsusHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var psuParams dto.PsuParams

	if err := c.BindQuery(&psuParams);  err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	psus, err := h.service.GetPsus(ctx, &psuParams)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get PSUs"})
		return
	}

	c.JSON(200, psus)
}