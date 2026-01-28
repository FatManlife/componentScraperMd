package dto

type PcMini struct {
	BaseAttrs BaseProduct
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
}

type PcMiniResponse struct {
	Product ProductResponse
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
}

type PcMiniParams struct {
	DefaultParams ProductParams 
	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
}