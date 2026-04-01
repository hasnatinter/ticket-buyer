package booking_test

import (
	"app/internal/api/booking"
	"app/internal/api/event"
	"app/internal/api/performer"
	"app/internal/api/ticket"
	"app/internal/api/venue"
	"app/pkg/helpers"
	"database/sql"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestConcurrentBooking(t *testing.T) {
	db := helpers.SetUp(t)

	e := CreateEvent(t, db)

	tRepo := ticket.NewRepository(db)
	tk := &ticket.Ticket{
		Seat:    "24H",
		Status:  ticket.Available.String(),
		EventId: e.ID,
	}
	_, err := tRepo.Create(t.Context(), tk)
	if err != nil {
		t.Fatal(err)
	}

	bRepo := booking.NewRepository(db)
	bFilter := booking.CreateFilter{
		UserName:    "Test user",
		UserAddress: "street 123",
		Tickets:     []int{tk.ID},
	}
	var wg sync.WaitGroup

	var totalErrors atomic.Int32
	for i := 0; i < 100; i++ {
		wg.Go(func() {
			_, err := bRepo.AddBooking(bFilter, t.Context())
			if err != nil {
				totalErrors.Add(1)
			}
		})
	}
	wg.Wait()

	if totalErrors.Load() != 99 {
		t.Fatal("Total concurrent booking errors not correct")
	}

	bk, err := bRepo.FetchBookings(t.Context())
	if len(bk) != 1 {
		t.Fatal("More than one booking found")
	}
}

func CreateEvent(t *testing.T, db *gorm.DB) *event.Event {
	performerRepo := performer.NewRepository(db)
	p := &performer.Performer{
		Name: "New performer",
	}
	_, err := performerRepo.Create(t.Context(), p)
	if err != nil {
		t.Fatal(err)
	}

	venueRepo := venue.NewRepository(db)
	v := &venue.Venue{
		Name: "New venue",
	}
	_, err = venueRepo.Create(t.Context(), v)
	if err != nil {
		t.Fatal(err)
	}

	eventRepo := event.NewRepository(db)
	tm := time.Now().Add(1 * time.Hour).UTC()
	e := &event.Event{
		Name:        "New event",
		Description: sql.NullString{String: "New description", Valid: true},
		Category:    sql.NullString{String: "Musik", Valid: true},
		VenueId:     v.ID,
		PerformerId: p.ID,
		StartTime:   &tm,
	}
	_, err = eventRepo.Create(t.Context(), e)
	if err != nil {
		t.Fatal(err)
	}

	return e
}
