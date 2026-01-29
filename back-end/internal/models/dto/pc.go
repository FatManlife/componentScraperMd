package dto

type Pc struct {
	BaseAttrs BaseProduct
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Motherboard string `json:"motherboard"`
	Psu string `json:"psu"`
	Case string `json:"case"`
}

type PcResponse struct {
	Product ProductResponse
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Motherboard string `json:"motherboard"`
	Psu string `json:"psu"`
	Case string `json:"case"`
}

type PcParams struct {
	DefaultParams ProductParams 
	Cpu string `form:"cpu"`
	Gpu string `form:"gpu"`
	Ram string `form:"ram"`
	Storage string `form:"storage"`
	Motherboard string `form:"motherboard"`
	Psu string `form:"psu"`
	Case string `form:"case"`
}