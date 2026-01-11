package models

type Psu struct {
	Name string `json:"name" selector:"a.xp-title"`
	ImageURL string `json:"image_url" selector:"a.img-wrap img" attr:"src"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`	
	Power int `json:"power"`
	Efficiency string `json:"efficiency"`
	Modularity string `json:"modularity"`
	FormFactor string `json:"form_factor"`
}