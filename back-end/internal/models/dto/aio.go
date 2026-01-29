package dto

type Aio struct {
	BaseAttrs BaseProduct 
	Diagonal string `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Gpu string `json:"gpu"`
}

type AioResponse struct {
	Product ProductResponse
	Diagonal string `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Gpu string `json:"gpu"`
}

type AioParams struct {
	DefaultParams ProductParams
	Diagonal string `form:"diagonal"`
	Cpu string `form:"cpu"`
	Ram string `form:"ram"`
	Storage string `form:"storage"`
	Gpu string `form:"gpu"`
}
