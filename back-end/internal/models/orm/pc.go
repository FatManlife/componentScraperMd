package orm

type Pc struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
	Motherboard string 
	Psu string 
	PcCase string 
}