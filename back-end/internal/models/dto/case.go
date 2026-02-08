package dto

type Case struct {
	BaseAttrs BaseProduct
	Format string
	MotherboardFormFactor string	
}

type CaseResponse struct {
	Product ProductResponse `json:"product"`
	Format string `json:"format"`
	MotherboardFormFactor string `json:"motherboard_form_factor"`	
}

type CaseParams struct {
	DefaultParams ProductParams `form:"default_params"`
	Format []string `form:"format"`
	MotherboardFormFactor []string `form:"motherboard_form_factor"`
}