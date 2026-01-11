package models

type Ram struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	FormFactor string `json:"form_factor"`
	Configuration string `json:"configuration"`
}