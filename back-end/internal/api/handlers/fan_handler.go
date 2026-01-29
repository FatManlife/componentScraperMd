package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type FanHandler struct {
	service *service.FanService
}

func NewFanHandler(service *service.FanService) *FanHandler {
	return &FanHandler{service: service}
}

func (h *FanHandler) GetFans(c *gin.Context){
	ctx := c.Request.Context()
	
	var params dto.FanParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	fans, err := h.service.GetFans(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get fans"})
		return
	}

	c.JSON(200, fans)
}