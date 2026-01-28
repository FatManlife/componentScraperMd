package handler

import (
	"strconv"

	"github.com/FatManlife/component-finder/back-end/internal/models/dto"
	"github.com/gin-gonic/gin"
)

func extractProductParams(c *gin.Context, productParams *dto.ProductParams)  error{
	productParams.Website = c.DefaultQuery("website", "")
	productParams.Brand = c.DefaultQuery("brand", "")
	productParams.Order = c.DefaultQuery("sort", "")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit <= 0 {
		c.JSON(400, gin.H{"error": "limit must be a positive integer"})
		return err
	}
	productParams.Limit = limit

	after, err := strconv.Atoi(c.DefaultQuery("after", "0"))
	if err != nil || after < 0 {
		c.JSON(400, gin.H{"error": "after must be a positive integer"})
		return err
	}
	productParams	.After = after

	min, err := strconv.ParseFloat(c.DefaultQuery("min", "0"), 64)
	if err != nil || min < 0 {
		c.JSON(400, gin.H{"error": "min must be a positive number"})
		return err
	}
	productParams.Min = min

	max, err := strconv.ParseFloat(c.DefaultQuery("max", "0"), 64)
	if err != nil || max < 0 {
		c.JSON(400, gin.H{"error": "max must be a positive number"})
		return err
	}
	productParams.Max = max
	
	return nil
}