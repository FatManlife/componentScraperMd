package orm

type Laptop struct {
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
	Diagonal string 
	Battery float64 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}