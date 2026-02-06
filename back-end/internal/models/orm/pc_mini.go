package orm

type PcMini struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Cpu string 
	Gpu string 
	Ram int 
	Storage int 
}