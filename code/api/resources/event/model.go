package event

import (
	"app/code/api/resources/performer"
	"app/code/api/resources/venue"
	"time"
)

type Event struct {
	ID           int64
	Name         string
	Description  string
	VenueId      int64
	StartTime    time.Time
	Venue        venue.Venue
	PerformerId  int64
	Performer    performer.Performer
	TotalTickets int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Events *[]Event
