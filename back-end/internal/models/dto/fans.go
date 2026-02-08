package dto

type Fan struct {
	BaseAttrs BaseProduct
	FanRPM int 
	Noise float64
	Size string 
}

type FanResponse struct {
	Product ProductResponse `json:"product"`
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
}	

type FanParams struct {
	DefaultParams ProductParams `form:"default_params"`
	MinFanRPM int 	`form:"min_fan_rpm"`
	MaxFanRPM int 	`form:"max_fan_rpm"`
	MinNoise float64 `form:"min_noise"`
	MaxNoise float64 `form:"max_noise"`
}