package orm

type Pc struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Cpu string 
	Gpu string 
	Ram int
	Storage int
	Motherboard string 
	Psu string 
	PcCase string 
}