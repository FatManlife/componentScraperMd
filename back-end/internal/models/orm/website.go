package orm

type Website struct {
	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string
	URL  string
	Image string

	Products []Product `gorm:"foreignKey:WebsiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}