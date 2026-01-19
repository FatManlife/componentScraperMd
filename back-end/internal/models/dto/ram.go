package dto

type Ram struct {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	Compatibility string `json:"compatibility"`
	Configuration string `json:"configuration"`
}