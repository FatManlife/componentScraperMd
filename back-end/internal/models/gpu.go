package models

type Gpu struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	ChipsetBrand string `json:"chipset_brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Vram int `json:"vram"`
	GpuSpeed int `json:"gpu_speed"`
	VramSpeed int `json:"vram_speed"`
}