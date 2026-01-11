package collector

import (
	"github.com/gocolly/colly"
)

func NewContext(key string, value string) *colly.Context{
	ctx := colly.NewContext()
	ctx.Put(key,value)
	return ctx
}