package booking

import (
	"app/internal/api/ticket"
	"strconv"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID          int
	UserName    string
	UserAddress string
	TicketIds   pq.Int32Array `gorm:"type:int"`
	Ticket      []ticket.Ticket
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoCreateTime"`
	DeletedAt   *time.Time
}

type Bookings []Booking

type BookingDTO struct {
	ID          string              `json:"id"`
	UserName    string              `json:"user_name"`
	UserAddress string              `json:"user_address"`
	Ticket      []*ticket.TicketDTO `json:"tickets"`
}

func (b *Booking) ToDTO(tickets ticket.Tickets) *BookingDTO {
	return &BookingDTO{
		ID:          strconv.Itoa(b.ID),
		UserName:    b.UserName,
		UserAddress: b.UserAddress,
		Ticket:      tickets.ToDTO(),
	}
}
