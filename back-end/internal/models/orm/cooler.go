package orm

type Cooler struct {
    ID            int     `gorm:"primaryKey;autoIncrement"`
    Name          string
    ImageURL      string
    Brand         string
    Price         float64
    URL           string
    WebsiteID     int

    Type          string
    FanRPM        int
    Noise         float64
    Size          string

    Website       Website `gorm:"foreignKey:WebsiteID"`           
    Compatibility []CoolerCompatibility `gorm:"foreignKey:CoolerID"`
}

type CoolerCompatibility struct {
    ID            int    `gorm:"primaryKey;autoIncrement"`
    CoolerID      int
    Compatibility string
}
