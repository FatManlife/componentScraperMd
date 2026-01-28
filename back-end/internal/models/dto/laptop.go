package dto

type Laptop struct {
	BaseAttrs BaseProduct
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Diagonal string `json:"diagonal"`
	Battery float64 `json:"battery"`
}

type LaptopResponse struct {
	Product ProductResponse
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Diagonal string `json:"diagonal"`
	Battery float64 `json:"battery"`
}

type LaptopParams struct {
	DefaultParams ProductParams 
	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
	Diagonal string 
	Battery float64 
}