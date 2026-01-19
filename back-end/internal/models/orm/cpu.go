package orm

type Cpu struct {
	ID 		   int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Cores      int     
	Threads    int     
	BaseClock  float64 
	BoostClock float64 
	Socket     string  
	Tdp        int     

	Website Website `gorm:"foreignKey:WebsiteID;"`	
}