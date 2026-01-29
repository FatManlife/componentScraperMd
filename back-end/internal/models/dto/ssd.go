package dto

type Ssd struct {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	ReadingSpeed int `json:"reading_speed"`	
	WritingSpeed int `json:"writing_speed"`	
	FormFactor string `json:"form_factor"`
}

type SsdResponse struct {
	Product ProductResponse`json:"product"`
	Capacity       int `json:"capacity"` 
	ReadingSpeed  int `json:"reading_speed"` 
	WritingSpeed  int `json:"writing_speed"` 
	FormFactor    string `json:"form_factor"`
}

type SsdParams struct {
	DefaultParams ProductParams 
	Capacity       int `form:"capacity"`
	ReadingSpeed  int `form:"reading_speed"`
	WritingSpeed  int `form:"writing_speed"`
	FormFactor    string `form:"form_factor"`
}