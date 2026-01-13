package models

type Hdd struct  {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	RotationSpeed int `json:"rotation_speed"`	
	FormFactor string `json:"form_factor"`
}