package orm

type Hdd struct  {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int


	Capacity int 
	RotationSpeed int 	
	FormFactor string 
	Website Website `gorm:"foreignKey:WebsiteID;"`
}