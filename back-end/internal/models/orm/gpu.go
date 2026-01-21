package orm

type Gpu struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ProductID  int

	Chipset string 
	Vram int 
	GpuFrequency int 
	VramFrequency int 
}