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
	Cores      int	 
	Threads    int     
	BaseClock  float64
	BoostClock float64
	Socket     string  
	Tdp        int     
}