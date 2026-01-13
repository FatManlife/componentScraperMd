package models

type Psu struct {
	BaseAttrs BaseProduct
	Power int `json:"power"`
	Efficiency string `json:"efficiency"`
	Modularity string `json:"modularity"`
	FormFactor string `json:"form_factor"`
}