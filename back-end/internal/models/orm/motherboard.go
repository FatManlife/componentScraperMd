package orm

type Motherboard struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	ProductID  int
	
	Chipset string 
	Socket string 
	FormFactor string 
	RamSupport string 
	FormFactorRam string 
}