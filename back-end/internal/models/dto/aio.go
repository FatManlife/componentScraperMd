package dto

type Aio struct {
	BaseAttrs BaseProduct 
	Diagonal string 
	Cpu string 
	Ram int 
	Storage int 
	Gpu string 
}

type AioResponse struct {
	Product ProductResponse
	Diagonal string `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram int `json:"ram"`
	Storage int `json:"storage"`
	Gpu string `json:"gpu"`
}

type AioParams struct {
	DefaultParams ProductParams
	Diagonal []string `form:"diagonal"`
	Cpu []string `form:"cpu"`
	Ram []int `form:"ram"`
	Storage []int `form:"storage"`
	Gpu []string `form:"gpu"`
}