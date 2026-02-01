package dto

type Ram struct {
	BaseAttrs BaseProduct
	Capacity int 
	Speed int 
	Type string 
	Compatibility string 
	Configuration string 
}

type RamResponse struct {
	Product ProductResponse
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	Compatibility string `json:"compatibility"`
	Configuration string `json:"configuration"`
}

type RamParams struct {
	DefaultParams ProductParams 
	MinCapacity int `form:"min_capacity"`
	MaxCapacity int `form:"max_capacity"`
	MinSpeed int`form:"min_speed"`
	MaxSpeed int`form:"max_speed"`
	Type []string `form:"type"`
	Compatibility []string `form:"compatibility"`
	Configuration []string `form:"configuration"`
}