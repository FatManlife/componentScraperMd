package dto

type BaseProduct struct {
	Name string 
	ImageURL string 
	Brand string 
	Price float64 
	Url string 
	Category string 
	Website_id int 
}

type ProductsResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Category string	`json:"category"`
	Website_id int `json:"website_id"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Category string	`json:"category"`
	Website_id int `json:"website_id"`
}

type ProductParams struct {
	Name string `form:"name"`
	Limit   int     `form:"limit"`
	Website []string  `form:"website"`
	After   int     `form:"after"`
	Min     float64 `form:"min"`
	Max     float64 `form:"max"`
	Order   string  `form:"order"`
}