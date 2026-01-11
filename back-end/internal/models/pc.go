package models

type Pc struct {
	Name string  `json:"name" selector:"a.xp-title"`
	ImageURL string `json:"image_url" selector:"a.img-wrap img" attr:"src"`
	Price float64 `json:"price"`
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Motherboard string `json:"motherboard"`
	Psu string `json:"psu"`
	Case string `json:"case"`
}