package ticket

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID        int64
	Seat      string
	UserID    int64
	Status    string
	EventId   int64
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoCreateTime"`
	DeletedAt *time.Time
}

type Tickets []Ticket

type TicketDTO struct {
	ID   int64  `json:"id"`
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
