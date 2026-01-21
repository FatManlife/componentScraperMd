package orm

type Case struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Format                string 
	MotherboardFormFactor string 
}
