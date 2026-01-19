package orm

type Gpu struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	ImageURL   string
	Brand      string
	Price      float64
	URL        string
	WebsiteID  int

	Chipset string 
	Vram int 
	GpuFrequency int 
	VramFrequency int 

	Website Website `gorm:"foreignKey:WebsiteID;"`
}