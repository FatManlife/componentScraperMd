package dto

type Laptop struct {
	BaseAttrs BaseProduct
	Cpu string 
	Gpu string 
	Ram int
	Storage int
	Diagonal float64
	Battery float64 
}

type LaptopResponse struct {
	Product ProductResponse
	Cpu string `json:"cpu"`
	Gpu string`json:"gpu"`
	Ram int`json:"ram"`
	Storage int `json:"storage"`
	Diagonal float64 `json:"diagonal"`
	Battery float64 `json:"battery"`
}

type LaptopParams struct {
	DefaultParams ProductParams 
	Cpu []string `form:"cpu"` 
	Gpu []string `form:"gpu"`
	Ram []int `form:"ram"`
	Storage []int `form:"storage"`
	Diagonal []float64 `form:"diagonal"`
}