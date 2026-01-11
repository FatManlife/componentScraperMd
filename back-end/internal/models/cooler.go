package models

type Cooler struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Type string `json:"type"`
	Ilumination string `json:"ilumination"`
	FanRPM int `json:"fan_rpm"`
	Compatibility []string `json:"compatibility"`
	Noise float64 `json:"noise"`
	Size string `json:"size"`

}