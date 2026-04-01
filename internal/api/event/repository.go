package event

import (
	"app/internal/api/ticket"
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
	queryDB := r.db.WithContext(ctx)
	queryDB = queryDB.Select("event.*, (select count(*) from ticket WHERE status = ? AND event_id = event.id) as total_tickets", ticket.Available.String())
	queryDB = queryDB.Joins("Venue")
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

func (r *Repository) ReadWithTickets(id string, ctx context.Context) (*Event, error) {
	event := &Event{}
	queryDB := r.db.WithContext(ctx)
	queryDB = r.db.Joins("Venue").Joins("Performer")
	queryDB.Select("event.*")
	queryDB.Where("event.id = ?", id)
	if err := queryDB.First(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *Repository) Create(ctx context.Context, event *Event) (*Event, error) {
	queryDB := r.db.WithContext(ctx)
	if err := queryDB.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}
