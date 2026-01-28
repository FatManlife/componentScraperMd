package dto

type Case struct {
	BaseAttrs BaseProduct
	Format string `json:"format"`
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
}

type CaseResponse struct {
	Product ProductResponse
	Format string `json:"format"`
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
}

type CaseParams struct {
	DefaultParams ProductParams 
	Format string 
	MotherboardFormFactor string 
}