package event

import (
	"app/code/api/resources/performer"
	"app/code/api/resources/ticket"
	"app/code/api/resources/venue"
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
	TotalTickets int
	CreatedAt    *time.Time `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoCreateTime"`
}

type TicketDTO struct {
	ID   string `json:"id"`
	Seat string `json:"seat"`
}

type EventDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StartTime    string `json:"start_time"`
	Category     string `json:"category"`
	Venue        string `json:"venue"`
	TotalTickets string `json:"total_available_tickets"`
}

func (e *Event) ToDTO() *EventDTO {
	return &EventDTO{
		ID:           strconv.Itoa(e.ID),
		Name:         e.Name,
		StartTime:    e.StartTime.String(),
		Venue:        e.Venue.Name,
		Category:     e.Category.String,
		TotalTickets: strconv.Itoa(e.TotalTickets),
	}
}

func (es Events) ToDTO() []*EventDTO {
	dto := make([]*EventDTO, len(es))
	for i, v := range es {
		dto[i] = v.ToDTO()
	}
	return dto
}

type Events []Event
