package dto

type Gpu struct {
	BaseAttrs BaseProduct
	Chipset string `json:"chipset"`
	Vram int `json:"vram"`
	GpuFrequency int `json:"gpu_frequency"`
	VramFrequency int `json:"vram_frequency"`
}

type GpuResponse struct {
	Product ProductResponse
	Chipset string `json:"chipset"`
	Vram int `json:"vram"`
	GpuFrequency int `json:"gpu_frequency"`
	VramFrequency int `json:"vram_frequency"`
}

type GpuParams struct {
	DefaultParams ProductParams 
	Chipset string `form:"chipset"`
	Vram int `form:"vram"`
	GpuFrequency int `form:"gpu_frequency"`
	VramFrequency int `form:"vram_frequency"`
}