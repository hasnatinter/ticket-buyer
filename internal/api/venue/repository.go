package venue

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, v *Venue) (*Venue, error) {
	queryDB := r.db.WithContext(ctx)
	if err := queryDB.Create(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
