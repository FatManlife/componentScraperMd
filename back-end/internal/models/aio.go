package models

type Aio struct {
	BaseAttrs BaseProduct 
	Diagonal string `json:"diagonal"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
	Gpu string `json:"gpu"`
}
