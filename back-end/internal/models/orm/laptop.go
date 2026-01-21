package orm

type Laptop struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int
	
	Cpu string 
	Gpu string 
	Ram string 
	Storage string 
	Diagonal string 
	Battery float64 
}