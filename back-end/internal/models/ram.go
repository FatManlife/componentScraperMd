package models

type Ram struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	Compatibility string `json:"compatibility"`
	Configuration string `json:"configuration"`
}