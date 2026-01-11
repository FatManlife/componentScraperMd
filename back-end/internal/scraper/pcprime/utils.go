package pcprime

import (
	"regexp"
	"strconv"
)

func castingFloat64Laptop(s string) float64{
	m := regexp.MustCompile(`(?i)(\d+(\.\d+)?)\s*Вт\*ч`).FindStringSubmatch(s)
	if len(m) > 1 {
		v, _ := strconv.ParseFloat(m[1], 64)
		return v
	}
	return 0
}

func CastingCastIntFan(s string) int{
	m := regexp.MustCompile(`(?i)(\d+)\s*rpm`).FindAllStringSubmatch(s, -1)
	if len(m) > 0 {
		v, _ := strconv.Atoi(m[len(m)-1][1])
		return v
	}
	return 0
}