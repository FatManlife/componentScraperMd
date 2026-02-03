package repo

import (
	"context"

	"github.com/FatManlife/component-finder/back-end/internal/models/orm"
	"gorm.io/gorm"
)

type WebsiteRepository struct {
	db *gorm.DB
}

func NewWebsiteRepository(db *gorm.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}	

func (r *WebsiteRepository) GetAllWebsites(ctx context.Context) ([]string, error) {
	var websites []string

	q := r.db.Model(&orm.Website{}).Select("name")

	if err := q.Pluck("name",&websites).Error; err != nil {
		return nil, err
	}

	return websites, nil
}