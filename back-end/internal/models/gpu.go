package models

type Gpu struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Chipset string `json:"chipset"`
	Price float64 `json:"price"`
	Vram int `json:"vram"`
	GpuFrequency int `json:"gpu_frequency"`
	VramFrequency int `json:"vram_frequency"`
}