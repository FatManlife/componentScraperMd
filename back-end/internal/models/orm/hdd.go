package orm

type Hdd struct  {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Capacity int 
	RotationSpeed int 	
	FormFactor string 
}