package dto

type Cpu struct {
	BaseAttrs  BaseProduct
	Cores      int     
	Threads    int     
	BaseClock  float64 
	BoostClock float64 
	Socket     string  
	Tdp        int     
}

type CpuResponse struct {
	Product    ProductResponse `json:"product"`
	Cores      int     `json:"cores"`
	Threads    int     `json:"threads"`
	BaseClock  float64 `json:"base_clock"`
	BoostClock float64 `json:"boost_clock"`
	Socket     string  `json:"socket"`
	Tdp        int     `json:"tdp"`
}

type CpuParams struct {
	DefaultParams ProductParams `form:"default_params"`
	Cores      []int	`form:"cores"` 
	Threads    []int     `form:"threads"`
	BaseClock  []float64 `form:"base_clock"`
	BoostClock []float64 `form:"boost_clock"`
	Socket     []string  `form:"socket"`
}