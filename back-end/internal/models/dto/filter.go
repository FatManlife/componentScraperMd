package dto

type DefaultFilters struct {
	Websites []string `json:"websites"`
	Prices []float64 `json:"prices"`
	Order []string `json:"order"`
}