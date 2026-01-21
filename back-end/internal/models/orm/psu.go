package orm

type Psu struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Power int 
	Efficiency string 
	FormFactor string 
}