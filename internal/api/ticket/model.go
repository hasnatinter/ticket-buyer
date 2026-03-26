package ticket

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID        int
	Seat      string
	UserID    int
	Status    string
	EventId   int
	BookingId sql.NullInt32
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoCreateTime"`
	DeletedAt *time.Time
}

type Tickets []Ticket

type TicketDTO struct {
	ID   int    `json:"id"`
	Seat string `json:"seat"`
}

func (t *Ticket) ToDTO() *TicketDTO {
	return &TicketDTO{
		ID:   t.ID,
		Seat: t.Seat,
	}
}

func (t Tickets) ToDTO() []*TicketDTO {
	dto := make([]*TicketDTO, len(t))
	for i, v := range t {
		dto[i] = v.ToDTO()
	}
	return dto
}
