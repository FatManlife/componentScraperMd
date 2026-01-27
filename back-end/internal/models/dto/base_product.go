package dto

type BaseProduct struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Website_id int `json:"website_id"`
	Category_id int `json:"category_id"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Website_id int `json:"website_id"`
	Category_id int `json:"category_id"`
}