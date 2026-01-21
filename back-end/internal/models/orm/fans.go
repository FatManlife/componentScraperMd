package orm

type Fan struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	FanRPM int 
	Noise float64 
	Size string 
}