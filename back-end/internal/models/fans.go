package models

type Fan struct {
	BaseAttrs BaseProduct
	Type string `json:"type"`
	Ilumination string `json:"ilumination"`
	FanRPM int `json:"fan_rpm"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`
}