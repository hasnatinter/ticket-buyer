package ticket

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

func (r *Repository) ListForEvent(eventId string, ctx context.Context) (Tickets, error) {
	tickets := make(Tickets, 0)
	queryDB := r.db.WithContext(ctx).Where("event_id = ?", eventId)
	if err := queryDB.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
