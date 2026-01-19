package orm

type Fan struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL string
	WebsiteID  int

	FanRPM int 
	Noise float64 
	Size string 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}