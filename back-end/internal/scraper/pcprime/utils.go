package pcprime

import (
	"regexp"
	"strconv"
)

func CastingFloat64Laptop(s string) float64{
	m := regexp.MustCompile(`(?i)(\d+(\.\d+)?)\s*wh`).FindStringSubmatch(s)
	if len(m) > 1 {
		v, _ := strconv.ParseFloat(m[1], 64)
		return v
	}
	return 0
}