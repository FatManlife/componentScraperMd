package orm

type Ram struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Capacity int 
	Speed int 
	Type string 
	Compatibility string 
	Configuration string 
	
	Website Website `gorm:"foreignKey:WebsiteID;"`
}