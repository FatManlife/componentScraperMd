package orm

type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"uniqueIndex;not null"`

	Products []Product `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}