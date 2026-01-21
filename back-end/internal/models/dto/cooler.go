package dto

type Cooler struct {
	BaseAttrs BaseProduct
	Type string `json:"type"`
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
	Compatibility []string `json:"compatibility"`
}