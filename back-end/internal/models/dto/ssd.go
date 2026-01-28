package dto

type Ssd struct {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	ReadingSpeed int `json:"reading_speed"`	
	WritingSpeed int `json:"writing_speed"`	
	FormFactor string `json:"form_factor"`
}

type SsdResponse struct {
	ProductResponse int `json:"product"`
	Capacity       int `json:"capacity"` 
	ReadingSpeed  int `json:"reading_speed"` 
	WritingSpeed  int `json:"writing_speed"` 
	FormFactor    string `json:"form_factor"`
}

type SsdParams struct {
	DefaultParams ProductParams 
	Capacity       int 
	ReadingSpeed  int 
	WritingSpeed  int 
	FormFactor    string
}