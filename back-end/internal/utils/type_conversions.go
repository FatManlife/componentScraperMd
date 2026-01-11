package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var patternNumber = regexp.MustCompile(`[0-9\s,]+\.?\d*`)

func CastInt(s string) int{
	newStr, err := strconv.Atoi(strings.ReplaceAll(patternNumber.FindString(s)," ", ""))

	if err != nil {
		fmt.Println(err)
	}

	return newStr
}

func CastFloat64(s string) float64{
	newStr, err := strconv.ParseFloat(strings.ReplaceAll(patternNumber.FindString(s)," ", ""),64)

	if err != nil {
		fmt.Println(err)
	}

	return newStr
}