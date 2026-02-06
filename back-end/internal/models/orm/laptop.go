package orm

type Laptop struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int
	
	Cpu string 
	Gpu string 
	Ram int
	Storage int
	Diagonal float64 
	Battery float64 
}