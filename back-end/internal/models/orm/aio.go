package orm

type Aio struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	ProductID int

	Diagonal string 
	Cpu string 
	Ram string 
	Storage string 
	Gpu string
}
