package dto

type Gpu struct {
	BaseAttrs BaseProduct
	Chipset string 
	Vram int 
	GpuFrequency int 
	VramFrequency int
}

type GpuResponse struct {
	Product ProductResponse `json:"product"`
	Chipset string `json:"chipset"`
	Vram int `json:"vram"`
	GpuFrequency int `json:"gpu_frequency"`
	VramFrequency int `json:"vram_frequency"`
}

type GpuParams struct {
	DefaultParams ProductParams `form:"default_params"`
	Chipset []string `form:"chipset"`
	MinVram int `form:"min_vram"`
	MaxVram int `form:"max_vram"`
	MinGpuFrequency int `form:"min_gpu_frequency"`
	MaxGpuFrequency int `form:"max_gpu_frequency"`
	MinVramFrequency int `form:"min_vram_frequency"`
	MaxVramFrequency int `form:"max_vram_frequency"`
}