package dto

type Hdd struct  {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	RotationSpeed int `json:"rotation_speed"`	
	FormFactor string `json:"form_factor"`
}

type HddResponse struct {
	Product ProductResponse
	Capacity int `json:"capacity"`
	RotationSpeed int `json:"rotation_speed"`	
	FormFactor string `json:"form_factor"`
}

type HddParams struct {
	DefaultParams ProductParams 
	Capacity int 
	RotationSpeed int 
	FormFactor string 
}