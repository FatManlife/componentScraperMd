package orm

type Psu struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Power int 
	Efficiency string 
	FormFactor string 

	Webiste Website `gorm:"foreignKey:WebsiteId;"`
}