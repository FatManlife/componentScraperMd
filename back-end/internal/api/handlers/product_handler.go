package handler

import (
	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.ProductParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}
	
	products, err := h.service.GetProducts(ctx, params)
	
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}
