package models

type Motherboard struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`	
	Chipset string `json:"chipset"`
	CpuSupport string `json:"cpu_support"`
	Socket string `json:"socket"`
	FormFactor string `json:"form_factor"`
	RamSupport string `json:"ram_support"`
	NetWork string `json:"networking"`
	Blueetooth bool `json:"bluetooth"`
	FormFactorRam string `json:"form_factor_ram"`
}