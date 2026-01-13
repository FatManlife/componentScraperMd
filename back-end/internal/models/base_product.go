package models

type BaseProduct struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Website_id int `json:"website_id"`
}