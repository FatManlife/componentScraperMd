package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type HddHandler struct {
	service *service.HddService
}

func NewHddHandler(hddService *service.HddService) *HddHandler {
	return &HddHandler{service: hddService}
}

func (h *HddHandler) GetHdds(c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.HddParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	hdds, err := h.service.GetHdds(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve hdds"})
		return
	}

	c.JSON(200, hdds)
}