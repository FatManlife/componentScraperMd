package orm

type CoolerCpu struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	CoolerID int
	CompatibilityID int
}