package orm

type Ssd struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Capacity int 
	ReadingSpeed int 
	WritingSpeed int 
	FormFactor string 
}
