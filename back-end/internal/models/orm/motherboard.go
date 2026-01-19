package orm

type Motherboard struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Chipset string 
	Socket string 
	FormFactor string 
	RamSupport string 
	FormFactorRam string 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}