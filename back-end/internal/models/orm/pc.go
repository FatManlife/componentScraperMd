package orm

type Pc struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
	Motherboard string 
	Psu string 
	Case string 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}