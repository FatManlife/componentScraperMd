package dto

type PcMini struct {
	BaseAttrs BaseProduct
	Cpu string 
	Gpu string 
	Ram int
	Storage int 
}

type PcMiniResponse struct {
	Product ProductResponse `json:"product"`
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram int `json:"ram"`
	Storage int `json:"storage"`
}


