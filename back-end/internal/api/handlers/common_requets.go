package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type commonFilters struct { 
	website string
	brand   string
	min    	float64
	max     float64
	sort    string
	limit  	int 
	after  	int
} 

func extractFilters(c *gin.Context, comStruct *commonFilters)  error{
	comStruct.website = c.DefaultQuery("website", "")
	comStruct.brand = c.DefaultQuery("brand", "")
	comStruct.sort = c.DefaultQuery("sort", "")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit <= 0 {
		c.JSON(400, gin.H{"error": "limit must be a positive integer"})
		return err
	}
	comStruct.limit = limit

	after, err := strconv.Atoi(c.DefaultQuery("after", "0"))
	if err != nil || after < 0 {
		c.JSON(400, gin.H{"error": "after must be a positive integer"})
		return err
	}
	comStruct.after = after

	min, err := strconv.ParseFloat(c.DefaultQuery("min", "0"), 64)
	if err != nil || min < 0 {
		c.JSON(400, gin.H{"error": "min must be a positive number"})
		return err
	}
	comStruct.min = min

	max, err := strconv.ParseFloat(c.DefaultQuery("max", "0"), 64)
	if err != nil || max < 0 {
		c.JSON(400, gin.H{"error": "max must be a positive number"})
		return err
	}
	comStruct.max = max
	
	return nil
}