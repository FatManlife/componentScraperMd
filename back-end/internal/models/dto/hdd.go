package dto

type Hdd struct  {
	BaseAttrs BaseProduct
	Capacity int 
	RotationSpeed int 	
	FormFactor string 
}

type HddResponse struct {
	Product ProductResponse
	Capacity int `json:"capacity"`
	RotationSpeed int `json:"rotation_speed"`	
	FormFactor string `json:"form_factor"`
}

type HddParams struct {
	DefaultParams ProductParams 
	MinCapacity int `form:"min_capacity"`
	MaxCapacity int `form:"max_capacity"`
	MinRotationSpeed int `form:"min_rotation_speed"`
	MaxRotationSpeed int `form:"max_rotation_speed"`
	FormFactor []string `form:"form_factor"`
}