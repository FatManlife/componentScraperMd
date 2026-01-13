package models

type Cpu struct {
	BaseAttrs BaseProduct 
	Cores int `json:"cores"`	
	Threads int `json:"threads"`
	BaseClock float64 `json:"base_clock"`
	BoostClock float64 `json:"boost_clock"`
	Cache string `json:"cache"`
	Socket string `json:"socket"`
	Tdp int `json:"tdp"`
}