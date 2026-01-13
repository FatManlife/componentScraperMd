package models

type PcMini struct {
	BaseAttrs BaseProduct
	Cpu string `json:"cpu"`
	Gpu string `json:"gpu"`
	Ram string `json:"ram"`
	Storage string `json:"storage"`
}