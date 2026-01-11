package models

type Motherboard struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`	
	Chipset string `json:"chipset"`
	Socket string `json:"socket"`
	FormFactor string `json:"form_factor"`
	RamSupport string `json:"ram_support"`
	NetWork string `json:"networking"`
	Blueetooth string `json:"bluetooth"`
	FormFactorRam string `json:"form_factor_ram"`
}