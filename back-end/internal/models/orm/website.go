package orm

type Website struct {
	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string
	URL  string
	Image string

	Aios []Aio `gorm:"foreignKey:WebsiteID;"`
	Cases []Case `gorm:"foreignKey:WebsiteID;"`
	Coolers []Cooler `gorm:"foreignKey:WebsiteID;"`
	Cpus []Cpu `gorm:"foreignKey:WebsiteID;"`
	Fans []Fan `gorm:"foreignKey:WebsiteID;"`
	Gpus []Gpu `gorm:"foreignKey:WebsiteID;"`
	Hdds []Hdd `gorm:"foreignKey:WebsiteID;"`
	Laptops []Laptop `gorm:"foreignKey:WebsiteID;"`
	Motherboards []Motherboard `gorm:"foreignKey:WebsiteID;"`
	PcMinis []PcMini `gorm:"foreignKey:WebsiteID;"`
	Pcs []Pc `gorm:"foreignKey:WebsiteID;"`
	Psus []Psu `gorm:"foreignKey:WebsiteID;"`
	Rams []Ram `gorm:"foreignKey:WebsiteID;"`
	Ssds []Ssd `gorm:"foreignKey:WebsiteID;"`
}