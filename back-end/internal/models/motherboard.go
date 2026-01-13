package models

type Motherboard struct {
	BaseAttrs BaseProduct
	Chipset string `json:"chipset"`
	Socket string `json:"socket"`
	FormFactor string `json:"form_factor"`
	RamSupport string `json:"ram_support"`
	FormFactorRam string `json:"form_factor_ram"`
}