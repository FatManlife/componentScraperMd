package orm

type Case struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Format                string 
	MotherboardFormFactor string 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}
