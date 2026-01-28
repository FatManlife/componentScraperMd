package orm

type Product struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string 
	ImageURL string 
	Brand string 
	Price float64 
	URL string `gorm:"uniqueIndex;not null"`
	WebsiteID int

	Aio *Aio `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Case *PcCase `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cooler *Cooler `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cpu *Cpu `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Fan *Fan `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Gpu *Gpu `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Hdd *Hdd `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Laptop *Laptop `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Motherboard *Motherboard `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PcMini *PcMini `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Pc *Pc `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Psu *Psu `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ram *Ram `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ssd *Ssd`gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}