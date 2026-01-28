package handler

import (
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

	var filters commonFilters

	if err := extractFilters(c, &filters); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	products, err := h.service.GetAios(ctx, filters.limit, filters.website, filters.after, filters.brand, filters.min, filters.max, filters.sort)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}