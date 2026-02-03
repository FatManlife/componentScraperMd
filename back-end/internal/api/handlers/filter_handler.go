package handler

import (
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type FilterHandler struct {
	serviec *service.FilterService
}

func NewFilterHandler(service *service.FilterService) *FilterHandler {
	return &FilterHandler{serviec: service}
}

func (h *FilterHandler) GetDefaultFilters(c *gin.Context) {
	category := c.Query("category")

	filters, err := h.serviec.GetDefaultFilters(c.Request.Context(), category)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, filters)
}

