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
	DefualtParams ProductParams
	Diagonal string
	Cpu string
	Ram string
	Storage string
	Gpu string
}
