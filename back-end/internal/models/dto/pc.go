package dto

type Pc struct {
	BaseAttrs BaseProduct
	Cpu string 
	Gpu string 
	Ram int
	Storage int
	Motherboard string 
	Psu string 
	Case string 
}

type PcResponse struct {
	Product ProductResponse
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram int`json:"ram"`
	Storage int`json:"storage"`
	Motherboard string `json:"motherboard"`
	Psu string `json:"psu"`
	Case string `json:"case"`
}

type PcParams struct {
	DefaultParams ProductParams 
	Cpu []string `form:"cpu"`
	Gpu []string `form:"gpu"`
	Ram []int `form:"ram"`
	Storage []int `form:"storage"`
}