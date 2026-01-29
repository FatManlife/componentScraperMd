package dto

type Cpu struct {
	BaseAttrs  BaseProduct
	Cores      int     `json:"cores"`
	Threads    int     `json:"threads"`
	BaseClock  float64 `json:"base_clock"`
	BoostClock float64 `json:"boost_clock"`
	Socket     string  `json:"socket"`
	Tdp        int     `json:"tdp"`
}

type CpuResponse struct {
	Product    ProductResponse
	Cores      int     `json:"cores"`
	Threads    int     `json:"threads"`
	BaseClock  float64 `json:"base_clock"`
	BoostClock float64 `json:"boost_clock"`
	Socket     string  `json:"socket"`
	Tdp        int     `json:"tdp"`
}

type CpuParams struct {
	DefaultParams ProductParams 
	Cores      int	`form:"cores"` 
	Threads    int     `form:"threads"`
	BaseClock  float64 `form:"base_clock"`
	BoostClock float64 `form:"boost_clock"`
	Socket     string  `form:"socket"`
	Tdp        int     `form:"tdp"`
}