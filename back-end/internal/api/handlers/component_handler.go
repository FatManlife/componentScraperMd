package handler

import (
	"strconv"

	service "github.com/FatManlife/component-finder/back-end/internal/services"
	"github.com/gin-gonic/gin"
)

type ComponentHandler[T any, P any] struct {
    service *service.ComponentService[T, P]
}

func NewComponentHandler[T any, P any](service *service.ComponentService[T, P]) *ComponentHandler[T, P] {
    return &ComponentHandler[T, P]{service: service}
}

func (h *ComponentHandler[T, P]) GetComponents(c *gin.Context) {
	ctx := c.Request.Context()

	var params P

	if err := c.BindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	products, count, err := h.service.GetComponents(ctx, params)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"products": products, "count": count})
}

func (h *ComponentHandler[T, P]) GetComponentByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err :=  strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID parameter"})
		return
	}

	product, err := h.service.GetComponentByID(ctx, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}
