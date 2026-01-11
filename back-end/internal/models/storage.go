package models

type Storage struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Capacity int `json:"capacity"`
	FormFactor string `json:"form_factor"`
}