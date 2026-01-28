package dto

type Psu struct {
	BaseAttrs BaseProduct
	Power int `json:"power"`
	Efficiency string `json:"efficiency"`
	FormFactor string `json:"form_factor"`
}

type PsuResponse struct {
	Product ProductResponse
	Power int `json:"power"`
	Efficiency string `json:"efficiency"`
	FormFactor string `json:"form_factor"`
}

type PsuParams struct {
	DefaultParams ProductParams 
	Power int 
	Efficiency string 
	FormFactor string 
}