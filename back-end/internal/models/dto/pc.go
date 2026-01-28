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
	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
	Motherboard string 
	Psu string 
	Case string 
}