package dto

type Ram struct {
	BaseAttrs BaseProduct
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	Compatibility string `json:"compatibility"`
	Configuration string `json:"configuration"`
}

type RamResponse struct {
	Product ProductResponse
	Capacity int `json:"capacity"`
	Speed int `json:"speed"`
	Type string `json:"type"`
	Compatibility string `json:"compatibility"`
	Configuration string `json:"configuration"`
}

type RamParams struct {
	DefaultParams ProductParams 
	Capacity int 
	Speed int 
	Type string 
	Compatibility string 
	Configuration string 
}