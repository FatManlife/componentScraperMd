package orm

type PcCase struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Format                string 
	MotherboardFormFactor string 
}
