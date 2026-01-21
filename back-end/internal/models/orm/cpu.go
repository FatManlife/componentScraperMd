package orm

type Cpu struct {
	ID 		   int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Cores      int     
	Threads    int     
	BaseClock  float64 
	BoostClock float64 
	Socket     string  
	Tdp        int     
}