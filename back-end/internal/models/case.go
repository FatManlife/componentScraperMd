package models

type Case struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Stock int `json:"stock"` 
	Price float64 `json:"price"`
	Format string `json:"format"`
	Material string `json:"material"`
	Cooling string `json:"cooling"`
	PsuLocation string `json:"psu_location"`
	GpuMaxLength string `json:"gpu_max_length"`	
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
	PsuFormFactor string `json:"psu_form_factor"`	
}
