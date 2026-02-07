package dto

type Aio struct {
	BaseAttrs BaseProduct 
	Diagonal float64
	Cpu string 
	Ram int 
	Storage int 
	Gpu string 
}

type AioResponse struct {
	Product ProductResponse
	Diagonal float64 `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram int `json:"ram"`
	Storage int `json:"storage"`
	Gpu string `json:"gpu"`
}

type AioParams struct {
	DefaultParams ProductParams
	Diagonal []float64`form:"diagonal"`
	Cpu []string `form:"cpu"`
	Ram []int `form:"ram"`
	Storage []int `form:"storage"`
	Gpu []string `form:"gpu"`
}