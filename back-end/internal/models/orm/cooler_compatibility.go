package orm

type CoolerCompatibility struct {
    ID            int    `gorm:"primaryKey;autoIncrement"`
    Cpu string
}
