package orm

type Aio struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	ProductID int

	Diagonal float64
	Cpu string 
	Ram int
	Storage int 
	Gpu string
}
