package event

import (
	"app/internal/api/resources/performer"
	"app/internal/api/resources/ticket"
	"app/internal/api/resources/venue"
	"database/sql"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID           int
	Name         string
	Description  sql.NullString
	Category     sql.NullString
	VenueId      int64
	Venue        venue.Venue
	StartTime    *time.Time
	PerformerId  int64
	Performer    performer.Performer
	Tickets      []ticket.Ticket
	TotalTickets *int       `gorm:"->;column:total_tickets"`
	CreatedAt    *time.Time `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoCreateTime"`
}

type EventDTO struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	StartTime    string             `json:"start_time"`
	Category     string             `json:"category"`
	Venue        string             `json:"venue"`
	Tickets      []ticket.TicketDTO `json:"tickets,omitempty"`
	TotalTickets string             `json:"total_available_tickets,omitempty"`
}

func (e *Event) ToDTO(tickets ticket.Tickets) *EventDTO {
	var ticketsDTO []ticket.TicketDTO
	var TotalTickets string
	if e.TotalTickets != nil {
		TotalTickets = strconv.Itoa(*e.TotalTickets)
	}
	if len(tickets) > 0 {
		ticketsDTO = make([]ticket.TicketDTO, len(tickets))
		for i, v := range tickets {
			ticketsDTO[i] = ticket.TicketDTO{
				ID:   v.ID,
				Seat: v.Seat,
			}
		}
	}
	return &EventDTO{
		ID:           strconv.Itoa(e.ID),
		Name:         e.Name,
		StartTime:    e.StartTime.String(),
		Venue:        e.Venue.Name,
		Category:     e.Category.String,
		Tickets:      ticketsDTO,
		TotalTickets: TotalTickets,
	}
}

func (es Events) ToDTO() []*EventDTO {
	dto := make([]*EventDTO, len(es))
	var tickets []ticket.Ticket
	for i, v := range es {
		dto[i] = v.ToDTO(tickets)
	}
	return dto
}

type Events []Event
