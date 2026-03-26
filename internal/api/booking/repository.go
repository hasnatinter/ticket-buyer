package booking

import (
	"app/internal/api/ticket"
	"context"
	"fmt"

	"gorm.io/gorm"
)

//go:generate stringer --type=BookingStatus
type BookingStatus int

const (
	reserved BookingStatus = iota
	booked
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create bookings in transcations, to prevent double booking against a ticket
func (r *Repository) AddBooking(c CreateFilter, ctx context.Context) (*Booking, error) {
	tx := r.db.Begin()
	defer tx.Rollback()

	tRepo := ticket.NewRepository(tx)
	tickets, err := tRepo.GetAvailableById(c.Tickets, ctx)
	// If tickets length < ids length return specific error
	if err != nil {
		return nil, err
	}

	// If total tickets are not equal to sent tickets then some were
	// already booked in another transaction and we should return error
	if len(tickets) != len(c.Tickets) {
		return nil, fmt.Errorf("some tickets are already booked")
	}

	ticketIds := make([]int32, 0)
	for _, t := range tickets {
		ticketIds = append(ticketIds, int32(t.ID))
	}

	// Create new booking
	booking := &Booking{
		UserName:    c.UserName,
		UserAddress: c.UserAddress,
		TicketIds:   ticketIds,
	}
	queryDB := tx.WithContext(ctx)
	if err := queryDB.Create(booking).Error; err != nil {
		return nil, err
	}

	// Update ticket status and attach booking id
	for _, t := range tickets {
		err = tRepo.BookTicket(&t, booking.ID, ctx)
		if err != nil {
			return nil, err
		}
	}

	booking.Ticket = tickets

	return booking, tx.Commit().Error
}
