package orm

type Aio struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string 
	ImageURL string 
	Brand string 
	Price float64 
	URL string 
	WebsiteID int

	Diagonal string 
	Cpu string 
	Ram string 
	Storage string 
	Gpu string

	Website Website `gorm:"foreignKey:WebsiteID;"`
}
