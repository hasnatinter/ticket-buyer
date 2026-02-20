package event

import (
	"context"
	"strconv"

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

func (r *Repository) ListWithTickets(filter *EventFilter, ctx context.Context) (Events, error) {
	events := make(Events, 0)
	queryDB := r.db.Joins("Venue")
	queryDB.WithContext(ctx)
	queryDB = queryDB.Select("*, (select count(*) from ticket WHERE status = 'available' AND event_id = event.id) as total_tickets")
	if len(filter.StartDate) > 0 {
		queryDB.Where("start_time >= ?", filter.StartDate)
	}
	if len(filter.EndDate) > 0 {
		queryDB.Where("end_time <= ?", filter.EndDate)
	}
	if len(filter.Category) > 0 {
		queryDB.Where("category = ?", filter.Category)
	}
	if len(filter.Venue) > 0 {
		queryDB.Where("venue_id = ?", filter.Venue)
	}
	if len(filter.Limit) > 0 {
		limit, _ := strconv.Atoi(filter.Limit)
		queryDB.Limit(limit)
	}
	if len(filter.Offset) > 0 {
		offset, _ := strconv.Atoi(filter.Offset)
		queryDB.Offset(offset)
	}
	if err := queryDB.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
