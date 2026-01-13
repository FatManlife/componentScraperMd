package models

type Case struct {
	BaseAttrs BaseProduct
	Format string `json:"format"`
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
}
