package performer

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

func (r *Repository) Create(ctx context.Context, p *Performer) (*Performer, error) {
	queryDB := r.db.WithContext(ctx)
	if err := queryDB.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
