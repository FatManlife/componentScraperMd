package dto

type Psu struct {
	BaseAttrs BaseProduct
	Power int 
	Efficiency string 
	FormFactor string 
}

type PsuResponse struct {
	Product ProductResponse
	Power int `json:"power"`
	Efficiency string `json:"efficiency"`
	FormFactor string `json:"form_factor"`
}

type PsuParams struct {
	DefaultParams ProductParams 
	MinPower int `form:"min_power"`
	MaxPower int `form:"max_power"`
	Efficiency []string `form:"efficiency"`
	FormFactor []string `form:"form_factor"`
}