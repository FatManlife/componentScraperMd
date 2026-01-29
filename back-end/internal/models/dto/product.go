package dto

type BaseProduct struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Website_id int `json:"website_id"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Website_id int `json:"website_id"`
}

type ProductParams struct {
	Limit   int     `form:"limit"`
	Website string  `form:"website"`
	After   int     `form:"after"`
	Brand   string  `form:"brand"`
	Min     float64 `form:"min"`
	Max     float64 `form:"max"`
	Order   string  `form:"order"`
}