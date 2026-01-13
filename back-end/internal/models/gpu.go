package models

type Gpu struct {
	BaseAttrs BaseProduct
	Chipset string `json:"chipset"`
	Vram int `json:"vram"`
	GpuFrequency int `json:"gpu_frequency"`
	VramFrequency int `json:"vram_frequency"`
}