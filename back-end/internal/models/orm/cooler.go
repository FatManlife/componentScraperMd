package orm

type Cooler struct {
    ID            int     `gorm:"primaryKey;autoIncrement"`
    ProductID     int

    Type          string
    FanRPM        int
    Noise         float64
    Size          string

    Cpus []CoolerCpu `gorm:"foreignKey:CoolerID;constraint:OnDelete:CASCADE"` 
}


