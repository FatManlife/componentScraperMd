package handler

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type SpecHandler[T any] struct {
	getDeafultSpecs func(ctx context.Context, category string) (dto.DefaultSpecs, error)
	getSpecs func (ctx context.Context) (T, error)
	category string
}

func NewSpecHandler[T any](getDefaultSpecs func(ctx context.Context, category string) (dto.DefaultSpecs, error), getSpecs func (ctx context.Context) (T, error), category string) *SpecHandler[T] {
	return &SpecHandler[T]{
		getDeafultSpecs: getDefaultSpecs,
		getSpecs: getSpecs,
		category: category,
	}
}

func (h *SpecHandler[T]) GetComponentSpecs(c *gin.Context) {
	defaultSpecs, err := h.getDeafultSpecs(c.Request.Context(), h.category)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	specificSpecs, err := h.getSpecs(c.Request.Context())

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	specs := gin.H{
		"defaultSpecs": defaultSpecs,
		"specificSpecs": specificSpecs,
	}

	c.JSON(200, specs)
}