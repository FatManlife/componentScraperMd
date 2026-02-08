package dto

type Cooler struct {
	BaseAttrs BaseProduct
	Type string 
	FanRPM int 
	Noise float64 
	Size string 
	Compatibility []string 
}

type CoolerResponse struct {
	Product ProductResponse `json:"product"`
	Type string `json:"type"`
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
	Compatibility []string `json:"compatibility"`
}

type CoolerParams struct {
	DefaultParams ProductParams `form:"default_params"`
	Type string `form:"type"`
	FanRPM int `form:"fan_rpm"`
	Noise float64 `form:"noise"`
	Compatibility []string `form:"compatibility"`
}