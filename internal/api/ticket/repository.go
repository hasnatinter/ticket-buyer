package ticket

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate stringer -type=TicketState
type TicketState int

const (
	Booked TicketState = iota
	Available
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(id int, ctx context.Context) (*Ticket, error) {
	ticket := Ticket{}
	queryDB := r.db.WithContext(ctx).Where("id = ?", id)
	if err := queryDB.First(&ticket).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *Repository) GetAvailableById(ids []int, ctx context.Context) (Tickets, error) {
	var tickets []Ticket
	err := r.db.WithContext(ctx).
		Where("ID IN ?", ids).
		Where("status = ?", Available.String()).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&tickets).
		Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *Repository) ListForEvent(eventId string, ctx context.Context) (Tickets, error) {
	var tickets []Ticket
	err := r.db.WithContext(ctx).
		Where("event_id = ?", eventId).
		Find(&tickets).
		Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *Repository) BookTicket(ticket *Ticket, booking_id int, ctx context.Context) error {
	return r.db.WithContext(ctx).Model(ticket).Updates(map[string]interface{}{"status": Booked, "booking_id": booking_id}).Error
}

func (r *Repository) Create(ctx context.Context, t *Ticket) (*Ticket, error) {
	queryDB := r.db.WithContext(ctx)
	if err := queryDB.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
