package orm

type Ram struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Capacity int 
	Speed int 
	Type string 
	Compatibility string 
	Configuration string 
}