package dto

type Fan struct {
	BaseAttrs BaseProduct
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
}

type FanResponse struct {
	Product ProductResponse
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
}	