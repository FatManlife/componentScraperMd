package repo

import "gorm.io/gorm"

type SSDRepository struct {
	db *gorm.DB
}

func NewSSDRepository(db *gorm.DB) *SSDRepository {
	return &SSDRepository{db: db}
}
