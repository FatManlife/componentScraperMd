package orm

type Ssd struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Capacity int 
	ReadingSpeed int 
	WritingSpeed int 
	FormFactor string 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}