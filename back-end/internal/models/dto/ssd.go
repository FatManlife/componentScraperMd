package dto

type Ssd struct {
	BaseAttrs BaseProduct
	Capacity int 
	ReadingSpeed int 
	WritingSpeed int 
	FormFactor string 
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
	MinCapacity       int `form:"min_capacity"`
	MaxCapacity       int `form:"max_capacity"`
	MinReadingSpeed  int `form:"min_reading_speed"`
	MaxReadingSpeed  int `form:"max_reading_speed"`
	MinWritingSpeed  int `form:"min_writing_speed"`
	MaxWritingSpeed  int `form:"max_writing_speed"`
	FormFactor    []string `form:"form_factor"`
}