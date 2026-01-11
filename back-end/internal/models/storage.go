package models

type Storage struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Capacity int `json:"capacity"`
	ReadSpeed int `json:"read_jspeed"`
	WriteSpeed int `json:"write_speed"`
	FormFactor string `json:"form_factor"`
	MTBF int `json:"mtbf"`
}