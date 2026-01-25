package orm

type CoolerCompatibility struct {
    ID            int    `gorm:"primaryKey;autoIncrement"`
    Cpu string

    Coolers []CoolerCpu `gorm:"foreignKey:CompatibilityID;references:ID;constraint:OnDelete:SET NULL"`
}
