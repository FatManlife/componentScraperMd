package dto

type Motherboard struct {
	BaseAttrs BaseProduct
	Chipset string 
	Socket string 
	FormFactor string 
	RamSupport string 
	FormFactorRam string 
}

type MotherboardResponse struct {
	Product ProductResponse
	Chipset string `json:"chipset"`
	Socket string `json:"socket"`
	FormFactor string `json:"form_factor"`
	RamSupport string `json:"ram_support"`
	FormFactorRam string `json:"form_factor_ram"`
}

type MotherboardParams struct {
	DefaultParams ProductParams 
	Chipset []string `form:"chipset"`
	Socket []string `form:"socket"`
	FormFactor []string `form:"form_factor"`
}