package dto

type Ssd struct {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	ReadingSpeed int `json:"reading_speed"`	
	WritingSpeed int `json:"writing_speed"`	
	FormFactor string `json:"form_factor"`
}