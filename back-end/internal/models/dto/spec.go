package dto

type DefaultSpecs struct {
	Websites []string `json:"websites"`
	Prices []float64 `json:"prices"`
	Order []string `json:"order"`
}

type AioSpecs struct {
	Diagonal []float64`form:"diagonal"`
	Cpu []string `form:"cpu"`
	Ram []int `form:"ram"`
	Storage []int `form:"storage"`
	Gpu []string `form:"gpu"`
}

type CoolerSpecs struct {
	Type []string `form:"type"`
	FanRPM []int `form:"fan_rpm"`
	Noise []float64 `form:"noise"`
	Compatibility []string `form:"compatibility"`
}

type CaseSpecs struct {
	Format []string `form:"format"`
	MotherboardFormFactor []string `form:"motherboard_form_factor"`
}

type CpuSpecs struct {
	Cores      []int	`form:"cores"` 
	Threads    []int     `form:"threads"`
	BaseClock  []float64 `form:"base_clock"`
	BoostClock []float64 `form:"boost_clock"`
	Socket     []string  `form:"socket"`
}

type FanSpecs struct {
	FanRPM []int 	`form:"fan_rpm"`
	Noise []float64 `form:"noise"`
}

type GpuSpecs struct {
	Chipset []string `json:"chipset"`
	Vram []int `json:"vram"`
	GpuFrequency []int `json:"gpu_frequency"`
	VramFrequency []int `json:"vram_frequency"`
}

type HddSpecs struct {
	Capacity []int `json:"capacity"`
	RotationSpeed []int `json:"rotation_speed"`	
	FormFactor []string `json:"form_factor"`
}

type LaptopSpecs struct {
	Cpu []string `json:"cpu"`
	Gpu []string `json:"gpu"`
	Ram []int `json:"ram"`
	Storage []int `json:"storage"`
	Diagonal []float64 `json:"diagonal"`
}

type MotherboardSpecs struct {
	Chipset []string `form:"chipset"`
	Socket []string `form:"socket"`
	FormFactor []string `form:"form_factor"`
}

type PcSpecs struct {
	Cpu []string `json:"cpu"`
	Gpu []string `json:"gpu"`
	Ram []int `json:"ram"`
	Storage []int `json:"storage"`
}

type PsuSpecs struct {
	Power []int `json:"power"`
	Efficiency []string `json:"efficiency"`
	FormFactor []string `json:"form_factor"`
}

type RamSpecs struct {
	Capacity []int `json:"capacity"`
	Speed []int `json:"speed"`
	Type []string `json:"type"`
	Compatibility []string `json:"compatibility"`
	Configuration []int`json:"configuration"`
}

type SsdSpecs struct {
	Capacity      []int `json:"capacity"` 
	ReadingSpeed  []int `json:"reading_speed"` 
	WritingSpeed  []int `json:"writing_speed"` 
	FormFactor    []string `json:"form_factor"`
}