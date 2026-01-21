package orm

type Cooler struct {
    ID            int     `gorm:"primaryKey;autoIncrement"`
    ProductID     int

    Type          string
    FanRPM        int
    Noise         float64
    Size          string

    Compatibility []CoolerCompatibility `gorm:"foreignKey:CoolerID"`
}

type CoolerCompatibility struct {
    ID            int    `gorm:"primaryKey;autoIncrement"`
    CoolerID      int
    Compatibility string
}
