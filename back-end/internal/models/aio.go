package models

type Aio struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Diagonal string `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Gpu string `json:"gpu"`
}
