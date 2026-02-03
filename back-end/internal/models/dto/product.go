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
	Website string `json:"website"`
	WebsiteImg string `json:"website_image"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Brand string `json:"brand"`
	Price float64 `json:"price"`
	Url string `json:"url"`
	Category string	`json:"category"`
	Website string `json:"website"`
	WebsiteImg string `json:"website_image"`
}

type ProductParams struct {
	Name string `form:"name"`
	Website []string  `form:"website"`
	Offset int     `form:"offset"`
	Min     float64 `form:"min"`
	Max     float64 `form:"max"`
	Order   string  `form:"order"`
}