package models

type Case struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Format string `json:"format"`
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
}
