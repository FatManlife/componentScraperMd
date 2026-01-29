package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type CaseHandler struct {
	service *service.CaseService
}

func NewCaseHandler(service *service.CaseService) *CaseHandler {
	return &CaseHandler{service: service}
}

func (h *CaseHandler) GetCases(c *gin.Context){
	ctx := c.Request.Context()
	
	var params dto.CaseParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	cases, err := h.service.GetCases(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get cases"})
		return
	}

	c.JSON(200, cases)
}